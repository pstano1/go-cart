package pkg

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONB) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	data, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(data, &j)
}

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
	Id           string         `gorm:"primarykey" json:"id"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	Names        JSONB          `gorm:"type:jsonb" json:"names"`
	Descriptions JSONB          `gorm:"type:jsonb" json:"descriptions"`
	Categories   pq.StringArray `gorm:"type:text[]" json:"categories"`
	Prices       JSONB          `gorm:"type:jsonb" json:"prices"`
	PriceHistory JSONB          `gotm:"type:jsonb" json:"priceHistory"`
}

type ProductCategory struct {
	CustomerSpecificModel
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Name      string         `json:"name"`
	Id        string         `gorm:"primarykey" json:"id"`
}

type Coupon struct {
	CustomerSpecificModel
	Id         string         `gorm:"primarykey" json:"id"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	PromoCode  string         `json:"promoCode"`
	Amount     int            `json:"amount"`
	Unit       string         `json:"unit"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	Categories pq.StringArray `gorm:"type:text[]" json:"categories"`
	IsActive   bool           `json:"isActive"`
}

type Order struct {
	CustomerSpecificModel
	Id         string         `gorm:"primarykey" json:"id"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	TotalCost  float32        `json:"totalCost"`
	Currency   string         `gorm:"size:3" json:"currency"`
	Country    string         `gorm:"size:2" json:"country"`
	City       string         `json:"city"`
	PostalCode string         `json:"postalCode"`
	Address    string         `json:"address"`
	Status     string         `json:"status"`
	Basket     JSONB          `gorm:"type:jsonb" json:"basket"`
	TaxId      string         `json:"taxId"`
}
