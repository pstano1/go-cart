package http

import (
	"encoding/json"

	"github.com/pstano1/go-cart/internal/pkg"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

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

func (i *HTTPInstanceAPI) deleteUser(ctx *fasthttp.RequestCtx) {
	i.log.Debug("got request for user deletion")
	userId := ctx.UserValue("id").(string)
	if userId == "" {
		ctx.SetBodyString(pkg.ErrUnableToReadPayload.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	if err := i.api.DeleteUser(&pkg.User{Id: userId}); err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	ctx.SetBodyString("deleted")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

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
