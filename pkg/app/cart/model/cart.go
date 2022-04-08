package model

import "gorm.io/gorm"


type Cart struct{
	gorm.Model
	UserId string
	Items []CartItem `gorm:"foreignKey:CartItemID; references: ItemID"`
	ItemID string
	CompleteOrder bool
	CartTotalPrice float64
	CartLength int
}