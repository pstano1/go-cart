package api

import (
	"github.com/pstano1/go-cart/internal/db"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func NewInstanceAPI(conf *APIConfig) *InstanceAPI {
	return &InstanceAPI{
		log:          conf.Logger,
		dbController: conf.DBCOntroller,
		secretKey:    conf.SecretKey,
	}
}

type APIConfig struct {
	Logger       *zap.Logger
	DBCOntroller db.IDBController
	SecretKey    string
}

type InstanceAPI struct {
	log          *zap.Logger
	dbController db.IDBController
	secretKey    string
}

func getHash(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}