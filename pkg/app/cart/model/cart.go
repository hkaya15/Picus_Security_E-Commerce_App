package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID         string  `gorm:"unique"`
	Items          []CartItem 
	CompleteOrder  bool	`gorm:"default:false"`
	CartTotalPrice float64
	CartLength     int
}
