package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/gorilla/securecookie"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/api/model"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/user/model"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/user/service"
	"github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/config"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/errors"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/jwt"

	"go.uber.org/zap"
)

type UserHandler struct {
	userService *UserService
	cfg         *config.Config
}

func NewUserHandler(r *gin.RouterGroup, u *UserService, cfg *config.Config) {
	h := &UserHandler{userService: u, cfg: cfg}
	u.Migrate()
	r.POST("/signup", h.signup)
}

func (u *UserHandler) signup(c *gin.Context) {
	var req SignUp
	if err := c.Bind(&req); err != nil {
		zap.L().Error("user.handler.signup", zap.Error(err))
		c.JSON(ErrorResponse(NewRestError(http.StatusBadRequest, "Check your request body", nil)))
		return
	}

	if err := req.Validate(strfmt.NewFormats()); err != nil {
		zap.L().Error("user.handler.signup", zap.Error(err))
		c.JSON(ErrorResponse(err))
		return
	}
	res, err := u.userService.CheckUser(ResponseToUser(&req))
	if err != nil {
		zap.L().Error("user.handler.signup", zap.Error(err))
		c.JSON(ErrorResponse(err))
		return
	}
	if res {
		zap.L().Error("user.handler.signup: User Already exist")
		c.JSON(ErrorResponse(NewRestError(http.StatusBadRequest, "User already exists", nil)))
		return
	}

	user, err := u.userService.Save(ResponseToUser(&req))
	if err != nil {
		zap.L().Error("user.handler.signup", zap.Error(err))
		c.JSON(ErrorResponse(err))
		return
	}

	tkn, err := GenerateToken(user, u.cfg)
	if err != nil {
		zap.L().Error("user.handler.signup: generatetoken", zap.Error(err))
		c.JSON(ErrorResponse(err))
		return
	}

	

	var hashKey = []byte("very-secret")
	var s = securecookie.New(hashKey, nil)
	encoded, err := s.Encode("token", tkn)
	if err == nil {
		cookie := &http.Cookie{
			Name:     "token",
			Value:    encoded,
			Path:     "/",
			Domain:   "127.0.0.1",
			Secure:   false,
			HttpOnly: false,
		}
		http.SetCookie(c.Writer, cookie)

		c.JSON(http.StatusCreated, APIResponseSignUp{Code: http.StatusCreated, Token: tkn})


	}
}

func (u *UserHandler) Migrate() {
	u.userService.Migrate()
}
