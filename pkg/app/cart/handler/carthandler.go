package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/api/model"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/cart/service"
	"github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/config"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/errors"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/jwt"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/middleware"
	"go.uber.org/zap"
)

type CartHandler struct {
	cartService *CartService
	cfg         *config.Config
}

func NewCartHandler(r *gin.RouterGroup, c *CartService, cfg *config.Config) {
	h := &CartHandler{cartService: c, cfg: cfg}
	c.Migrate()
	r.POST("/add", AuthenticationMiddleware(h.cfg), h.add)
}

func (c *CartHandler) Migrate() {
	c.cartService.Migrate()
}

func (crt *CartHandler) add(c *gin.Context) {
	val, res := c.Get("User")
	if res == false {
		zap.L().Error("cart.handler.add", zap.Bool("value: ", res))
		c.JSON(ErrorResponse(NewRestError(http.StatusInternalServerError, os.Getenv("NO_CONTEXT"), nil)))
		return
	}
	user := val.(*AccessTokenDetails)

	var req CartItem
	if err := c.Bind(&req); err != nil {
		zap.L().Error("cart.handler.add", zap.Error(err))
		c.JSON(ErrorResponse(NewRestError(http.StatusBadRequest, os.Getenv("CHECK_YOUR_REQUEST"), nil)))
		return
	}

	if err := req.Validate(strfmt.NewFormats()); err != nil {
		zap.L().Error("user.handler.signup", zap.Error(err))
		c.JSON(ErrorResponse(err))
		return
	}
	err := crt.cartService.Add(user, &req)
	if err != nil {
		zap.L().Error("cart.handler.add", zap.Error(err))
		c.JSON(http.StatusInternalServerError,APIResponse{Code: http.StatusInternalServerError,Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, APIResponse{Code: http.StatusCreated, Message: os.Getenv("CREATE_CART_ITEM_SUCCESS")})
	return
}
