package model

import (
	"time"

	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/cart/model"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderId    string
	UserId     string
	Cart       Cart `gorm:"foreignKey:UserId; references:UserID"`
	OrderDate  time.Time
	IsCanceled bool
	IsOrdered  bool
}
