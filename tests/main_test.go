package tests

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/pstano1/customer-api/client"
	"github.com/pstano1/go-cart/internal/api"
	"github.com/pstano1/go-cart/internal/db"
	"github.com/pstano1/go-cart/internal/pkg"
	exchange "github.com/pstano1/go-cart/internal/pkg/exchangeProvider"
	"github.com/pstano1/go-cart/internal/pkg/stripeProvider"
	"github.com/spf13/viper"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var API *api.InstanceAPI

const (
	confOptBindPort               = "BIND_PORT"
	confOptSecretKey              = "SECRET_KEY"
	confOptCustomerServiceAddress = "CUSTOMER_SERVICE_ADDRESS"
	confOptStripeKey              = "STRIPE_KEY"
)

const (
	dbName     = "testDB"
	dbUser     = "testDBUser"
	dbPassword = "strong-password"
)

func setupTestEnvironment() {
	ctx := context.Background()
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("could not load config file: %v", err)
	}

	req := testcontainers.ContainerRequest{
		Image:        "postgres:16-alpine",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     dbUser,
			"POSTGRES_PASSWORD": dbPassword,
			"POSTGRES_DB":       dbName,
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}
	postgresContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("failed to start container: %s", err)
	}

	port, err := postgresContainer.MappedPort(ctx, "5432")
	if err != nil {
		log.Fatalf("failed to get container's port: %s", err)
	}

	connURI := fmt.Sprintf("host=localhost port=%s user=%s password=%s dbname=%s sslmode=disable", port.Port(), dbUser, dbPassword, dbName)
	logger := zap.Must(zap.NewDevelopment())
	defer logger.Sync()

	gormDB, err := gorm.Open(postgres.Open(connURI), &gorm.Config{})
	if err != nil {
		log.Fatalf("error occured while connecting to database %v", err)
	}
	err = gormDB.AutoMigrate(
		&pkg.User{},
		&pkg.Permission{},
		&pkg.Product{},
		&pkg.ProductCategory{},
		&pkg.Order{},
		&pkg.Coupon{},
	)
	if err != nil {
		log.Fatalf("error occured while migrating database: %v", err)
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

	sp := stripeProvider.NewProvider(viper.GetString(confOptStripeKey), logger)

	API = api.NewInstanceAPI(&api.APIConfig{
		Logger:           logger,
		DBController:     dbController,
		CustomerClient:   customerClient,
		ExchangeProvider: exchangeProvider,
		StripeProvider:   sp,
		SecretKey:        viper.GetString(confOptSecretKey),
	})
}

func TestMain(m *testing.M) {
	setupTestEnvironment()
	code := m.Run()
	os.Exit(code)
}
