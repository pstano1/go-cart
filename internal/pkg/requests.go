package pkg

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserCreate struct {
	CustomerId    string   `json:"customerId"`
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
