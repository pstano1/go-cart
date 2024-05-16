package api

import (
	"github.com/bcmills/unsafeslice"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/pstano1/go-cart/internal/pkg"
	"go.uber.org/zap"
)

func (a *InstanceAPI) CreateUser(request *pkg.UserCreate) (*string, error) {
	a.log.Debug("creating account",
		zap.String("username", request.Username),
		zap.String("customer", request.CustomerId),
	)
	var user pkg.User
	err := copier.Copy(&user, request)
	if err != nil {
		a.log.Error("error while copying request",
			zap.Error(err),
		)
		return nil, err
	}
	user.Password = request.PasswordCheck
	users, err := a.dbController.GetUsers(&pkg.UserFilter{
		Username: user.Username,
	})
	if err != nil {
		a.log.Error("error while retreving users",
			zap.Error(err),
		)
		return nil, pkg.ErrRetrievingUsers
	}
	if len(users) != 0 {
		return nil, pkg.ErrUserAlreadyExists
	}
	hashPassword, err := getHash(unsafeslice.OfString(user.Password))
	if err != nil {
		a.log.Error("error while hashing password",
			zap.Error(err),
		)
		return nil, err
	}
	user.Id = uuid.New().String()
	user.Password = hashPassword
	err = a.dbController.Create(&user)
	if err != nil {
		a.log.Error("error while creating account",
			zap.Error(err),
		)
		return nil, pkg.ErrCreatingUser
	}
	return &user.Id, nil
}

func (a *InstanceAPI) CreateProduct(request *pkg.ProductCreate) (*string, error) {
	a.log.Debug("creating product",
		zap.String("name", request.Name),
		zap.String("customer", request.CustomerId),
	)
	var product pkg.Product
	err := copier.Copy(&product, request)
	if err != nil {
		a.log.Error("error while copying request",
			zap.Error(err),
		)
		return nil, err
	}
	for key, value := range product.Descriptions {
		if !isValidDescription(key, value) {
			return nil, pkg.ErrInvalidDescriptionKeyOrValue
		}
	}
	// TODO: retrieve categories and check if they exist
	for key, value := range product.Prices {
		if !isValidPrice(key, value) {
			return nil, pkg.ErrInvalidPriceKeyOrValue
		}
	}
	product.Id = uuid.New().String()
	err = a.dbController.Create(&product)
	if err != nil {
		a.log.Error("error while creating product",
			zap.String("name", request.Name),
			zap.Error(err),
		)
		return nil, pkg.ErrCreatingProduct
	}
	return &product.Id, nil
}

func (a *InstanceAPI) CreateCategory(request *pkg.CategoryCreate) (*string, error) {
	a.log.Debug("creating category",
		zap.String("name", request.Name),
		zap.String("customer", request.CustomerId),
	)
	var category pkg.ProductCategory
	err := copier.Copy(&category, request)
	if err != nil {
		a.log.Error("error while copying request",
			zap.Error(err),
		)
		return nil, err
	}
	category.Id = uuid.New().String()
	err = a.dbController.Create(&category)
	if err != nil {
		a.log.Error("error while creating category",
			zap.String("name", request.Name),
			zap.Error(err),
		)
		return nil, pkg.ErrCreatingCategory
	}
	return &category.Id, nil
}

func (a *InstanceAPI) CreateCoupon(request *pkg.CouponCreate) (*string, error) {
	a.log.Debug("creating coupon",
		zap.String("code", request.PromoCode),
		zap.String("customer", request.CustomerId),
	)
	var coupon pkg.Coupon
	err := copier.Copy(&coupon, request)
	if err != nil {
		a.log.Error("error while copying request",
			zap.Error(err),
		)
		return nil, err
	}
	coupon.Id = uuid.New().String()
	coupon.IsActive = true
	err = a.dbController.Create(&coupon)
	if err != nil {
		a.log.Error("error while creating coupon",
			zap.String("code", request.PromoCode),
			zap.Error(err),
		)
		return nil, pkg.ErrCreatingCategory
	}
	return &coupon.Id, nil
}
