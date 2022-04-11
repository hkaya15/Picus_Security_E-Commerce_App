package service

import (
	"net/http"

	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/user/model"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/user/repository"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/errors"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/helper"
)

type UserService struct {
	UserRepo *UserRepository
}

func NewUserService(u *UserRepository) *UserService {
	return &UserService{UserRepo: u}
}
func (u *UserService) CheckUser(user *User) (bool, error) {
	return u.UserRepo.CheckUser(user)

}

func (u *UserService) Save(user *User) (*User, error) {
	if VerifyEMail(user.Email) {
		ver := VerifyPassword(user.Password)
		if ver == true {
			hash, err := HashPassword(user.Password)
			if err != nil {
				return nil, NewRestError(http.StatusBadRequest, "Problem on creating password", nil)
			}
			user.Password = hash
			return u.UserRepo.Save(user)
		}
		return nil, NewRestError(http.StatusBadRequest, "Password is weak! Please Enter password that contains 1 Number, 1 Upper, 1 Lower, 1 Special Char and 7 length", nil)
	}
	return nil, NewRestError(http.StatusBadRequest, "Please enter a valid e-mail", nil)
}

func (u *UserService) Login(email string, password string) (*User,error){
	user,err:= u.UserRepo.Login(email)
	if err != nil {
		return nil, NewRestError(http.StatusBadRequest, "Problem on decoding password", nil)
	}
	res:=CheckPasswordHash(password,user.Password)
	if !res{
		return nil, NewRestError(http.StatusBadRequest, "Wrong e-mail or password", nil)
	}
	return user,nil
}

func (u *UserService) Migrate() {
	u.UserRepo.Migrate()
}
