package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model `json:"-"`
	ID uint `gorm:"primary key" json:"id"`
	CartID uint `json:"cart_id"`
	ProductID uint `json:"product_id"`
	Product Product `gorm:"foreignKey:ProductID" json:"product"`
	Quantity uint `json:"quantity"`
	Price float64 `json:"-"`
}

type RequestCartItem struct {
	ProductID uint `json:"product_id"`
	Quantity uint `json:"quantity"`
}

type RequestRemoveCartItem struct {
	ProductID uint `json:"product_id"`
}


