package service

import (
	"net/http"
	"os"

	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/product/model"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/product/repository"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/errors"
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

func (p *ProductService) Create(pr *ProductBase) (*ProductBase, error) {
	return p.ProductRepo.Create(pr)
}

func (p *ProductService) Search(query string) (*ProductList, error) {
	return p.ProductRepo.Search(query)
}

func (p *ProductService) Update(pr *ProductBase, id string) (*ProductBase, error) {
	res,err:=p.ProductRepo.CheckProduct(id)
	if err != nil {
		return nil, NewRestError(http.StatusBadRequest, os.Getenv("UPDATE_CHECK_PRODUCT_ISSUE"), nil)
	}
	if res{
		return p.ProductRepo.Update(pr, id)
	}else{
		return nil, NewRestError(http.StatusBadRequest, os.Getenv("NO_PRODUCT"), nil)
	}
	
}

func (p *ProductService) Delete(id string) (bool, error) {
	res, err := p.ProductRepo.CheckProduct(id)
	if err != nil {
		return false, NewRestError(http.StatusBadRequest, os.Getenv("DELETE_CHECK_PRODUCT_ISSUE"), nil)
	}
	if res{
		return p.ProductRepo.Delete(id)
	}else{
		return false, NewRestError(http.StatusBadRequest, os.Getenv("NO_PRODUCT"), nil)
	}
	
}
