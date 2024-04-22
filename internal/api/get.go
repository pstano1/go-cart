package api

import "github.com/pstano1/go-cart/internal/pkg"

func (a *InstanceAPI) GetUsers(filter *pkg.UserFilter) ([]pkg.User, error) {
	return a.dbController.GetUsers(filter)
}
