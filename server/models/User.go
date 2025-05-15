package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	Password string `json:"password" example:"12345678"`
	Email    string `gorm:"unique;not null" json:"email" example:"example@example.com"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
}