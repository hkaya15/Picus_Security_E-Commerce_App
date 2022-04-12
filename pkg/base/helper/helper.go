package helper

import (
	"encoding/csv"
	"mime/multipart"
	"net/http"
	"net/mail"
	"strings"
	"unicode"

	"github.com/google/uuid"
	"github.com/gorilla/securecookie"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/api/model"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/category/model"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/user/model"
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

func DecodeToken(req *http.Request, user *User) (*Token, error) {
	var hashKey = []byte("very-secret")
	var s = securecookie.New(hashKey, nil)
	var value Token
	if cookie, err := req.Cookie(user.UserId); err == nil {
		if err = s.Decode("token", cookie.Value, &value); err != nil {
			return nil, err
		}

	}
	return &value, nil
}

func ReadCSV(file *multipart.File) (CategoryList, error) {
	csvReader := csv.NewReader(*file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	var categorieslist CategoryList
	for _, line := range records[1:] {
		categorieslist = append(categorieslist, Category{
			CategoryID:   uuid.New().String(),
			CategoryName: line[0],
			IconURL:      line[1],
		})
	}
	return categorieslist, nil
}

func CompareCategories(db, uploaded *CategoryList) CategoryList {

	var out CategoryList

	up := *uploaded
	d := *db

	for i := 0; i < len(up); i++ {
		res := contains(d, up[i])
		if !res {
			out = append(out, up[i])
		}
	}
	return out
}

func contains(s CategoryList, e Category) bool {
	for _, a := range s {
		if strings.ToLower(a.CategoryName) == strings.ToLower(e.CategoryName) {
			return true
		}
	}
	return false
}
