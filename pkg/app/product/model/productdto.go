package model

import (
	api "github.com/hkaya15/PicusSecurity/Final_Project/pkg/api/model"
	"github.com/google/uuid"
)

func ResponseToProduct(p api.Product) *ProductBase {
	return &ProductBase{
		ProductId: uuid.New().String(),
		ProductName: *p.Name,
		ProductDescription: *p.Description,
		ProductQuantity: uint(*p.Quantity),
		CategoryId: *p.CategoryID,
		ImageURL: *p.ImageURL,
		Price: *p.Price,
	}
}

func ProductToResponse(p *ProductBase) *api.Product{
	quantity:=int64(p.ProductQuantity)
	return &api.Product{
		CategoryID:  &p.CategoryId,
		Description: &p.ProductDescription,
		ImageURL:   &p.ImageURL,
		Name:        &p.ProductName,
		Price:       &p.Price,
		Quantity: &quantity,
	}
}


func SearchToResponse(pl *ProductList) []*api.Product{
	products := make([]*api.Product, 0)
	for _, p := range *pl {
		products = append(products, ProductToResponse(&p))
	}
	return products
}