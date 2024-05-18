package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/pstano1/go-cart/internal/pkg"
	"github.com/spf13/viper"
	"go.uber.org/zap"
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

func flushDatabase(db *gorm.DB) {
	models := []interface{}{
		&pkg.User{},
		&pkg.Permission{},
		&pkg.Product{},
		&pkg.ProductCategory{},
		&pkg.Order{},
		&pkg.Coupon{},
	}
	for _, model := range models {
		err := db.Migrator().DropTable(model)
		if err != nil {
			log.Fatal("error while dropping table",
				zap.Error(err),
			)
		}
	}
}

func createPermissions(db *gorm.DB) {
	db.FirstOrCreate(&pkg.Permission{Name: pkg.GetUser})
	db.FirstOrCreate(&pkg.Permission{Name: pkg.GetUserSelf})
	db.FirstOrCreate(&pkg.Permission{Name: pkg.CreateUser})
	db.FirstOrCreate(&pkg.Permission{Name: pkg.UpdateUser})
	db.FirstOrCreate(&pkg.Permission{Name: pkg.UpdateUserSelf})
	db.FirstOrCreate(&pkg.Permission{Name: pkg.DeleteUser})
	db.FirstOrCreate(&pkg.Permission{Name: pkg.DeleteUserSelf})
	db.FirstOrCreate(&pkg.Permission{Name: pkg.CreateProduct})
	db.FirstOrCreate(&pkg.Permission{Name: pkg.UpdateProduct})
	db.FirstOrCreate(&pkg.Permission{Name: pkg.DeleteProduct})
	db.FirstOrCreate(&pkg.Permission{Name: pkg.CreateCategory})
	db.FirstOrCreate(&pkg.Permission{Name: pkg.UpdateCategory})
	db.FirstOrCreate(&pkg.Permission{Name: pkg.DeleteCategory})
	db.FirstOrCreate(&pkg.Permission{Name: pkg.CreateCoupon})
	db.FirstOrCreate(&pkg.Permission{Name: pkg.UpdateCoupon})
	db.FirstOrCreate(&pkg.Permission{Name: pkg.DeleteCoupon})
	db.FirstOrCreate(&pkg.Permission{Name: pkg.GetOrders})
	db.FirstOrCreate(&pkg.Permission{Name: pkg.UpdateOrder})
	db.FirstOrCreate(&pkg.Permission{Name: pkg.DeleteOrder})
}

func createServiceAccount(db *gorm.DB) {
	var permissions []string
	rows, err := db.Model(pkg.Permission{}).Rows()
	if err != nil {
		return
	}
	for rows.Next() {
		var name string
		if err = rows.Scan(
			&name,
		); err != nil {
			return
		}
		permissions = append(permissions, name)
	}

	randomBytes := make([]byte, 32)
	_, err = rand.Read(randomBytes)
	if err != nil {

		return
	}
	password := base64.URLEncoding.EncodeToString(randomBytes)

	serviceAccount := &pkg.UserCreate{
		Username:      "service-account",
		Password:      password,
		PasswordCheck: password,
		Email:         "",
		Permissions:   permissions,
		IsActive:      true,
	}
	println("Save the printed password since it can't be accessed later")
	fmt.Printf("%s's password: %s\n", serviceAccount.Username, password)
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
		case "flush":
			flushDatabase(gormDB)
			migrateDatabase(gormDB)
		case "create-service-account":
			println("not implemented yet")
			createServiceAccount(gormDB)
		default:
			log.Fatalf("invalid action %s. skipping...", action)
		}
	}
}
