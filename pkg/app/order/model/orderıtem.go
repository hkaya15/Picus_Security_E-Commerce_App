package model

import (
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/product/model"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/user/model"
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	OrderID   string `gorm:"primary_key;" json:"order_id"`
	UserID    string
	User      User         `gorm:"foreignKey:UserID; references: Id" json:"user"`
	ProductID string       `json:"product_id"`
	Product   *ProductBase `gorm:"foreignkey:ProductID; references:Id" json:"product"`
	Quantity uint
}
