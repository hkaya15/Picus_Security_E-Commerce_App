package model

import (
	"time"

	"gorm.io/gorm"
	//	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/cart/model"
)

type Order struct {
	gorm.Model
	OrderId    string
	UserId     string
	CartItemID string
	//Cart Cart `gorm:"foreignKey:ItemID; references: CartItemID"`
	OrderDate  time.Time
	IsCanceled bool
	IsOrdered  bool
}
