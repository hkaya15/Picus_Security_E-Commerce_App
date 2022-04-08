package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/api/model"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/user/model"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/user/service"
	"github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/config"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/errors"
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
		c.JSON(ErrorResponse(NewRestError(http.StatusBadRequest, "Check your request body", nil)))
		return
	}

	if err := req.Validate(strfmt.NewFormats()); err != nil {
		c.JSON(ErrorResponse(err))
		return
	}
	res,err:= u.userService.CheckUser(ResponseToUser(&req))
	if err!=nil{
		c.JSON(ErrorResponse(err))
		return
	}
	if res{
		c.JSON(ErrorResponse(NewRestError(http.StatusBadRequest, "User already exists", nil)))
		return
	}
	
	user, err := u.userService.Save(ResponseToUser(&req))
	if err != nil {
		c.JSON(ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, APIResponseSignUp{Code: http.StatusOK, Token: user.UserId})

}

func (u *UserHandler) Migrate() {
	u.userService.Migrate()

}
