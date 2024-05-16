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
	products, err := a.GetProducts(&pkg.ProductFilter{
		Id: request.Id,
	})
	if err != nil {
		a.log.Error("Could not retrieve product",
			zap.Error(err),
		)
		return pkg.ErrProductNotFound
	}
	product := products[0]
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

func (a *InstanceAPI) UpdateCategory(request *pkg.CategoryUpdate) error {
	a.log.Debug("updating category info",
		zap.String("id", request.Id),
	)
	err := a.dbController.Update(&pkg.ProductCategory{
		Id:                    request.Id,
		Name:                  request.Name,
		CustomerSpecificModel: request.CustomerSpecificModel,
	})
	if err != nil {
		a.log.Debug(err.Error())
		return pkg.ErrUpdatingCategory
	}
	return nil
}

func (a *InstanceAPI) UpdateCoupon(request *pkg.CouponUpdate) error {
	a.log.Debug("updating coupon info",
		zap.String("id", request.Id),
	)
	coupons, err := a.GetCoupons(&pkg.CouponFilter{
		Id: request.Id,
	})
	if err != nil {
		a.log.Error("Could not retrieve coupon",
			zap.Error(err),
		)
		return pkg.ErrCouponNotFound
	}
	coupon := coupons[0]
	err = copier.Copy(&coupon, request)
	if err != nil {
		a.log.Debug(err.Error())
		return err
	}
	err = a.dbController.Update(coupon)
	if err != nil {
		a.log.Debug(err.Error())
		return pkg.ErrUpdatingCoupon
	}
	return nil
}
