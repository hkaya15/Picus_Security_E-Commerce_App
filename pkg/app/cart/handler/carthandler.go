package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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
	fmt.Println(val.(*AccessTokenDetails).Email)
}
