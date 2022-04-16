package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/api/model"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/category/service"
	"github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/config"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/errors"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/middleware"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/pagination"
	"go.uber.org/zap"
)

type CategoryHandler struct {
	categoryService *CategoryService
	cfg             *config.Config
}

func NewCategoryHandler(r *gin.RouterGroup, c *CategoryService, cfg *config.Config) {
	h := &CategoryHandler{categoryService: c, cfg: cfg}
	c.Migrate()
	r.POST("/upload",AuthorizationMiddleware(h.cfg), h.upload)
	r.GET("/",PaginationMiddleware(h.cfg), h.getcategories)

}

func (c *CategoryHandler) Migrate() {
	c.categoryService.Migrate()
}

// upload helps to user create bulk category. If category has implemented before db, upload checks is without any fault. 
func (ct *CategoryHandler) upload(c *gin.Context) {
	file, handler, err := c.Request.FormFile("myFile")
	if err != nil {
		zap.L().Error("category.handler.upload", zap.Error(err))
		c.JSON(ErrorResponse(NewRestError(http.StatusInternalServerError, os.Getenv("READ_FILE_FAULT"), nil)))
		return
	}
	defer file.Close()
	zap.L().Debug("category.handler.upload:", zap.String("filename", fmt.Sprintf("Uploaded File: %+v\n", handler.Filename)))
	zap.L().Debug("category.handler.upload:", zap.String("filesize", fmt.Sprintf("File Size: %+v\n", handler.Size)))
	zap.L().Debug("category.handler.upload:", zap.String("MIME", fmt.Sprintf("MIME Header: %+v\n", handler.Header)))

	count,str, err := ct.categoryService.Upload(&file)

	if err != nil {
		zap.L().Error("category.handler.upload", zap.Error(err))
		c.JSON(ErrorResponse(NewRestError(http.StatusInternalServerError, os.Getenv("UPLOAD_FILE_FAULT"), err.Error())))
		return
	}
	c.JSON(http.StatusOK, APIResponse{Code: http.StatusOK,Message: fmt.Sprintf("%v Record",count),Details:str })

}


func (ct *CategoryHandler) getcategories(c *gin.Context){
	val,res:=c.Get("Pagination")
	if res == false {
		zap.L().Error("category.handler.getcategories", zap.Bool("value: ", res))
		c.JSON(ErrorResponse(NewRestError(http.StatusInternalServerError, os.Getenv("NO_CONTEXT"), nil)))
		return
	}
	pag:=val.(Pagination)
	p,err:= ct.categoryService.GetAllCategoriesWithPagination(pag)
	if err != nil {
		zap.L().Error("category.handler.getcategories", zap.Error(err))
		c.JSON(ErrorResponse(NewRestError(http.StatusInternalServerError, os.Getenv("PAGINATION_FAULT"), nil)))
		return
	}
	c.JSON(http.StatusOK, APIResponse{Code: http.StatusOK, Message: os.Getenv("CATEGORIES_GET_SUCCESS"), Details: NewPage(*p)})
	return
}