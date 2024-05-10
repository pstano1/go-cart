package pkg

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type CustomerSpecificModel struct {
	CustomerId string `json:"customerId"`
}

type User struct {
	CustomerSpecificModel
	Id          string         `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Username    string         `gorm:"unique" json:"username"`
	Password    string         `json:"password"`
	Email       string         `json:"email"`
	IsActive    bool           `json:"isActive"`
	Permissions pq.StringArray `gorm:"type:text[]" json:"permissions"`
}

func (u *User) IsEmpty() bool {
	return u == &User{}
}

type Permission struct {
	Name string `gorm:"primarykey" json:"name"`
}

type Product struct {
	CustomerSpecificModel
	Id          string         `gorm:"primarykey" json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Categories  pq.StringArray `gorm:"type:text[]" json:"categories"`
	Prices      []Price        `gorm:"type:jsonb" json:"prices"`
}

type ProductCategory struct {
	CustomerSpecificModel
	Name string `gorm:"primarykey" json:"name"`
}

type Price struct {
	Rate     float32 `json:"rate"`
	Currency string  `json:"size:3"`
}

type Coupon struct {
	CustomerSpecificModel
	PromoCode  string         `gorm:"primarykey" json:"promoCode"`
	Categories pq.StringArray `gorm:"type:text[]" json:"categories"`
}

type Order struct {
	CustomerSpecificModel
	Id         string `gorm:"primarykey" json:"id"`
	TotalCost  int    `json:"totalCost"`
	Currency   string `gorm:"size:3" json:"currency"`
	Country    string `gorm:"size:2" json:"country"`
	City       string `json:"city"`
	PostalCode string `json:"postalCode"`
	Address    string `json:"address"`
}
