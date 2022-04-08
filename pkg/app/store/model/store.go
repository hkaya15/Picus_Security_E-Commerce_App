package model

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	StoreName string
	StoreId string
	Products []string
}
