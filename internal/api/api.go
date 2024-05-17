package api

import (
	"regexp"

	"github.com/pstano1/customer-api/client"
	"github.com/pstano1/go-cart/internal/db"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func NewInstanceAPI(conf *APIConfig) *InstanceAPI {
	return &InstanceAPI{
		log:             conf.Logger,
		dbController:    conf.DBController,
		customerService: conf.CustomerClient,
		secretKey:       conf.SecretKey,
	}
}

type APIConfig struct {
	Logger         *zap.Logger
	DBController   db.IDBController
	CustomerClient client.ICustomerServiceClient
	SecretKey      string
}

type InstanceAPI struct {
	log             *zap.Logger
	dbController    db.IDBController
	customerService client.ICustomerServiceClient
	secretKey       string
}

func getHash(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

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

func isValidDescription(key string, value interface{}) bool {
	regex := regexp.MustCompile(`^[A-Z]{2}$`)
	if !regex.MatchString(key) {
		return false
	}

	if desc, ok := value.(string); ok {
		return len(desc) > 0
	}

	return false
}

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
