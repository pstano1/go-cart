package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pstano1/go-cart/internal/pkg"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	confOptDatabaseHost     = "DATABASE_HOST"
	confOptDatabaseName     = "DATABASE_NAME"
	confOptDatabaseUsername = "DATABASE_USERNAME"
	confOptDatabasePassword = "DATABASE_PASSWORD"
	confOptDatabasePort     = "DATABASE_PORT"
)

func migrateDatabase(db *gorm.DB) {
	err := db.AutoMigrate(
		&pkg.User{},
		&pkg.Permission{},
		&pkg.Product{},
		&pkg.ProductCategory{},
		&pkg.Order{},
		&pkg.Coupon{},
	)
	if err != nil {
		log.Fatalf("error occured while migrating database: %v", err)
		return
	}
}

func createPermissions(db *gorm.DB) {

}

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("error occured when loading config %v", err)
		return
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
		return
	}
	if len(os.Args) < 2 {
		log.Fatal("please provide a valid action")
		return
	}

	for _, action := range os.Args[1:] {
		switch action {
		case "migrate":
			migrateDatabase(gormDB)
		case "create-permissions":
			createPermissions(gormDB)
		default:
			log.Fatalf("invalid action %s. skipping...", action)
		}
	}
}
