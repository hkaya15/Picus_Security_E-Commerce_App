package model

import (
	"time"

	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/user/model"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID     string
	User       User         `gorm:"foreignKey:UserID;references:UserId"`
	OrderItems []*OrderItem `gorm:"foreignkey:OrderID;" json:"order_items"`
	OrderPrice float64      `json:"order_price"`
	OrderDate  time.Time    `json:"order_date"`
	IsCanceled bool         `json:"is_canceled"`
}
