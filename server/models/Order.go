package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model `swaggerignore:"true"`
	UserID uint `json:"user_id"`
	User User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Total float64 `json:"total" example:"53.999"`
	Address string `json:"address" example:"Россия, Москва, Верхняя Первомайская, 52"`
	Items []OrderItem `json:"foreignKey:OrderID"`
	Status string `json:"status" example:"created"`
	OrderNumber string `json:"order_number" example:"ORD-2025-1010"`
	Date time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"date"`
}

type OrderItem struct {
	gorm.Model `swaggerignore:"true"`
	OrderID uint `json:"cart_id"`
	Product Product `gorm:"foreignKey:ProductID;constraint:OnDelete:Cascade"`
	ProductID uint `json:"product_id"`
	Quantity uint `json:"quantity"`
}

func (order *Order) BeforeCreate(tx *gorm.DB) error {
	if order.OrderNumber == "" {
			// Формат: ORD-ГОД-ПОСЛЕДНИЙ_ID+1
			var lastOrder Order
			var lastOrderID uint
			if err := tx.Order("id desc").First(&lastOrder).Error; err != nil {
				lastOrderID = 0
			} else {
				lastOrderID = lastOrder.ID
			}

			year := time.Now().Year()
			newID := int(lastOrderID) + 1

			order.OrderNumber = fmt.Sprintf("ORD-%d-%04d", year, newID)
	}
	return nil
}