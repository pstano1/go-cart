package pkg

import (
	"github.com/valyala/fasthttp"
)

type Filter interface {
	Populate(ctx *fasthttp.RequestCtx) Filter
}

type UserFilter struct {
	Id         string
	Username   string
	CustomerId string
}

func (f UserFilter) Populate(ctx *fasthttp.RequestCtx) Filter {
	args := ctx.QueryArgs()

	return &UserFilter{
		Id:         string(args.Peek("id")),
		Username:   string(args.Peek("username")),
		CustomerId: string(args.Peek("customerId")),
	}
}
