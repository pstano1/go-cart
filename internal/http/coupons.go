// Package http provides server implementation for the application
// This file includes /coupon routes handlers
package http

import (
	"encoding/json"

	"github.com/pstano1/go-cart/internal/pkg"
	"github.com/valyala/fasthttp"
)

// getCoupon retrieves coupon(s) based on provided query params
// (or lack of) and returns slice of pkg.Coupon to user
func (i *HTTPInstanceAPI) getCoupon(ctx *fasthttp.RequestCtx) {
	i.log.Debug("got request for retrieving coupons")
	filter, err := validateFilter[pkg.CouponFilter](ctx)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	coupons, err := i.api.GetCoupons(&filter)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(coupons)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBody(response)
	ctx.SetStatusCode(fasthttp.StatusOK)
}

// createCoupon handles coupon creation based on request's body
// & returns object of type pkg.ObjectCreateResponse
func (i *HTTPInstanceAPI) createCoupon(ctx *fasthttp.RequestCtx) {
	i.log.Debug("got request for creating coupon")
	if ok := validatePermissions([]string{pkg.CreateCoupon}, ctx); !ok {
		ctx.SetBodyString(pkg.ErrUserForbidden.Error())
		ctx.SetStatusCode(fasthttp.StatusForbidden)
		return
	}
	request, err := validateBody[pkg.CouponCreate](ctx)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	couponId, err := i.api.CreateCoupon(request)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(&pkg.ObjectCreateResponse{
		Id: *couponId,
	})
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBody(response)
	ctx.SetStatusCode(fasthttp.StatusCreated)
}

// updateCoupon handles coupon update based on request's body
func (i *HTTPInstanceAPI) updateCoupon(ctx *fasthttp.RequestCtx) {
	i.log.Debug("got request for updating coupon")
	if ok := validatePermissions([]string{pkg.UpdateCoupon}, ctx); !ok {
		ctx.SetBodyString(pkg.ErrUserForbidden.Error())
		ctx.SetStatusCode(fasthttp.StatusForbidden)
		return
	}
	request, err := validateBody[pkg.CouponUpdate](ctx)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	if err = i.api.UpdateCoupon(request); err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBodyString("successfully updated coupon")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

// deleteCoupon deletes coupon with id specified in route
func (i *HTTPInstanceAPI) deleteCoupon(ctx *fasthttp.RequestCtx) {
	i.log.Debug("got request for deleting coupon")
	if ok := validatePermissions([]string{pkg.DeleteCoupon}, ctx); !ok {
		ctx.SetBodyString(pkg.ErrUserForbidden.Error())
		ctx.SetStatusCode(fasthttp.StatusForbidden)
		return
	}
	couponId := ctx.UserValue("id").(string)
	if couponId == "" {
		ctx.SetBodyString(pkg.ErrUnableToReadPayload.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	if err := i.api.Delete(&pkg.Coupon{Id: couponId}); err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	ctx.SetBodyString("deleted")
	ctx.SetStatusCode(fasthttp.StatusOK)
}
