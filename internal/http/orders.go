// Package http provides server implementation for the application
// This file includes /order routes handlers
package http

import (
	"encoding/json"

	"github.com/pstano1/go-cart/internal/pkg"
	"github.com/valyala/fasthttp"
)

// getCoupon retrieves order(s) based on provided query params
// (or lack of) and returns slice of pkg.Order to user
func (i *HTTPInstanceAPI) getOrder(ctx *fasthttp.RequestCtx) {
	i.log.Debug("got request for retrieving orders")
	if ok := validatePermissions([]string{pkg.GetOrders}, ctx); !ok {
		ctx.SetBodyString(pkg.ErrUserForbidden.Error())
		ctx.SetStatusCode(fasthttp.StatusForbidden)
		return
	}
	filter, err := validateFilter[pkg.OrderFilter](ctx)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	orders, err := i.api.GetOrders(&filter)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(orders)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBody(response)
	ctx.SetStatusCode(fasthttp.StatusOK)
}

// createOrder handles order creation based on request's body
// & returns object of type pkg.ObjectCreateResponse
func (i *HTTPInstanceAPI) createOrder(ctx *fasthttp.RequestCtx) {
	i.log.Debug("got request for creating order")
	request, err := validateBody[pkg.OrderCreate](ctx)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	orderId, err := i.api.CreateOrder(request)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(&pkg.ObjectCreateResponse{
		Id: *orderId,
	})
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBody(response)
	ctx.SetStatusCode(fasthttp.StatusCreated)
}

// updateOrder handles order update based on request's body
func (i *HTTPInstanceAPI) updateOrder(ctx *fasthttp.RequestCtx) {
	i.log.Debug("got request for updating order")
	if ok := validatePermissions([]string{pkg.UpdateOrder}, ctx); !ok {
		ctx.SetBodyString(pkg.ErrUserForbidden.Error())
		ctx.SetStatusCode(fasthttp.StatusForbidden)
		return
	}
	request, err := validateBody[pkg.OrderUpdate](ctx)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	if err = i.api.UpdateOrder(request); err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBodyString("successfully updated order")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

// deleteOrder deletes order with id specified in route
func (i *HTTPInstanceAPI) deleteOrder(ctx *fasthttp.RequestCtx) {
	i.log.Debug("got request for deleting order")
	if ok := validatePermissions([]string{pkg.DeleteOrder}, ctx); !ok {
		ctx.SetBodyString(pkg.ErrUserForbidden.Error())
		ctx.SetStatusCode(fasthttp.StatusForbidden)
		return
	}
	orderId := ctx.UserValue("id").(string)
	if orderId == "" {
		ctx.SetBodyString(pkg.ErrUnableToReadPayload.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	if err := i.api.Delete(&pkg.Order{Id: orderId}); err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	ctx.SetBodyString("deleted")
	ctx.SetStatusCode(fasthttp.StatusOK)
}
