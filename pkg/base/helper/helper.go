package helper

import (
	"net/http"
	"net/mail"
	"unicode"

	"github.com/gorilla/securecookie"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/api/model"
	"golang.org/x/crypto/bcrypt"
)

func VerifyEMail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func VerifyPassword(s string) bool {
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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func DecodeToken(req *http.Request) (*Token,error){
	var hashKey = []byte("very-secret")
	var s = securecookie.New(hashKey, nil)
	var value Token
	if cookie, err := req.Cookie("token"); err == nil {
		if err = s.Decode("token", cookie.Value, &value); err != nil {
			return nil,err
		}
		
	}
	return &value,nil
}
