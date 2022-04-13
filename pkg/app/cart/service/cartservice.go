package service

import(
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/cart/repository"
)

type CartService struct {
	CartRepo *CartRepository
}

func NewCartService(c *CartRepository) *CartService {
	return &CartService{CartRepo: c}
}

func (c *CartService) Migrate() {
	c.CartRepo.Migrate()
}