package handler

import (
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/api/model"
	model "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/product/model"
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
	r.GET("/search", h.search)
	r.PUT("/:id", AuthorizationMiddleware(h.cfg), h.update)
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
	pr, err := p.productService.Create(model.ResponseToProduct(req))
	if err != nil {
		zap.L().Error("product.handler.create", zap.Error(err))
		c.JSON(ErrorResponse(err))
		return
	}
	c.JSON(http.StatusCreated, APIResponse{Code: http.StatusCreated, Message: os.Getenv("CREATE_PRODUCT_SUCCESS"), Details: model.ProductToResponse(pr)})
	return

}

// search helps user to search by product_id,name,description. It can be expanded.
func (p *ProductHandler) search(c *gin.Context) {
	query := c.Query("query")
	if len(query) == 0 {
		zap.L().Error("product.handler.search", zap.Error(errors.New(os.Getenv("NULL_SEARCH"))))
		c.JSON(http.StatusInternalServerError, APIResponse{Code: http.StatusInternalServerError, Message: os.Getenv("NULL_SEARCH")})
		return
	}
	products, err := p.productService.Search(query)
	if err != nil {
		zap.L().Error("product.handler.search", zap.Error(err))
		c.JSON(ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, APIResponse{Code: http.StatusOK, Message: os.Getenv("SEARCH_PRODUCT_SUCCESS"), Details: model.SearchToResponse(products)})
	return
}

// update helps user to update product by id 
func (p *ProductHandler) update(c *gin.Context) {
	id := c.Param("id")
	var req Product
	if err := c.Bind(&req); err != nil {
		zap.L().Error("product.handler.update", zap.Error(err))
		c.JSON(ErrorResponse(NewRestError(http.StatusBadRequest, os.Getenv("CHECK_YOUR_REQUEST"), err)))
		return
	}
	if err := req.Validate(strfmt.NewFormats()); err != nil {
		zap.L().Error("product.handler.validate", zap.Error(err))
		c.JSON(ErrorResponse(err))
		return
	}
	pr, err := p.productService.Update(model.ResponseToProduct(req), id)
	if err != nil {
		zap.L().Error("product.handler.update", zap.Error(err))
		c.JSON(ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, APIResponse{Code: http.StatusOK, Message: os.Getenv("UPDATE_PRODUCT_SUCCESS"), Details: model.ProductToResponse(pr)})
	return
}
