package model

import (
	"gorm.io/gorm"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/category/model"
)

type ProductBase struct{
	gorm.Model
	ProductId string `gorm:"unique"`
	ProductName string
	ProductDescription string
	ProductQuantity uint
	CategoryId string `json:"categoryId"`
	Category Category `gorm:"foreignKey:CategoryID; references:CategoryId"`
	ImageURL string
	Price float64
	Counter uint `gorm:"default:0"`
}

type ProductList []ProductBase