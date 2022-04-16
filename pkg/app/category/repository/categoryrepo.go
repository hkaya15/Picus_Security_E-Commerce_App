package repository

import (
	"os"

	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/category/model"
	"gorm.io/gorm"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/api/model"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (c *CategoryRepository) Migrate() {
	c.db.AutoMigrate(&Category{})
}

func (c *CategoryRepository) Upload(categories *CategoryList) (int,string, error) {
	var count int64
	err := c.db.Create(&categories).Count(&count).Error
	return int(count),os.Getenv("CREATE_FILE"), err
}

func (c *CategoryRepository) GetAll() (*CategoryList, error) {
	var categories CategoryList
	err := c.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return &categories, nil
}

func (c *CategoryRepository) GetAllCategoriesWithPagination(pag Pagination)([]Category, int,error){
	var categories []Category
	var count int64
	err:=c.db.Offset(int(pag.Page) - 1 * int(pag.PageSize)).Limit(int(pag.PageSize)).Find(&categories).Count(&count).Error
	if err!=nil{
		return nil,-1,err
	}
	return categories,int(count),nil
}
