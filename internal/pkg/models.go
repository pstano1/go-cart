package pkg

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type User struct {
	Id          string         `gorm:"primarykey" json:"id"`
	CustomerId  string         `json:"customerId"`
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
