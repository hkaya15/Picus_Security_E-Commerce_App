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
	c.db.AutoMigrate(&CartsItem{})
	c.db.AutoMigrate(&Cart{})

}

// AddItem helps to add CartItem
func (c *CartRepository) AddItem(cartItem *CartsItem) error {
	zap.L().Debug("cart.repo.create.cartıtem", zap.Reflect("cartItem", cartItem))
	if err := c.db.Create(cartItem).Error; err != nil {
		zap.L().Error("cart.repo.create.cartitem failed to create cartitem", zap.Error(err))
		return err
	}
	return nil
}

// FindByID helps to check the item that has product
func (c *CartRepository) FindByID(productId string, cartId string) (*CartsItem,bool) {
	var item *CartsItem
	var exists bool = false
	zap.L().Debug("cart.repo.findbyıd.cartıtem", zap.String("cartItem", productId))
	r := c.db.Where("product_id=?", productId).Limit(1).Find(&item)
	if r.RowsAffected > 0 {
		exists = true
	}
	return item, exists
}

// CreateCart helps to check cart is exist. if not create for you
func (c *CartRepository) CreateCart(crt *Cart) (*Cart,error) {
	var cart *Cart
	zap.L().Debug("cart.repo.create", zap.Reflect("cart", crt))
	if err := c.db.Where(Cart{UserID: crt.UserID}).Attrs(crt).FirstOrCreate(&cart).Error; err != nil {
		zap.L().Error("cart.repo.create failed to create cart", zap.Error(err))
		return nil, err
	}
	return cart, nil
}

// GetCartList helps to get cartlist
func (c *CartRepository) GetCartList(crt *Cart) (*Cart, error) {
	zap.L().Debug("cart.repo.getCartItem", zap.Reflect("cart", crt))
	var cart *Cart
	err := c.db.Where(&Cart{UserID: crt.UserID}).Attrs(crt).Preload("Items.Product.Category").Find(&cart).Error
	if err != nil {
		zap.L().Error("cart.repo.getCartItem failed to get cart", zap.Error(err))
		return nil, err
	}
	return cart, nil
}

func (c *CartRepository) Update(crt *CartsItem)  error {
	zap.L().Debug("cart.repo.update", zap.Reflect("item", crt))
	if err := c.db.Save(&crt).Error; err != nil {
		zap.L().Error("product.repo.update failed to update product", zap.Error(err))
		return err
	}
	return nil
}