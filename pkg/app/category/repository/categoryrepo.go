package repository

import (
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/category/model"
	"gorm.io/gorm"
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
	return int(count),"Created!", err
}

func (c *CategoryRepository) GetAll() (*CategoryList, error) {
	var categories CategoryList
	err := c.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return &categories, nil
}
