package repository

import (
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/cart/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (c *CartRepository) Migrate() {
	c.db.AutoMigrate(&CartItem{})
	c.db.AutoMigrate(&Cart{})

}

// AddItem helps to add CartItem
func (c *CartRepository) AddItem(cartItem *CartItem) error {
	zap.L().Debug("cart.repo.create.cartıtem", zap.Reflect("cartItem", cartItem))
	if err := c.db.Create(cartItem).Error; err != nil {
		zap.L().Error("cart.repo.create.cartitem failed to create cartitem", zap.Error(err))
		return err
	}
	return nil
}

// FindByID helps to check the item that has product
func (c *CartRepository) FindByID(cartItem *CartItem) bool {
	var item *CartItem
	var exists bool = false
	zap.L().Debug("cart.repo.findbyıd.cartıtem", zap.Reflect("cartItem", cartItem))
	r := c.db.Where("product_id=?", cartItem.ProductID).Limit(1).Find(&item)
	if r.RowsAffected > 0 {
		exists = true
	}
	return exists
}

// CreateCart helps to check cart is exist. if not create for you
func (c *CartRepository) CreateCart(crt *Cart) error {
	var cart *Cart
	zap.L().Debug("cart.repo.create", zap.Reflect("cart", crt))
	if err := c.db.Where(Cart{UserID: crt.UserID}).Attrs(crt).Preload("Items").Preload("Items.Product").Preload("Items.Product.Category").FirstOrCreate(&cart).Error; err != nil {
		zap.L().Error("cart.repo.create failed to create cart", zap.Error(err))
		return err
	}
	return nil
}
