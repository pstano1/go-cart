package db

import (
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
	var product pkg.Product
	products := make([]pkg.Product, 0)
	gormQuery := d.gormDB.Table("products").Select(`
		products.id,
		products.customer_id,
		products.name,
		products.descriptions,
		products.categories,
		products.prices
	`)
	if filter.Id != "" {
		gormQuery = gormQuery.Where("products.id = ?", filter.Id)
	}
	if filter.CustomerId != "" {
		gormQuery = gormQuery.Where("products.customer_id = ?", filter.CustomerId)
	}
	if len(filter.Categories) > 0 {
		// TODO: make it select all products with at least one matching category
	}
	gormQuery = gormQuery.Where("products.deleted_at is null")
	rows, err := gormQuery.Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err = rows.Scan(
			&product.Id,
			&product.CustomerId,
			&product.Name,
			&product.Descriptions,
			&product.Categories,
			&product.Prices,
		); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
