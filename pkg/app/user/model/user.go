package model

import (

	"gorm.io/gorm"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/order/model"
)

type User struct {
	gorm.Model
	UserId    string `gorm:"unique"`
	FirstName string
	LastName  string
	Email     string
	Password  string
	IsAdmin   bool `gorm:"default:false"`
	Order []Order 
}
