package service

import (
	"mime/multipart"
	"os"

	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/category/model"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/category/repository"
	"go.uber.org/zap"

	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/helper"
)

type CategoryService struct {
	CategoryRepo *CategoryRepository
}

func NewCategoryService(c *CategoryRepository) *CategoryService {
	return &CategoryService{CategoryRepo: c}
}

func (c *CategoryService) Migrate() {
	c.CategoryRepo.Migrate()
}

// Upload helps to read file and compare after that create on db
func (c *CategoryService) Upload(file *multipart.File) (int, string, error) {
	var count int
	var str string

	categorylist, err := ReadCSV(file)
	if err != nil {
		zap.L().Error("category.service.readcsv", zap.Error(err))
		return 0,os.Getenv("ERROR"), err
	}
	categoriesOnDb, err := c.GetAll()
	if err != nil {
		return 0,os.Getenv("ERROR"), err
	}

	if len(*categoriesOnDb) > 0 {
		compared := CompareCategories(categoriesOnDb, &categorylist)
		if len(compared) > 0 {
			count,str, err = c.CategoryRepo.Upload(&compared)
			if err != nil {
				return count,os.Getenv("ERROR"), err
			}
			return count,str,nil
			
		}
		return 0,os.Getenv("SAME_CATEGORY"),nil
	}
	count,str, err = c.CategoryRepo.Upload(&categorylist)
	if err != nil {
		return count,os.Getenv("ERROR"), err
	}

	return count,str, nil

}

// GetAll returns all categories
func (c *CategoryService) GetAll() (*CategoryList, error) {
	return c.CategoryRepo.GetAll()
}
