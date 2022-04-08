package model

import "gorm.io/gorm"


type Product struct{
	gorm.Model
	ProductId string
	ProductName string
	ProductDescription string
	ProductQantity uint
	CategoryId int
	ImageURL string
	Price float64
	DiscountedPrice float64
	StoreId string
	Counter uint
	UserId string
}