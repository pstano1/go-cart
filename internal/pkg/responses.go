package pkg

type SignInResponse struct {
	Username string `json:"username"`
	Token    string `json:"sessionToken"`
}
