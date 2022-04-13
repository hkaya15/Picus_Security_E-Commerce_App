package repository

import (
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/cart/model"
	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (c *CartRepository) Migrate() {
	c.db.AutoMigrate(&Cart{})
}
