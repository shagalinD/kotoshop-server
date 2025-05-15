package models

import "gorm.io/gorm"

type Feedback struct {
	gorm.Model `json:"-"`
	ID        uint `gorm:"primaryKey" json:"-"`
	Comment   string `json:"comment" example:"Кот просто восторг!"`
	Rating    float64 `gorm:"required;constraint:CHECK(rating >= 1 AND rating <= 5)" json:"rating" example:"5"`
	ProductID uint `json:"product_id"`
	Product   Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
	UserID uint `json:"user_id" gorm:"constraint:unique(user_id AND product_id)"`
	User User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
}