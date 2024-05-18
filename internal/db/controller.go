package db

import (
	"database/sql"

	"github.com/pstano1/go-cart/internal/pkg"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IDBController interface {
	Create(model interface{}) error
	Update(model interface{}) error
	Delete(model interface{}, conditions ...interface{}) error

	GetUsers(filter *pkg.UserFilter) ([]pkg.User, error)
	GetProducts(filter *pkg.ProductFilter) ([]pkg.Product, error)
	GetCategories(filter *pkg.CategoryFilter) ([]string, error)
	GetCoupons(filter *pkg.CouponFilter) ([]pkg.Coupon, error)
	GetOrders(filter *pkg.OrderFilter) ([]pkg.Order, error)
	GetPermissions() ([]string, error)
}

type DBController struct {
	db     *sql.DB
	gormDB *gorm.DB
	log    *zap.Logger
}

func NewDBController(gormDB *gorm.DB, log *zap.Logger) IDBController {
	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Error("could not initialize database controller",
			zap.Error(err),
		)
		return nil
	}
	return &DBController{
		log:    log,
		db:     sqlDB,
		gormDB: gormDB,
	}
}

func (d *DBController) Create(model interface{}) error {
	return d.gormDB.Create(model).Error
}

func (d *DBController) Update(model interface{}) error {
	return d.gormDB.Save(model).Error
}

func (d *DBController) Delete(model interface{}, conditions ...interface{}) error {
	return d.gormDB.Delete(model, conditions...).Error
}
