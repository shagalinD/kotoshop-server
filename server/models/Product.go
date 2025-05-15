package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model `json:"-"`
	ID uint `gorm:"primary key" json:"id"`
	Title string `gorm:"required" json:"title" example:"MacBook Pro"`
	Price float64 `gorm:"required" json:"price" example:"1500000"`
	Description string `json:"description" example:"15.6 дюймов" `
	Image string `json:"image" example:"/assets/cat-surprised.gif"`
	Category string `gorm:"required" json:"category" example:"electronics"`
}