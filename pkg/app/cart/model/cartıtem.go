package model

import (
. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/product/model"
)
type CartItem struct{
	CartID string `gorm:"primary_key" json:"cart_id"`
	Product ProductBase `gorm:"foreignKey:ProductID; references:ProductId" json:"product"`
	ProductID string `gorm:"primary_key" json:"product_id"`
	Quantity uint `gorm:"not null" json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}