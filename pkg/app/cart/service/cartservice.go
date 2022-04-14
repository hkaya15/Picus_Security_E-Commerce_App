package service

import (
	"errors"
	"net/http"
	"os"

	api "github.com/hkaya15/PicusSecurity/Final_Project/pkg/api/model"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/cart/model"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/cart/repository"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/product/repository"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/errors"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/jwt"
)

type CartService struct {
	CartRepo    *CartRepository
	ProductRepo *ProductRepository
}

func NewCartService(c *CartRepository, pr *ProductRepository) *CartService {
	return &CartService{CartRepo: c, ProductRepo: pr}
}

func (c *CartService) Migrate() {
	c.CartRepo.Migrate()
}

// Add checks that ordinary with checks product, cart
func (c *CartService) Add(user *AccessTokenDetails, item *api.CartItem) error {

	product, err := c.ProductRepo.GetProductById(*item.ProductID)
	if err != nil {
		return NewRestError(http.StatusBadRequest, os.Getenv("CHECK_USER_CART_PRODUCT_ISSUE"), err.Error())
	}

	err = c.CartRepo.CreateCart(ResponseToCart(user.UserID))
	if err != nil {
		return NewRestError(http.StatusBadRequest, os.Getenv("CREATE_CART_ISSUE"), err.Error())
	}

	cartItem := ResponseToCartItem(item, product, user.UserID)

	bool := c.CartRepo.FindByID(cartItem)
	if bool == true {
		return NewRestError(http.StatusBadRequest, os.Getenv("CART_HAS_PRODUCT"), nil)
	}

	if product.ProductQuantity < uint(*item.Quantity) {
		return errors.New(os.Getenv("NOT_ENOUGH_PRODUCT"))
	}

	err = c.CartRepo.AddItem(cartItem)
	if err != nil {
		return NewRestError(http.StatusBadRequest, os.Getenv("CREATE_CART_ITEM_ISSUE"), nil)
	}

	return nil
}

// GetCartList get cartlist includes all cart items
func (c *CartService) GetCartList(cart *Cart) (*Cart,error){
	cart, err := c.CartRepo.GetCartList(cart)
	if err != nil {
		return nil, NewRestError(http.StatusBadRequest, os.Getenv("GET_CART_ISSUE"), err.Error())
	}
	return cart,nil
}
