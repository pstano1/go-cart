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
	return res.Ok, err
}

func (a *InstanceAPI) GetProducts(filter *pkg.ProductFilter) ([]pkg.Product, error) {
	return a.dbController.GetProducts(filter)
}
