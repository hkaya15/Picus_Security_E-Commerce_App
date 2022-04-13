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

func (p *ProductRepository) Search(query string) (*ProductList,error) {
	zap.L().Debug("product.repo.search", zap.Reflect("query", query))
	var products ProductList
	if err:=p.db.Preload("Category").Where("product_name ILIKE ? OR product_description ILIKE ? OR product_id ILIKE ?", "%"+query+"%", "%"+query+"%", "%"+query+"%").Find(&products).Error; err!=nil{
		zap.L().Error("product.repo.search failed to search product", zap.Error(err))
		return nil,err
	}

	return &products,nil
}

// Update take product id. It can be uses without id with Save. But it creates different product id's. Here it can be uses with Save() without Where() conditions,
// because id is implemented product. Here I would like to differentiate usage practices.
func (p *ProductRepository) Update(pr *ProductBase, id string) (*ProductBase,error) {
	zap.L().Debug("product.repo.update", zap.Reflect("query", pr))
	pr.ProductId=id
	if err := p.db.Where("product_id=?",id).Updates(&pr).Error; err!=nil{
		zap.L().Error("product.repo.update failed to update product", zap.Error(err))
		return nil,err
	}

	return pr, nil
	
}