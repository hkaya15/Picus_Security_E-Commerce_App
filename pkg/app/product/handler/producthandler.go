package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/api/model"
	api "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/product/model"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/product/service"
	"github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/config"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/errors"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/middleware"
	"go.uber.org/zap"
)

type ProductHandler struct {
	productService *ProductService
	cfg            *config.Config
}

func NewProductHandler(r *gin.RouterGroup, p *ProductService, cfg *config.Config) {
	h := &ProductHandler{productService: p, cfg: cfg}
	p.Migrate()
	r.POST("/create", AuthorizationMiddleware(h.cfg), h.create)
}
func (p *ProductHandler) Migrate() {
	p.productService.Migrate()
}

// create helps to create product
func (p *ProductHandler) create(c *gin.Context) {
	var req Product
	if err := c.Bind(&req); err != nil {
		zap.L().Error("product.handler.create", zap.Error(err))
		c.JSON(ErrorResponse(NewRestError(http.StatusBadRequest, os.Getenv("CHECK_YOUR_REQUEST"), nil)))
		return
	}

	if err := req.Validate(strfmt.NewFormats()); err != nil {
		zap.L().Error("product.handler.validate", zap.Error(err))
		c.JSON(ErrorResponse(err))
		return
	}
	pr, err := p.productService.Create(api.ResponseToProduct(req))
	if err != nil {
		zap.L().Error("product.handler.create", zap.Error(err))
		c.JSON(ErrorResponse(err))
		return
	}
	c.JSON(http.StatusCreated, APIResponse{Code: http.StatusCreated, Message: os.Getenv("CREATE_PRODUCT_SUCCESS"), Details: api.ProductToResponse(pr)})
	return

}
