package service

import (
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/product/repository"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/product/model"
)

type ProductService struct {
	ProductRepo *ProductRepository
}

func NewProductService(p *ProductRepository) *ProductService {
	return &ProductService{ProductRepo: p}
}

func (p *ProductService) Migrate() {
	p.ProductRepo.Migrate()
}

func (p *ProductService) Create(pr *ProductBase)(*ProductBase, error){
	return p.ProductRepo.Create(pr)
}