// Package pkg provides models & provider implementations for the application
// This file contains reponses definitions for the http service
package pkg

type SignInResponse struct {
	Username    string   `json:"username"`
	Permissions []string `json:"permissions"`
	Token       string   `json:"sessionToken"`
}

type ObjectCreateResponse struct {
	Id string `json:"id"`
}

type OrderCreateResponse struct {
	Id          string `json:"id"`
	CheckoutURL string `json:"checkoutURL"`
}
