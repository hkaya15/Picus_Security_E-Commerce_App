package model

import "gorm.io/gorm"

type Category struct{
	gorm.Model
	CategoryID string `gorm:"unique"`
	CategoryName string
	IconURL string
}

type CategoryList []Category