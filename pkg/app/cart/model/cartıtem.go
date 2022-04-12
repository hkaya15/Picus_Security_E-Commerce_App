package model

import (
//	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/product/model"
)
type CartItem struct{
//	Product Product `gorm:"foreignKey:ProductId; references: ProductID"`
//	ProductID string
	CartItemID string 
	Quantity uint
	TotalPrice float64
}