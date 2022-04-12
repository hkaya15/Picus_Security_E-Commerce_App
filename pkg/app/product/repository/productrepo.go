package repository

import (
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/product/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p *ProductRepository) Migrate() {
	p.db.AutoMigrate(&ProductBase{})
}

func (p *ProductRepository) Create(pr *ProductBase) (*ProductBase, error) {
	zap.L().Debug("product.repo.create", zap.Reflect("product", pr.ProductName))
	if err := p.db.Create(&pr).Error; err != nil {
		zap.L().Error("product.repo.create failed to create product", zap.Error(err))
		return nil,err
	}
	return pr,nil
}

