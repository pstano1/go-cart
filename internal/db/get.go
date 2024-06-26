package db

import (
	"fmt"
	"strings"

	"github.com/pstano1/go-cart/internal/pkg"
)

func (d *DBController) GetUsers(filter *pkg.UserFilter) ([]pkg.User, error) {
	var user pkg.User
	users := make([]pkg.User, 0)
	gormQuery := d.gormDB.Table("users").Select(`
		users.id,
		users.customer_id,
		users.username,
		users.password,
		users.email,
		users.is_active,
		users.permissions
	`)
	if filter.Username != "" {
		gormQuery = gormQuery.Where("users.username = ?", filter.Username)
	}
	if filter.Id != "" {
		gormQuery = gormQuery.Where("users.id = ?", filter.Id)
	}
	if filter.CustomerId != "" {
		gormQuery = gormQuery.Where("users.customer_id = ?", filter.CustomerId)
	}
	gormQuery = gormQuery.Where("users.deleted_at is null")
	rows, err := gormQuery.Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err = rows.Scan(
			&user.Id,
			&user.CustomerId,
			&user.Username,
			&user.Password,
			&user.Email,
			&user.IsActive,
			&user.Permissions,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (d *DBController) GetProducts(filter *pkg.ProductFilter) ([]pkg.Product, error) {
	products := make([]pkg.Product, 0)
	gormQuery := d.gormDB.Table("products").Select(`
		products.id,
		products.customer_id,
		products.names,
		products.descriptions,
		products.categories,
		products.prices,
		products.price_history
	`)
	if filter.Id != "" {
		gormQuery = gormQuery.Where("products.id = ?", filter.Id)
	}
	if filter.CustomerId != "" {
		gormQuery = gormQuery.Where("products.customer_id = ?", filter.CustomerId)
	}
	if len(filter.Categories) > 0 && filter.Categories[0] != "" {
		var conditions []string
		for _, category := range filter.Categories {
			conditions = append(conditions, fmt.Sprintf("'%s' = ANY(products.categories)", category))
		}
		gormQuery = gormQuery.Where("(" + strings.Join(conditions, " OR ") + ")")
	}
	gormQuery = gormQuery.Where("products.deleted_at is null")
	rows, err := gormQuery.Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var product pkg.Product
		if err = rows.Scan(
			&product.Id,
			&product.CustomerId,
			&product.Names,
			&product.Descriptions,
			&product.Categories,
			&product.Prices,
			&product.PriceHistory,
		); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (d *DBController) GetCategories(filter *pkg.CategoryFilter) ([]pkg.ProductCategory, error) {
	categories := make([]pkg.ProductCategory, 0)
	gormQuery := d.gormDB.Table("product_categories").Select(`
		product_categories.name,
		product_categories.id
	`)
	if filter.Id != "" {
		gormQuery = gormQuery.Where("product_categories.id = ?", filter.Id)
	}
	if filter.CustomerId != "" {
		gormQuery = gormQuery.Where("product_categories.customer_id = ?", filter.CustomerId)
	}
	gormQuery = gormQuery.Where("product_categories.deleted_at is null")
	rows, err := gormQuery.Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var category pkg.ProductCategory
		if err = rows.Scan(
			&category.Name,
			&category.Id,
		); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (d *DBController) GetCoupons(filter *pkg.CouponFilter) ([]pkg.Coupon, error) {
	coupons := make([]pkg.Coupon, 0)
	gormQuery := d.gormDB.Table("coupons").Select(`
		coupons.id,
		coupons.customer_id,
		coupons.promo_code,
		coupons.amount,
		coupons.unit,
		coupons.is_active
	`)
	if filter.Id != "" {
		gormQuery = gormQuery.Where("coupons.id = ?", filter.Id)
	}
	if filter.Code != "" {
		gormQuery = gormQuery.Where("coupons.promo_code = ?", filter.Code)
	}
	if filter.CustomerId != "" {
		gormQuery = gormQuery.Where("coupons.customer_id = ?", filter.CustomerId)
	}
	gormQuery = gormQuery.Where("coupons.deleted_at is null")
	rows, err := gormQuery.Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var coupon pkg.Coupon
		if err = rows.Scan(
			&coupon.Id,
			&coupon.CustomerId,
			&coupon.PromoCode,
			&coupon.Amount,
			&coupon.Unit,
			&coupon.IsActive,
		); err != nil {
			return nil, err
		}
		coupons = append(coupons, coupon)
	}
	return coupons, nil
}

func (d *DBController) GetOrders(filter *pkg.OrderFilter) ([]pkg.Order, error) {
	orders := make([]pkg.Order, 0)
	gormQuery := d.gormDB.Table("orders").Select(`
		orders.id,
		orders.customer_id,
		orders.total_cost,
		orders.currency,
		orders.city,
		orders.postal_code,
		orders.address,
		orders.country,
		orders.basket,
		orders.status
	`)
	if filter.Id != "" {
		gormQuery = gormQuery.Where("orders.id = ?", filter.Id)
	}
	if filter.CustomerId != "" {
		gormQuery = gormQuery.Where("orders.customer_id = ?", filter.CustomerId)
	}
	gormQuery = gormQuery.Where("orders.deleted_at is null")
	rows, err := gormQuery.Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var order pkg.Order
		if err = rows.Scan(
			&order.Id,
			&order.CustomerId,
			&order.TotalCost,
			&order.Currency,
			&order.City,
			&order.PostalCode,
			&order.Address,
			&order.Country,
			&order.Basket,
			&order.Status,
		); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (d *DBController) GetPermissions() ([]string, error) {
	var permissions []string
	rows, err := d.gormDB.Model(pkg.Permission{}).Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var name string
		if err = rows.Scan(
			&name,
		); err != nil {
			return nil, err
		}
		permissions = append(permissions, name)
	}
	return permissions, nil
}
