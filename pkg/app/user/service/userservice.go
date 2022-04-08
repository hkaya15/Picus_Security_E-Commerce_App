package service

import (
	"net/http"
	"net/mail"
	"unicode"

	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/user/model"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/user/repository"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo *UserRepository
}

func NewUserService(u *UserRepository) *UserService {
	return &UserService{UserRepo: u}
}
func (u *UserService) CheckUser(user *User) (bool,error){
	return u.UserRepo.CheckUser(user)
	
}

func (u *UserService) Save(user *User) (*User, error) {
	if verifyEMail(user.Email) {
		ver := verifyPassword(user.Password)
		if ver == true {
			hash,err:=hashPassword(user.Password)
			if err!=nil{
				return nil, NewRestError(http.StatusBadRequest,"Problem on creating password",nil)
			}
			user.Password=hash
			return u.UserRepo.Save(user)
		}
		return nil, NewRestError(http.StatusBadRequest, "Password is weak! Please Enter password that contains 1 Number, 1 Upper, 1 Lower, 1 Special Char and 7 length", nil)
	}
	return nil, NewRestError(http.StatusBadRequest, "Please enter a valid e-mail", nil)
}

func (u *UserService) Migrate() {
	u.UserRepo.Migrate()
}

func verifyEMail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func verifyPassword(s string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(s) >= 7 {
		hasMinLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}
