package http

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/fasthttp/router"
	"github.com/gbrlsnchs/jwt"
	"github.com/jinzhu/copier"
	"github.com/pstano1/go-cart/internal/api"
	"github.com/pstano1/go-cart/internal/pkg"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"
)

const (
	bearerToken          = "Bearer"
	corsAllowHeaders     = "Content-Type,Authorization"
	corsAllowMethods     = "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS"
	corsAllowOrigin      = "*"
	corsAllowCredentials = "true"
)

type HTTPInstanceAPI struct {
	bind string
	log  *zap.Logger
	api  *api.InstanceAPI
}

type HTTPConfig struct {
	Logger   *zap.Logger
	BindPath string
	API      *api.InstanceAPI
}

func NewHTTPInstanceAPI(conf *HTTPConfig) *HTTPInstanceAPI {
	return &HTTPInstanceAPI{
		bind: conf.BindPath,
		log:  conf.Logger,
		api:  conf.API,
	}
}

func (i *HTTPInstanceAPI) GetRouter() *router.Router {
	r := router.New()
	api := r.Group("/api")

	customers := api.Group("/customer")
	customers.GET("/id/{tag}", i.getCustomerId)

	users := api.Group("/user")
	users.GET("/", i.authMiddleware(i.sameCustomerOperation(i.getUser)))
	users.POST("/", i.createUser)
	// users.POST("/", i.authMiddleware(i.createUser))
	users.PATCH("/", i.authMiddleware(i.updateUser))
	users.DELETE("/{id}", i.authMiddleware(i.deleteUser))
	users.POST("/signin", i.signUserIn)
	users.POST("/refresh", i.refreshToken)

	products := api.Group("/product")
	products.GET("/", i.authMiddleware(i.sameCustomerOperation(i.getProduct)))
	products.POST("/", i.authMiddleware(i.sameCustomerOperation(i.createProduct)))
	products.PATCH("/", i.authMiddleware(i.sameCustomerOperation(i.updateProduct)))
	products.DELETE("/{id}", i.authMiddleware(i.sameCustomerOperation(i.deleteProduct)))

	_ = api.Group("/orders")

	return r
}

func (i *HTTPInstanceAPI) Run() {
	r := i.GetRouter()
	i.log.Info("Starting server at port",
		zap.String("port", i.bind),
	)
	log.Fatal(fasthttp.ListenAndServe(i.bind, i.corsMiddleware(r.Handler)))
}

func (i *HTTPInstanceAPI) authMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		path := ctx.Path()
		i.log.Debug("checking authentication for",
			zap.ByteString("path", path),
		)

		token := string(ctx.Request.Header.Peek("Authorization"))
		if token == "" {
			ctx.Error("user unauthenticated", fasthttp.StatusUnauthorized)
			return
		}
		token = strings.TrimPrefix(token, bearerToken+" ")
		rawDecodeText, err := jwt.FromString(token)
		if err != nil {
			ctx.Error("user unauthenticated", fasthttp.StatusUnauthorized)
			return
		}
		time_exp := time.Until(rawDecodeText.ExpirationTime())
		if time_exp.Seconds() < 0 {
			ctx.Error("user unauthenticated", fasthttp.StatusUnauthorized)
			return
		}

		next(ctx)
	}
}

func (i *HTTPInstanceAPI) corsMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", corsAllowOrigin)
		ctx.Response.Header.Set("Access-Control-Allow-Methods", corsAllowMethods)
		ctx.Response.Header.Set("Access-Control-Allow-Headers", corsAllowHeaders)
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", corsAllowCredentials)

		if string(ctx.Method()) == "OPTIONS" {
			ctx.SetStatusCode(fasthttp.StatusNoContent)
			return
		}

		next(ctx)
	}
}

func (i *HTTPInstanceAPI) sameCustomerOperation(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		var customerId []byte
		if string(ctx.Method()) == "GET" || string(ctx.Method()) == "DELETE" {
			customerId = ctx.Request.URI().QueryArgs().Peek("customerId")
			if customerId == nil {
				ctx.Error("no customerId specified", fasthttp.StatusBadRequest)
				return
			}
		} else {
			payload, err := validateBody[pkg.CustomerSpecificModel](ctx)
			if err != nil {
				ctx.Error("error while processing payload", fasthttp.StatusInternalServerError)
				return
			}
			if payload.CustomerId == "" {
				ctx.Error("no customerId specified", fasthttp.StatusBadRequest)
				return
			}
			customerId = []byte(payload.CustomerId)
		}
		if ok, _ := i.api.ValidateCustomerId(string(customerId)); !ok {
			ctx.Error("customer with this id does not exist", fasthttp.StatusBadRequest)
			return
		}
		user, err := i.api.GetUserFromRequest(ctx)
		if err != nil {
			ctx.Error("error while retrieving user", fasthttp.StatusInternalServerError)
			return
		}
		if user.CustomerId != string(customerId) {
			ctx.Error("cross customer operation", fasthttp.StatusForbidden)
			return
		}

		next(ctx)
	}
}

func validateBody[T any](ctx *fasthttp.RequestCtx) (*T, error) {
	ctx.SetUserValue("startTime", time.Now())
	var postBody T
	validate := validator.New()
	body := ctx.Request.Body()
	if err := json.Unmarshal(body, &postBody); err != nil {
		zap.Error(err)
		return &postBody, pkg.ErrUnableToReadPayload
	}
	if err := validate.Struct(postBody); err != nil {
		return &postBody, err
	}
	return &postBody, nil
}

func validateFilter[T pkg.Filter](ctx *fasthttp.RequestCtx) (T, error) {
	ctx.SetUserValue("start_time", time.Now())
	var filter T
	validate := validator.New()
	f := filter.Populate(ctx)
	if err := validate.Struct(f); err != nil {
		return filter, err
	}
	err := copier.Copy(&filter, f)
	if err != nil {
		return filter, err
	}
	return filter, nil
}

func (i *HTTPInstanceAPI) getCustomerId(ctx *fasthttp.RequestCtx) {
	i.log.Debug("got request for getting customer id")
	tag := ctx.UserValue("tag").(string)
	if tag == "" {
		ctx.SetBodyString(pkg.ErrUnableToReadPayload.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	response, err := i.api.ExchangeTagForId(tag)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	body, err := json.Marshal(response)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBody(body)
	ctx.SetStatusCode(fasthttp.StatusOK)
}
