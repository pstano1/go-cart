package http

import (
	"encoding/json"

	"github.com/pstano1/go-cart/internal/pkg"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

// getUser retrieves user(s) based on provided query params
// (or lack of) and returns slice of pkg.User to user
func (i *HTTPInstanceAPI) getUser(ctx *fasthttp.RequestCtx) {
	i.log.Debug("got request for retrieving users")
	filter, err := validateFilter[pkg.UserFilter](ctx)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	users, err := i.api.GetUsers(&filter)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(users)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBody(response)
	ctx.SetStatusCode(fasthttp.StatusOK)
}

// getPermission provides user with list of possible permissions
func (i *HTTPInstanceAPI) getPermission(ctx *fasthttp.RequestCtx) {
	i.log.Debug("got request for permission retrieval")
	if ok := validatePermissions([]string{pkg.CreateUser}, ctx); !ok {
		ctx.SetBodyString(pkg.ErrUserForbidden.Error())
		ctx.SetStatusCode(fasthttp.StatusForbidden)
		return
	}
	permissions, err := i.api.GetPermissions()
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(permissions)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBody(response)
	ctx.SetStatusCode(fasthttp.StatusOK)
}

// createUser handles user creation based on request's body
func (i *HTTPInstanceAPI) createUser(ctx *fasthttp.RequestCtx) {
	i.log.Debug("got request for user creation")
	request, err := validateBody[pkg.UserCreate](ctx)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	if request.Password != request.PasswordCheck {
		ctx.SetBodyString("passwords don't match")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	if _, err := i.api.CreateUser(request); err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBodyString("successfully created user")
	ctx.SetStatusCode(fasthttp.StatusCreated)
}

// updateUser handles user update based on request's body
func (i *HTTPInstanceAPI) updateUser(ctx *fasthttp.RequestCtx) {
	i.log.Debug("got request for user update")
	request, err := validateBody[pkg.UserUpdate](ctx)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	if err = i.api.UpdateUser(request); err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBodyString("successfully updated user")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

// deleteUser deletes user with id specified in route
func (i *HTTPInstanceAPI) deleteUser(ctx *fasthttp.RequestCtx) {
	i.log.Debug("got request for user deletion")
	userId := ctx.UserValue("id").(string)
	if userId == "" {
		ctx.SetBodyString(pkg.ErrUnableToReadPayload.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	if err := i.api.Delete(&pkg.User{Id: userId}); err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	ctx.SetBodyString("deleted")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

// signUserIn attemps to sign user in based on provided credentials
func (i *HTTPInstanceAPI) signUserIn(ctx *fasthttp.RequestCtx) {
	i.log.Debug("got request to sign user in")
	request, err := validateBody[pkg.Credentials](ctx)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	response, err := i.api.SignUserIn(request)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
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

// refreshToken allows for prolonging user session
func (i *HTTPInstanceAPI) refreshToken(ctx *fasthttp.RequestCtx) {
	i.log.Debug("got request for access token refresh")
	user, err := i.api.GetUserFromRequest(ctx)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetUserValue("user", user.Username)
	response, err := i.api.RefreshUserJWT(user)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	i.log.Debug("successfully refreshed token for user",
		zap.String("username", user.Username),
	)
	body, err := json.Marshal(response)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}
	ctx.SetBody(body)
	ctx.SetStatusCode(fasthttp.StatusOK)
}
