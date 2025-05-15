package postgres

import (
	"fmt"
	"kotoshop/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Open(postgresString string) {
	var err error

	DB, err = gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	os.Getenv("DB_HOST"),
	os.Getenv("DB_USER"),    // "postgres"
	os.Getenv("DB_PASSWORD"),// ваш пароль
	os.Getenv("DB_NAME"),    // имя БД
	os.Getenv("DB_PORT"))), &gorm.Config{})

	if err != nil {
		log.Fatal("Error on accessing database")
	}

	migratingErr := DB.AutoMigrate(&models.Product{}, &models.User{}, &models.Feedback{}, &models.Cart{}, &models.CartItem{}, &models.Order{}, &models.OrderItem{})

	if migratingErr != nil {
		log.Fatal("Error on migrating")
	}
}