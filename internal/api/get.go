package api

import (
	pb "github.com/pstano1/customer-api/client/proto"
	"github.com/pstano1/go-cart/internal/pkg"
)

func (a *InstanceAPI) GetUsers(filter *pkg.UserFilter) ([]pkg.User, error) {
	return a.dbController.GetUsers(filter)
}

func (a *InstanceAPI) ExhchangeTagForId(tag string) (*pb.ExchangeTagForIdResponse, error) {
	return a.customerService.ExchangeTagForId(tag)
}
