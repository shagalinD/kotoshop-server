package postgres

import (
	"kotoshop/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Open(postgresString string) {
	var err error

	DB, err = gorm.Open(postgres.Open(os.Getenv("POSTGRES_STRING")), &gorm.Config{})

	if err != nil {
		log.Fatal("Error on accessing database")
	}

	migratingErr := DB.AutoMigrate(&models.Product{}, &models.User{}, &models.Feedback{}, &models.Cart{}, &models.CartItem{}, &models.Order{}, &models.OrderItem{})

	if migratingErr != nil {
		log.Fatal("Error on migrating")
	}
}