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
