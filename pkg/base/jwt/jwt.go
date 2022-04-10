package jwtpack

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/api/model"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/user/model"
	"github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/config"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/errors"
	"go.uber.org/zap"
)

type TokenDetails struct {
	UserID       string
	Email        string
	Role         bool
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
	jwt.StandardClaims
}

func GenerateToken(user *User, cfg *config.Config) (*Token, error) {
	td := &TokenDetails{}
	tkn := &Token{}

	td.AtExpires = time.Now().Add(time.Duration(cfg.JWTConfig.AccessSessionTime) * time.Second).Unix()
	td.AccessUuid = uuid.New().String()
	td.UserID = user.UserId
	td.Email = user.Email

	td.RtExpires = time.Now().Add(time.Duration(cfg.JWTConfig.RefreshSessionTime) * time.Second).Unix()
	td.RefreshUuid = uuid.New().String()

	atClaims := jwt.MapClaims{}
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = td.UserID
	atClaims["email"] = td.Email
	atClaims["role"] = td.Role
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	accessTokenResult, err := at.SignedString([]byte(cfg.JWTConfig.SecretKey))
	td.AccessToken = accessTokenResult
	if err != nil {
		zap.L().Debug("generatetoken.accesstoken: ", zap.Error(err))
		return nil, NewRestError(http.StatusBadRequest, "Issue on generating access token", nil)
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = td.UserID
	rtClaims["email"] = td.Email
	rtClaims["role"] = td.Role
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(cfg.JWTConfig.SecretKey))
	if err != nil {
		zap.L().Debug("generatetoken.refreshtoken: ", zap.Error(err))
		return nil, NewRestError(http.StatusBadRequest, "Issue on generating refresh token", nil)
	}
	tkn.AccessToken = td.AccessToken
	tkn.RefreshToken = td.RefreshToken
	return tkn,nil
}
