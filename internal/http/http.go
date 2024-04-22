package http

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/fasthttp/router"
	"github.com/gbrlsnchs/jwt"
	"github.com/pstano1/go-cart/internal/api"
	"github.com/pstano1/go-cart/internal/pkg"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"
)

const (
	bearerToken = "Bearer"
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

	users := api.Group("/user")
	users.GET("/", i.authMiddleware(i.getUser))
	users.POST("/", i.createUser)
	// users.POST("/", i.authMiddleware(i.createUser))
	users.PATCH("/", i.authMiddleware(i.updateUser))
	users.DELETE("/{id}", i.authMiddleware(i.deleteUser))
	users.POST("/signin", i.signUserIn)
	users.POST("/refresh", i.refreshToken)

	return r
}

func (i *HTTPInstanceAPI) Run() {
	r := i.GetRouter()
	i.log.Info("Starting server at port",
		zap.String("port", i.bind),
	)
	s := &fasthttp.Server{
		Handler: r.Handler,
		Name:    "Casker",
	}
	log.Fatal(s.ListenAndServe(i.bind))
}

func (i *HTTPInstanceAPI) authMiddleware(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
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

		handler(ctx)
	}
}

func validateBody[T any](ctx *fasthttp.RequestCtx) (*T, error) {
	ctx.SetUserValue("startTime", time.Now())
	var postBody T
	validate := validator.New()
	body := ctx.Request.Body()
	if err := json.Unmarshal(body, &postBody); err != nil {
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
	return filter, nil
}
