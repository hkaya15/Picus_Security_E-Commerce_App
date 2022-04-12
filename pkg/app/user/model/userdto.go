package model

import (
	"github.com/google/uuid"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/api/model"
)

func ResponseToUser(u *SignUp) *User {
	return &User{
		UserId:    uuid.New().String(),
		FirstName: *u.Firstname,
		LastName:  *u.Lastname,
		Password:  *u.Password,
		Email:  *u.Email,
		IsAdmin: false,
		Order: nil,
	}
}


