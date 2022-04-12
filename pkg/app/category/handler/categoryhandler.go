package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/category/service"
	"github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/config"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/errors"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/api/model"
	"go.uber.org/zap"
)

type CategoryHandler struct {
	categoryService *CategoryService
	cfg             *config.Config
}

func NewCategoryHandler(r *gin.RouterGroup, c *CategoryService, cfg *config.Config) {
	h := &CategoryHandler{categoryService: c, cfg: cfg}
	c.Migrate()
	r.POST("/upload", h.upload)

}

func (c *CategoryHandler) Migrate() {
	c.categoryService.Migrate()
}

func (ct *CategoryHandler) upload(c *gin.Context) {
	file, handler, err := c.Request.FormFile("myFile")
	if err != nil {
		zap.L().Error("category.handler.upload", zap.Error(err))
		c.JSON(ErrorResponse(NewRestError(http.StatusInternalServerError, "Fault on getting file", nil)))
		return
	}
	defer file.Close()
	zap.L().Debug("category.handler.upload:", zap.String("filename", fmt.Sprintf("Uploaded File: %+v\n", handler.Filename)))
	zap.L().Debug("category.handler.upload:", zap.String("filesize", fmt.Sprintf("File Size: %+v\n", handler.Size)))
	zap.L().Debug("category.handler.upload:", zap.String("MIME", fmt.Sprintf("MIME Header: %+v\n", handler.Header)))

	count,str, err := ct.categoryService.Upload(&file)

	if err != nil {
		zap.L().Error("category.handler.upload", zap.Error(err))
		c.JSON(ErrorResponse(NewRestError(http.StatusInternalServerError, "Fault on getting file", err.Error())))
		return
	}
	c.JSON(http.StatusOK, APIResponse{Code: http.StatusOK,Message: fmt.Sprintf("%v Record",count),Details:str })

}
