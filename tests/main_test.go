package tests

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/pstano1/customer-api/client"
	"github.com/pstano1/go-cart/internal/api"
	"github.com/pstano1/go-cart/internal/db"
	exchange "github.com/pstano1/go-cart/internal/pkg/exchangeProvider"
	"github.com/pstano1/go-cart/internal/pkg/stripeProvider"
	"github.com/spf13/viper"
	"github.com/testcontainers/testcontainers-go"
	postgresMock "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var API *api.InstanceAPI

const (
	dbName           = "testDB"
	dbUser           = "testDBUser"
	dbPassword       = "strong-password"
	confOptStripeKey = "STRIPE_KEY"
	customerId       = "c6a2b5a1-6851-438b-a055-2ae0d1116b50"
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

	postgresContainer, err := postgresMock.RunContainer(ctx,
		testcontainers.WithImage("docker.io/postgres:16-alpine"),
		postgresMock.WithInitScripts(filepath.Join("mock_data.sql")),
		postgresMock.WithDatabase(dbName),
		postgresMock.WithUsername(dbUser),
		postgresMock.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(30*time.Second)),
	)
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

	dbController := db.NewDBController(
		gormDB,
		logger,
	)

	lis, serve := mockGRPCServer()
	go serve(lis)
	addr := lis.Addr().String()
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
		SecretKey:        "example-secret-key",
	})
}

func mockGRPCServer() (net.Listener, func(lis net.Listener) error) {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()

	return lis, grpcServer.Serve
}

func TestMain(m *testing.M) {
	setupTestEnvironment()
	code := m.Run()
	os.Exit(code)
}
