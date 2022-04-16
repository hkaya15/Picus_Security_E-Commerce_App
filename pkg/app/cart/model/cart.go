package model

import "time"

type Cart struct {
	UserID         string       `gorm:"primary_key;" json:"user_id"`
	Items          []*CartsItem `gorm:"foreignkey:CartID;" json:"items"`
	CompleteOrder  bool         `gorm:"default:false" json:"complete_order"`
	CartTotalPrice float64      `json:"total_price"`
	CartLength     int          `json:"cart_len"`
	CreatedAt      time.Time    `json:"createdAt"`
}
