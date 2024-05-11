package api

import (
	"github.com/jinzhu/copier"
	"github.com/pstano1/go-cart/internal/pkg"
	"go.uber.org/zap"
)

func (a *InstanceAPI) UpdateUser(request *pkg.UserUpdate) error {
	a.log.Debug("updating account info",
		zap.String("account", request.Id),
	)
	users, err := a.GetUsers(&pkg.UserFilter{
		Id: request.Id,
	})
	if err != nil {
		a.log.Error("Could not retrieve user",
			zap.Error(err),
		)
		return pkg.ErrUserNotFound
	}
	user := users[0]
	err = copier.Copy(&user, request)
	if err != nil {
		a.log.Debug(err.Error())
		return err
	}
	err = a.dbController.Update(user)
	if err != nil {
		a.log.Debug(err.Error())
		return pkg.ErrUpdatingUser
	}
	return nil
}

func (a *InstanceAPI) UpdateProduct(request *pkg.ProductUpdate) error {
	a.log.Debug("updating product info",
		zap.String("id", request.Id),
	)
	product, err := a.GetProducts(&pkg.ProductFilter{
		Id: request.Id,
	})
	if err != nil {
		a.log.Error("Could not retrieve product",
			zap.Error(err),
		)
		return pkg.ErrProductNotFound
	}
	err = copier.Copy(&product, request)
	if err != nil {
		a.log.Debug(err.Error())
		return err
	}
	err = a.dbController.Update(product)
	if err != nil {
		a.log.Debug(err.Error())
		return pkg.ErrUpdatingProduct
	}
	return nil
}
