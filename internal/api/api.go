// Package api provides a logic for the application
// This file contains definition & some helper functions
package api

import (
	"regexp"

	"github.com/pstano1/customer-api/client"
	"github.com/pstano1/go-cart/internal/db"
	exchange "github.com/pstano1/go-cart/internal/pkg/exchangeProvider"
	"github.com/pstano1/go-cart/internal/pkg/stripeProvider"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func NewInstanceAPI(conf *APIConfig) *InstanceAPI {
	return &InstanceAPI{
		log:              conf.Logger,
		dbController:     conf.DBController,
		customerService:  conf.CustomerClient,
		exchangeProvider: conf.ExchangeProvider,
		stripeProvider:   conf.StripeProvider,
		secretKey:        conf.SecretKey,
	}
}

type APIConfig struct {
	Logger           *zap.Logger
	DBController     db.IDBController
	CustomerClient   client.ICustomerServiceClient
	ExchangeProvider exchange.IExchangeProvider
	StripeProvider   stripeProvider.IStripeProvider
	SecretKey        string
}

type InstanceAPI struct {
	log              *zap.Logger
	dbController     db.IDBController
	customerService  client.ICustomerServiceClient
	exchangeProvider exchange.IExchangeProvider
	stripeProvider   stripeProvider.IStripeProvider
	secretKey        string
}

// GetHash generates hash from given password
func getHash(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// isValidPrice checks generic interface{} if it's a valid price
// following the pattern of being a string, number pair
// where string is 3 characters long
func isValidPrice(key string, value interface{}) bool {
	regex := regexp.MustCompile(`^[A-Z]{3}$`)
	if !regex.MatchString(key) {
		return false
	}

	switch value.(type) {
	case float64:
		return true
	case int:
		return true
	case float32:
		return true
	default:
		return false
	}
}

// isValidNameOrDescription checks generic interface{} if it's a valid name or description
// following the pattern of being a string, string pair
// where first string (a key) is 2 characters long
func isValidNameOrDescription(key string, value interface{}) bool {
	regex := regexp.MustCompile(`^[A-Z]{2}$`)
	if !regex.MatchString(key) {
		return false
	}

	if desc, ok := value.(string); ok {
		return len(desc) > 0
	}

	return false
}

// isValidProductDescription is a helper function for `isValidbasketEntry`
// checks if field is valid by checking for it matching criteria of type & length
func isValidBasketProductDescription(key string, value interface{}) bool {
	regex := regexp.MustCompile(`^[A-Z]{3}$`)
	switch val := value.(type) {
	case int:
		return key == "price" || key == "quantity"
	case float32, float64:
		return key == "price" || key == "quantity"
	case string:
		return key == "name" || (key == "currency" && regex.MatchString(val))
	default:
		return false
	}
}

// checks generic interface{} if it's a valid `pkg.ProductSummary` implementation
// checks if interface{} contains all erquired fields & if type and length are valid
func isValidBasketEntry(key string, value interface{}) bool {
	switch val := value.(type) {
	case map[string]interface{}:
		requiredKeys := map[string]bool{
			"price":    true,
			"currency": true,
			"quantity": true,
			"name":     true,
		}
		for k, v := range val {
			if !isValidBasketProductDescription(k, v) {
				return false
			}
		}
		for k := range requiredKeys {
			if _, found := val[k]; !found {
				return false
			}
		}
		return true
	default:
		return false
	}
}

// FetchExchangeRates fetches all possible exhange rates from NBP
func (a *InstanceAPI) FetchExchangeRates() error {
	for _, name := range []string{"a", "b", "c"} {
		if err := a.exchangeProvider.FetchNBPTable(name); err != nil {
			return err
		}
	}
	return nil
}
