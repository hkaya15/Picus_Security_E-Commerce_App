package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/order/service"

	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/order/model"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/api/model"
	"github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/config"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/errors"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/jwt"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/middleware"
	"go.uber.org/zap"
)

type OrderHandler struct {
	orderService *OrderService
	cfg          *config.Config
}

func NewOrderHandler(r *gin.RouterGroup, o *OrderService, cfg *config.Config) {
	h := &OrderHandler{orderService: o, cfg: cfg}
	o.Migrate()
	r.POST("/complete", AuthenticationMiddleware(h.cfg), h.completeorder)
	r.GET("/", AuthenticationMiddleware(h.cfg), h.getallorders)
}

func (o *OrderHandler) Migrate() {
	o.orderService.Migrate()
}

func (o *OrderHandler) completeorder(c *gin.Context) {
	val, res := c.Get("User")
	if res == false {
		zap.L().Error("order.handler.completeorder", zap.Bool("value: ", res))
		c.JSON(ErrorResponse(NewRestError(http.StatusInternalServerError, os.Getenv("NO_CONTEXT"), nil)))
		return
	}
	user := val.(*AccessTokenDetails)

	err := o.orderService.CompleteOrder(user)
	if err != nil {
		zap.L().Error("order.handler.completeorder", zap.Error(err))
		c.JSON(ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, APIResponse{Code: http.StatusOK, Message: os.Getenv("ORDER_SUCCESS")})
	return
}

func (o *OrderHandler) getallorders(c *gin.Context) {
	val, res := c.Get("User")
	if res == false {
		zap.L().Error("order.handler.completeorder", zap.Bool("value: ", res))
		c.JSON(ErrorResponse(NewRestError(http.StatusInternalServerError, os.Getenv("NO_CONTEXT"), nil)))
		return
	}
	user := val.(*AccessTokenDetails)
	orders,err:= o.orderService.GetAllOrders(user.UserID)
	if err != nil {
		zap.L().Error("order.handler.getallorders", zap.Error(err))
		c.JSON(ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, APIResponse{Code: http.StatusOK, Message: os.Getenv("ORDER_SUCCESS"),Details:GetAllOrderToAPI(orders)})
	return
}
