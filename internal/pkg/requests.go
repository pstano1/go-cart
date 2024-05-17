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
	Id          string   `json:"id"`
	Email       string   `json:"email"`
	Permissions []string `json:"permissions"`
}

type ProductCreate struct {
	CustomerSpecificModel
	Name         string   `json:"name"`
	Descriptions JSONB    `json:"descriptions"`
	Categories   []string `json:"categories"`
	Prices       JSONB    `json:"prices"`
}

type ProductUpdate struct {
	CustomerSpecificModel
	Id           string         `json:"id"`
	Name         string         `json:"name"`
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
}

type CouponUpdate struct {
	CustomerSpecificModel
	Id        string `json:"id"`
	PromoCode string `json:"promoCode"`
	Amount    int    `json:"amount"`
}

type OrderCreate struct {
	CustomerSpecificModel
	TotalCost  float32 `json:"totalCost"`
	Currency   string  `gorm:"size:3" json:"currency"`
	Country    string  `gorm:"size:2" json:"country"`
	City       string  `json:"city"`
	PostalCode string  `json:"postalCode"`
	Address    string  `json:"address"`
	Basket     JSONB   `json:"basket"`
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
}
