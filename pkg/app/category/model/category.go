package model

import "gorm.io/gorm"

type Category struct{
	gorm.Model
	CategoryID string
	CategoryName string
	IconURL string
	SubCategories []int
}