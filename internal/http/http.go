// Package http provides server implementation for the application
// It includes route definitions, middlewares & handlers for various API endpoints
package http

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fasthttp/router"
	"github.com/gbrlsnchs/jwt"
	gojwt "github.com/golang-jwt/jwt"
	"github.com/jinzhu/copier"
	"github.com/mitchellh/mapstructure"
	"github.com/pstano1/go-cart/internal/api"
	"github.com/pstano1/go-cart/internal/pkg"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"
)

// Constants used through the `http` package
const (
	bearerToken          = "Bearer"
	corsAllowHeaders     = "Content-Type,Authorization"
	corsAllowMethods     = "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS"
	corsAllowOrigin      = "*"
	corsAllowCredentials = "true"
)

// HTTPInstanceAPI represents the HTTP server instance and its configuration
type HTTPInstanceAPI struct {
	bind string
	log  *zap.Logger
	api  *api.InstanceAPI
}

// HTTPConfig holds the configuration for initializing an HTTPInstanceAPI
type HTTPConfig struct {
	Logger   *zap.Logger
	BindPath string
	API      *api.InstanceAPI
}

// NewHTTPInstanceAPI creates a new instance of HTTPInstanceAPI with the given configuration
func NewHTTPInstanceAPI(conf *HTTPConfig) *HTTPInstanceAPI {
	return &HTTPInstanceAPI{
		bind: conf.BindPath,
		log:  conf.Logger,
		api:  conf.API,
	}
}

// GetRouter sets up the router and defines all the API routes
func (i *HTTPInstanceAPI) GetRouter() *router.Router {
	r := router.New()
	api := r.Group("/api")

	api.GET("/error/{lang}", i.getError)

	customers := api.Group("/customer")
	customers.GET("/id/{tag}", i.getCustomerId)

	users := api.Group("/user")
	users.GET("/", i.authMiddleware(i.sameCustomerOperation(i.getUser)))
	users.GET("/permission", i.authMiddleware(i.getPermission))
	users.POST("/", i.createUser)
	// users.POST("/", i.authMiddleware(i.sameCustomerOperation(i.createUser)))
	users.PUT("/", i.authMiddleware(i.updateUser))
	users.DELETE("/{id}", i.authMiddleware(i.deleteUser))
	users.POST("/signin", i.signUserIn)
	users.POST("/refresh", i.refreshToken)

	products := api.Group("/product")
	products.GET("/", i.getProduct)
	products.POST("/", i.authMiddleware(i.sameCustomerOperation(i.createProduct)))
	products.PUT("/", i.authMiddleware(i.sameCustomerOperation(i.updateProduct)))
	products.DELETE("/{id}", i.authMiddleware(i.sameCustomerOperation(i.deleteProduct)))
	products.GET("/category", i.getCategory)
	products.POST("/category", i.authMiddleware(i.sameCustomerOperation(i.createCategory)))
	products.PUT("/category", i.authMiddleware(i.sameCustomerOperation(i.updateCategory)))
	products.DELETE("/category/{id}", i.authMiddleware(i.sameCustomerOperation(i.deleteCategory)))

	coupons := api.Group("/coupon")
	coupons.GET("/", i.getCoupon)
	coupons.POST("/", i.authMiddleware(i.sameCustomerOperation(i.createCoupon)))
	coupons.PUT("/", i.authMiddleware(i.sameCustomerOperation(i.updateCoupon)))
	coupons.DELETE("/{id}", i.authMiddleware(i.sameCustomerOperation(i.deleteCoupon)))

	orders := api.Group("/order")
	orders.GET("/", i.authMiddleware(i.sameCustomerOperation(i.getOrder)))
	orders.POST("/", i.createOrder)
	orders.PUT("/", i.authMiddleware(i.sameCustomerOperation(i.updateOrder)))
	orders.DELETE("/{id}", i.authMiddleware(i.sameCustomerOperation(i.deleteOrder)))

	return r
}

// Run starts the HTTP server
func (i *HTTPInstanceAPI) Run() {
	swaggerUIHandler := func(ctx *fasthttp.RequestCtx) {
		if string(ctx.Path()) == "/swagger/" || string(ctx.Path()) == "/swagger/index.html" {
			ctx.SendFile("./swagger/index.html")
			return
		}
		ctx.SendFile("./" + string(ctx.Path()))
	}
	r := i.GetRouter()
	r.GET("/swagger/{any}", swaggerUIHandler)
	i.log.Info("Starting server at port",
		zap.String("port", i.bind),
	)
	log.Fatal(fasthttp.ListenAndServe(i.bind, i.corsMiddleware(r.Handler)))
}

// authMiddleware checks if user is authenticated when accessing protected routes
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

// corsMiddleware handles preflights and allows for cross domain operations
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

// sameCustomerOperation is middleware that handles request to see
// if request made to route protected this way contains `customerId`
// then check if user requesting has the same `customerId`
// so the server will not respond with different customer's data
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

// validatePermissions checks if user possess' any od required permissions
func validatePermissions(requiredPermissions []string, ctx *fasthttp.RequestCtx) bool {
	ctx.SetUserValue("startTime", time.Now())
	token := string(ctx.Request.Header.Peek("Authorization"))
	token = strings.TrimPrefix(token, bearerToken+" ")
	parsedToken, _, err := new(gojwt.Parser).ParseUnverified(token, gojwt.MapClaims{})
	if err != nil {
		return false
	}
	claims, ok := parsedToken.Claims.(gojwt.MapClaims)
	if !ok {
		return false
	}
	var jwtConfig api.JWTConfig
	if err := mapstructure.Decode(claims, &jwtConfig); err != nil {
		return false
	}
	content, ok := claims["user"].(map[string]interface{})
	if !ok {
		return false
	}
	userPermissions, ok := content["permissions"].([]interface{})
	if !ok {
		return false
	}
	permissionSet := make(map[string]bool)
	for _, inter := range userPermissions {
		permission, ok := inter.(string)
		if !ok {
			return false
		}
		permissionSet[permission] = true
	}
	for _, permission := range requiredPermissions {
		if permissionSet[permission] {
			return true
		}
	}
	return false
}

// validateBody unmarshals request body into `T` type struct
// & validates it's integrity
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

// validateFilter populates filter of specified `T` type
// & validates it's integrity
func validateFilter[T pkg.Filter](ctx *fasthttp.RequestCtx) (T, error) {
	ctx.SetUserValue("startTime", time.Now())
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

// getCustomerId handles a translation of tag into customerId
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

// getError parses the error dict file of specified language
// & returns it to user as response body
func (i *HTTPInstanceAPI) getError(ctx *fasthttp.RequestCtx) {
	i.log.Debug("got request for retrieving error messages")
	supportedLanguages := []string{"en", "pl"}
	requestedLanguage := ctx.UserValue("lang").(string)
	ok := false
	for _, lang := range supportedLanguages {
		if lang == strings.ToLower(requestedLanguage) {
			ok = true
			break
		}
	}
	if !ok {
		ctx.SetBodyString(pkg.ErrUnsupportedLang.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	filename := fmt.Sprintf("common/errors/%s.json", strings.ToLower(requestedLanguage))
	errors, err := os.ReadFile(filename)
	if err != nil {
		ctx.SetBodyString(pkg.ErrUnsupportedLang.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	ctx.Response.SetBody(errors)
	ctx.SetStatusCode(fasthttp.StatusOK)
}
