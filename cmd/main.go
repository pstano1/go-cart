package main

import (
	"fmt"
	"log"

	"github.com/pstano1/go-cart/internal/api"
	"github.com/pstano1/go-cart/internal/db"
	"github.com/pstano1/go-cart/internal/http"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	confOptBindPort         = "BIND_PORT"
	confOptDatabaseHost     = "DATABASE_HOST"
	confOptDatabaseName     = "DATABASE_NAME"
	confOptDatabaseUsername = "DATABASE_USERNAME"
	confOptDatabasePassword = "DATABASE_PASSWORD"
	confOptDatabasePort     = "DATABASE_PORT"
	confOptSecretKey        = "SECRET_KEY"
)

func main() {
	logger := zap.Must(zap.NewProduction())
	defer logger.Sync()
	server := createServerFromConfig(logger)
	server.Run()
}

func createServerFromConfig(logger *zap.Logger) *http.HTTPInstanceAPI {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("could not load config file: %v", err)
		return nil
	}

	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=require",
		viper.GetString(confOptDatabaseUsername),
		viper.GetString(confOptDatabasePassword),
		viper.GetString(confOptDatabaseHost),
		viper.GetString(confOptDatabasePort),
		viper.GetString(confOptDatabaseName),
	)

	gormDB, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("error occured while connecting to database %v", err)
		return nil
	}

	dbController := db.NewDBController(
		gormDB,
		logger,
	)

	instanceAPI := api.NewInstanceAPI(&api.APIConfig{
		Logger:       logger,
		DBCOntroller: dbController,
		SecretKey:    viper.GetString(confOptSecretKey),
	})

	return http.NewHTTPInstanceAPI(&http.HTTPConfig{
		Logger:   logger,
		BindPath: fmt.Sprintf(":%s", viper.GetString(confOptBindPort)),
		API:      instanceAPI,
	})
}