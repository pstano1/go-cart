package main

import (
	"fmt"
	"log"

	"github.com/pstano1/customer-api/client"
	"github.com/pstano1/go-cart/internal/api"
	"github.com/pstano1/go-cart/internal/db"
	"github.com/pstano1/go-cart/internal/http"
	exchange "github.com/pstano1/go-cart/internal/pkg/exchangeProvider"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	confOptBindPort               = "BIND_PORT"
	confOptDatabaseHost           = "DATABASE_HOST"
	confOptDatabaseName           = "DATABASE_NAME"
	confOptDatabaseUsername       = "DATABASE_USERNAME"
	confOptDatabasePassword       = "DATABASE_PASSWORD"
	confOptDatabasePort           = "DATABASE_PORT"
	confOptSecretKey              = "SECRET_KEY"
	confOptCustomerServiceAddress = "CUSTOMER_SERVICE_ADDRESS"
)

func main() {
	logger := zap.Must(zap.NewDevelopment())
	defer logger.Sync()
	server := createServerFromConfig(logger)

	go server.GetExchangeRates()
	go server.Run()

	select {}
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

	conn, err := grpc.Dial(viper.GetString(confOptCustomerServiceAddress), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("could not establish connection",
			zap.Error(err),
		)
	}
	customerClient := client.NewCustomerService(conn, logger)

	exchangeProvider := exchange.NewProvider(logger)

	instanceAPI := api.NewInstanceAPI(&api.APIConfig{
		Logger:           logger,
		DBController:     dbController,
		CustomerClient:   customerClient,
		ExchangeProvider: exchangeProvider,
		SecretKey:        viper.GetString(confOptSecretKey),
	})

	return http.NewHTTPInstanceAPI(&http.HTTPConfig{
		Logger:   logger,
		BindPath: fmt.Sprintf(":%s", viper.GetString(confOptBindPort)),
		API:      instanceAPI,
	})
}
