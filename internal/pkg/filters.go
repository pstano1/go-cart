package pkg

import (
	"strings"

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

type ProductFilter struct {
	Id         string
	CustomerId string
	Categories []string
}

func (f ProductFilter) Populate(ctx *fasthttp.RequestCtx) Filter {
	args := ctx.QueryArgs()

	return &ProductFilter{
		Id:         string(args.Peek("id")),
		CustomerId: string(args.Peek("customerId")),
		Categories: strings.Split(string(args.Peek("categories")), ","),
	}
}

type CategoryFilter struct {
	Id         string
	CustomerId string
}

func (f CategoryFilter) Populate(ctx *fasthttp.RequestCtx) Filter {
	args := ctx.QueryArgs()

	return &ProductFilter{
		Id:         string(args.Peek("id")),
		CustomerId: string(args.Peek("customerId")),
	}
}

type CouponFilter struct {
	Id         string
	CustomerId string
}

func (f CouponFilter) Populate(ctx *fasthttp.RequestCtx) Filter {
	args := ctx.QueryArgs()

	return &ProductFilter{
		Id:         string(args.Peek("id")),
		CustomerId: string(args.Peek("customerId")),
	}
}

type OrderFilter struct {
	Id         string
	CustomerId string
}

func (f OrderFilter) Populate(ctx *fasthttp.RequestCtx) Filter {
	args := ctx.QueryArgs()

	return &ProductFilter{
		Id:         string(args.Peek("id")),
		CustomerId: string(args.Peek("customerId")),
	}
}
