package http

import (
	"encoding/json"

	"github.com/pstano1/go-cart/internal/pkg"
	"github.com/valyala/fasthttp"
)

// getProduct retrieves product(s) based on provided query params
// (or lack of) and returns slice of pkg.Product to user
func (i *HTTPInstanceAPI) getProduct(ctx *fasthttp.RequestCtx) {
	i.log.Debug("got request for retrieving products")
	filter, err := validateFilter[pkg.ProductFilter](ctx)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	products, err := i.api.GetProducts(&filter)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(products)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBody(response)
	ctx.SetStatusCode(fasthttp.StatusOK)
}

// createProduct handles product creation based on request's body
// & returns object of type pkg.ObjectCreateResponse
func (i *HTTPInstanceAPI) createProduct(ctx *fasthttp.RequestCtx) {
	i.log.Debug("got request for product creation")
	request, err := validateBody[pkg.ProductCreate](ctx)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	productId, err := i.api.CreateProduct(request)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(&pkg.ObjectCreateResponse{
		Id: *productId,
	})
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBody(response)
	ctx.SetStatusCode(fasthttp.StatusCreated)
}

// updateProduct handles product update based on request's body
func (i *HTTPInstanceAPI) updateProduct(ctx *fasthttp.RequestCtx) {
	i.log.Debug("got request for product update")
	request, err := validateBody[pkg.ProductUpdate](ctx)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	if err = i.api.UpdateProduct(request); err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBodyString("successfully updated product")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

// deleteProduct deletes product with id specified in route
func (i *HTTPInstanceAPI) deleteProduct(ctx *fasthttp.RequestCtx) {
	i.log.Debug("got request for product deletion")
	productId := ctx.UserValue("id").(string)
	if productId == "" {
		ctx.SetBodyString(pkg.ErrUnableToReadPayload.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	if err := i.api.Delete(&pkg.Product{Id: productId}); err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	ctx.SetBodyString("deleted")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

// getCategory retrieves coupon(s) based on provided query params
// (or lack of) and returns slice of pkg.ProductCategory to user
func (i *HTTPInstanceAPI) getCategory(ctx *fasthttp.RequestCtx) {
	i.log.Debug("got request for retrieving categories")
	filter, err := validateFilter[pkg.CategoryFilter](ctx)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	categories, err := i.api.GetCategories(&filter)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(categories)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBody(response)
	ctx.SetStatusCode(fasthttp.StatusOK)
}

// createCreate handles category creation based on request's body
// & returns object of type pkg.ObjectCreateResponse
func (i *HTTPInstanceAPI) createCategory(ctx *fasthttp.RequestCtx) {
	i.log.Debug("got request for creating category")
	request, err := validateBody[pkg.CategoryCreate](ctx)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	categoryId, err := i.api.CreateCategory(request)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(&pkg.ObjectCreateResponse{
		Id: *categoryId,
	})
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBody(response)
	ctx.SetStatusCode(fasthttp.StatusCreated)
}

// updateCategory handles category update based on request's body
func (i *HTTPInstanceAPI) updateCategory(ctx *fasthttp.RequestCtx) {
	i.log.Debug("got request for category update")
	request, err := validateBody[pkg.CategoryUpdate](ctx)
	if err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	if err = i.api.UpdateCategory(request); err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBodyString("successfully updated category")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

// deleteCategory deletes category with id specified in route
func (i *HTTPInstanceAPI) deleteCategory(ctx *fasthttp.RequestCtx) {
	i.log.Debug("got request for category deletion")
	categoryId := ctx.UserValue("id").(string)
	if categoryId == "" {
		ctx.SetBodyString(pkg.ErrUnableToReadPayload.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	if err := i.api.Delete(&pkg.ProductCategory{Id: categoryId}); err != nil {
		ctx.SetBodyString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	ctx.SetBodyString("deleted")
	ctx.SetStatusCode(fasthttp.StatusOK)
}
