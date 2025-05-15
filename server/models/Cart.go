package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Cart struct {
	gorm.Model `json:"-"`
	ID uint `gorm:"primary key" json:"id"`
	UserID     uint `gorm:"foreignKey" json:"user_id"`
	Items      []CartItem `gorm:"foreignKey:CartID;references:ID;constraint:OnDelete:CASCADE" json:"items"`
	Total float64 `json:"total"`
}

func (cart *Cart) AfterDelete(tx *gorm.DB) (err error) {
	tx.Clauses(clause.Returning{}).Where("cart_id = ?", cart.ID).Delete(&CartItem{})
	return
}