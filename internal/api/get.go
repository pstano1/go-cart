// Package api provides a logic for the application
// This file contains definitions of get methods
package api

import (
	pb "github.com/pstano1/customer-api/client/proto"
	"github.com/pstano1/go-cart/internal/pkg"
)

func (a *InstanceAPI) GetUsers(filter *pkg.UserFilter) ([]pkg.User, error) {
	return a.dbController.GetUsers(filter)
}

func (a *InstanceAPI) ExchangeTagForId(tag string) (*pb.ExchangeTagForIdResponse, error) {
	return a.customerService.ExchangeTagForId(tag)
}

func (a *InstanceAPI) ValidateCustomerId(id string) (bool, error) {
	res, err := a.customerService.ValidateId(id)
	if err != nil {
		return false, err
	}
	return res.Ok, err
}

func (a *InstanceAPI) GetProducts(filter *pkg.ProductFilter) ([]pkg.Product, error) {
	return a.dbController.GetProducts(filter)
}

func (a *InstanceAPI) GetCategories(filter *pkg.CategoryFilter) ([]pkg.ProductCategory, error) {
	return a.dbController.GetCategories(filter)
}

func (a *InstanceAPI) GetCoupons(filter *pkg.CouponFilter) ([]pkg.Coupon, error) {
	return a.dbController.GetCoupons(filter)
}

func (a *InstanceAPI) GetOrders(filter *pkg.OrderFilter) ([]pkg.Order, error) {
	return a.dbController.GetOrders(filter)
}

func (a *InstanceAPI) GetPermissions() ([]string, error) {
	return a.dbController.GetPermissions()
}
