package model

import (
	"github.com/google/uuid"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/api/model"
)

func ResponseToUser(u *SignUp) *User {
	userID := createId()
	return &User{
		UserId:    userID,
		FirstName: *u.Firstname,
		LastName:  *u.Lastname,
		Password:  *u.Password,
		Email:  *u.Email,
		IsAdmin: false,
	}
}

func createId() string {
	id := uuid.New().String()
	return id
}
