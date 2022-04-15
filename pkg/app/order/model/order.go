package model

import (
	"time"

	"gorm.io/gorm"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/user/model"
)

type Order struct {
	gorm.Model
	ID string 
	UserID     string		`json:"user_id"`
	User       User         `gorm:"foreignKey:UserID;references:UserId"`
	OrderItems []OrderItem `gorm:"foreignkey:OrderID;" json:"order_items"`
	OrderPrice float64      `json:"order_price"`
	OrderDate  time.Time    `json:"order_date"`
	IsCanceled bool         `json:"is_canceled"`
}
