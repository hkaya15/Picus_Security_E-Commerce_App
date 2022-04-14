package model

import (
	"time"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/product/model"
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	OrderID   string       `json:"order_id"`
	ProductID string       `json:"product_id"`
	Product   *ProductBase `gorm:"foreignkey:ProductID" json:"product"`
	CreatedAt time.Time    `gorm:"<-:create" json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}
