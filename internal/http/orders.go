package http

import "github.com/valyala/fasthttp"

func (i *HTTPInstanceAPI) getOrder(ctx *fasthttp.RequestCtx) {
	i.log.Debug("got request for retrieving orders")
}
