// Package pkg provides models & provider implementations for the application
// This file contains requests definitons for http service
package pkg

import "github.com/lib/pq"

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserCreate struct {
	CustomerSpecificModel
	Username      string   `json:"username"`
	Password      string   `json:"password"`
	PasswordCheck string   `json:"passwordCheck"`
	Email         string   `json:"email"`
	Permissions   []string `json:"permissions"`
	IsActive      bool     `json:"isActive"`
}

type UserUpdate struct {
	CustomerSpecificModel
	Id          string   `json:"id"`
	Email       string   `json:"email"`
	Permissions []string `json:"permissions"`
}

type ProductCreate struct {
	CustomerSpecificModel
	Names        JSONB    `json:"names"`
	Descriptions JSONB    `json:"descriptions"`
	Categories   []string `json:"categories"`
	Prices       JSONB    `json:"prices"`
}

type ProductUpdate struct {
	CustomerSpecificModel
	Id           string         `json:"id"`
	Names        JSONB          `json:"names"`
	Descriptions JSONB          `json:"descriptions"`
	Categories   pq.StringArray `json:"categories"`
	Prices       JSONB          `json:"prices"`
}

type CategoryCreate struct {
	CustomerSpecificModel
	Name string `json:"name"`
}

type CategoryUpdate struct {
	CustomerSpecificModel
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CouponCreate struct {
	CustomerSpecificModel
	PromoCode string `json:"promoCode"`
	Amount    int    `json:"amount"`
	Unit      string `json:"unit"`
}

type CouponUpdate struct {
	CustomerSpecificModel
	Id        string `json:"id"`
	PromoCode string `json:"promoCode"`
	Amount    int    `json:"amount"`
	Unit      string `json:"unit"`
}

type OrderCreate struct {
	CustomerSpecificModel
	TotalCost  float32 `json:"totalCost"`
	Currency   string  `gorm:"size:3" json:"currency"`
	Country    string  `gorm:"size:2" json:"country"`
	Coupon     string  `json:"promoCode"`
	City       string  `json:"city"`
	PostalCode string  `json:"postalCode"`
	Address    string  `json:"address"`
	Basket     JSONB   `json:"basket"`
	TaxId      string  `json:"taxId"`
}

type OrderUpdate struct {
	CustomerSpecificModel
	Id         string  `json:"id"`
	TotalCost  float32 `json:"totalCost"`
	Currency   string  `gorm:"size:3" json:"currency"`
	Country    string  `gorm:"size:2" json:"country"`
	City       string  `json:"city"`
	PostalCode string  `json:"postalCode"`
	Address    string  `json:"address"`
	Status     string  `json:"status"`
	Basket     JSONB   `json:"basket"`
	TaxId      string  `json:"taxId"`
}

// ProductSummary is what is being stored
// as basket product in database
type ProductSummary struct {
	Price    float32 `json:"price"`
	Currency string  `json:"currency"`
	Quantity int     `json:"quantity"`
	Name     string  `json:"name"`
}
