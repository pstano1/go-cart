package pkg

type SignInResponse struct {
	Username    string   `json:"username"`
	Permissions []string `json:"permissions"`
	Token       string   `json:"sessionToken"`
}

type ObjectCreateResponse struct {
	Id string `json:"id"`
}
