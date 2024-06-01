// Package api provides a logic for the application
// This file contains definitions of create methods
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

// CreateProduct processes request data - validates it & then creates model instance
// with given data
func (a *InstanceAPI) CreateProduct(request *pkg.ProductCreate) (*string, error) {
	a.log.Debug("creating product",
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
	for key, value := range product.Names {
		print(key)
		if !isValidNameOrDescription(key, value) {
			return nil, pkg.ErrInvalidNameKeyOrValue
		}
	}
	for key, value := range product.Descriptions {
		if !isValidNameOrDescription(key, value) {
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

func (a *InstanceAPI) CreateOrder(request *pkg.OrderCreate) (pkg.OrderCreateResponse, error) {
	a.log.Debug("creating order",
		zap.String("customer", request.CustomerId),
	)
	var res pkg.OrderCreateResponse
	var order pkg.Order
	err := copier.Copy(&order, request)
	if err != nil {
		a.log.Error("error while copying request",
			zap.Error(err),
		)
		return res, err
	}
	var total float32 = 0
	summaries := make([]pkg.ProductSummary, 0)
	for key, value := range request.Basket {
		if !isValidBasketEntry(key, value) {
			return res, pkg.ErrInvalidBasketValue
		}
		product, err := a.GetProducts(&pkg.ProductFilter{
			Id: key,
		})
		if err != nil || len(product) == 0 {
			return res, pkg.ErrInvalidBasketValue
		}
		summaryMap, ok := value.(map[string]interface{})
		if !ok {
			return res, pkg.ErrInvalidBasketValue
		}
		productSummary := &pkg.ProductSummary{
			Price:    float32(summaryMap["price"].(float64)),
			Name:     summaryMap["name"].(string),
			Currency: summaryMap["currency"].(string),
			Quantity: int(summaryMap["quantity"].(float64)),
		}
		if productSummary.Price != float32(product[0].Prices[productSummary.Currency].(float64)) {
			return res, pkg.ErrInvalidBasketValue
		}
		total += productSummary.Price * float32(productSummary.Quantity)
		coupons, err := a.GetCoupons(&pkg.CouponFilter{
			Code:       request.Coupon,
			CustomerId: request.CustomerId,
		})
		if len(coupons) != 0 && coupons[0].IsActive && err == nil {
			if coupons[0].Unit == "percentage" {
				total = total - (total * (float32(coupons[0].Amount) / 100))
			} else {
				rate, _ := a.exchangeProvider.GetExchangeRate(coupons[0].Unit, request.Currency)
				total = total - (float32(coupons[0].Amount) * rate)
			}
		}
		if order.TotalCost != total {
			return res, pkg.ErrInvalidBasketValue
		}
		summaries = append(summaries, *productSummary)
	}
	order.Id = uuid.New().String()
	order.Status = "placed"
	err = a.dbController.Create(&order)
	res.Id = order.Id
	if err != nil {
		a.log.Error("error while creating order",
			zap.Error(err),
		)
		return res, pkg.ErrCreatingOrder
	}
	url, err := a.stripeProvider.CreatePayment(summaries)
	if err != nil {
		a.log.Error("error while creating checkout session",
			zap.Error(err),
		)
		return res, pkg.ErrCreatingOrder
	}
	res.CheckoutURL = url
	// TO-DO: send confirmation e-mail
	return res, nil
}
