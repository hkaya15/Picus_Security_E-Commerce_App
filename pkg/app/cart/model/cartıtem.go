package model

import (
. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/product/model"
)
type CartItem struct{
	CartID string
	Product ProductBase `gorm:"foreignKey:ProductID; references:ProductId"`
	ProductID string 
	Quantity uint
	TotalPrice float64
}