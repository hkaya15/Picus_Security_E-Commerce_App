package service

import (
	"errors"
	"net/http"
	"os"

	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/api/model"
	api "github.com/hkaya15/PicusSecurity/Final_Project/pkg/api/model"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/cart/model"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/cart/repository"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/product/repository"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/errors"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/jwt"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/helper"
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

// Add checks that ordinary with checks product, cart and add item to cart
func (c *CartService) Add(user *AccessTokenDetails, item *api.CartItem) error {
	var cart *Cart
	product, err := c.ProductRepo.GetProductById(*item.ProductID)
	if err != nil {
		return NewRestError(http.StatusBadRequest, os.Getenv("CHECK_USER_CART_PRODUCT_ISSUE"), err.Error())
	}

	cart, err = c.CartRepo.CreateCart(ResponseToCart(user.UserID))
	if err != nil {
		return NewRestError(http.StatusBadRequest, os.Getenv("CREATE_CART_ISSUE"), err.Error())
	}

	cartItem := ResponseToCartItem(item, product, user.UserID)

	_, bool := c.CartRepo.FindCartItemByID(cartItem.ProductID, cartItem.CartID)
	if bool == true {
		return NewRestError(http.StatusBadRequest, os.Getenv("CART_HAS_PRODUCT"), nil)
	}

	if product.ProductQuantity < uint(*item.Quantity) {
		return NewRestError(http.StatusBadRequest, os.Getenv("NOT_ENOUGH_PRODUCT"), nil)
	}

	err = c.CartRepo.AddItem(cartItem)
	if err != nil {
		return NewRestError(http.StatusBadRequest, os.Getenv("CREATE_CART_ITEM_ISSUE"), nil)
	}

	cartitems, err := c.CartRepo.GetCartItems(cart.UserID)

	
	cart= UpdateValues(*cart,cartitems)
	

	err = c.CartRepo.UpdateCart(cart)
	if err != nil {
		return NewRestError(http.StatusBadRequest, os.Getenv("ISSUE_ON_UPDATE_CART"), nil)
	}
	return nil
}

// GetCartList get cartlist includes all cart items
func (c *CartService) GetCartList(cart *Cart) (*Cart, error) {
	cart, err := c.CartRepo.GetCartList(cart)
	if err != nil {
		return nil, NewRestError(http.StatusBadRequest, os.Getenv("GET_CART_ISSUE"), err.Error())
	}
	return cart, nil
}

// Update  checks product whether exist or not. After that checks Cart that created before. After that it process logic and update cart item.
func (c *CartService) Update(req *UpdatedCartItem, userid string) error {
	var cart *Cart
	product, err := c.ProductRepo.GetProductById(*req.ProductID)
	if err != nil {
		return NewRestError(http.StatusBadRequest, os.Getenv("CHECK_USER_CART_PRODUCT_ISSUE"), err.Error())
	}

	cart, err = c.CartRepo.CreateCart(ResponseToCart(userid))
	if err != nil {
		return NewRestError(http.StatusBadRequest, os.Getenv("CREATE_CART_ISSUE"), err.Error())
	}

	item, bool := c.CartRepo.FindCartItemByID(*req.ProductID, userid)
	if !bool {
		return NewRestError(http.StatusBadRequest, os.Getenv("CART_HASNT_PRODUCT"), nil)
	}

	if product.ProductQuantity < uint(*req.Quantity) {
		return errors.New(os.Getenv("NOT_ENOUGH_PRODUCT"))
	}

	if uint(*req.Quantity) == 0 {
		res, err := c.CartRepo.Delete(item)
		if err != nil {
			return err
		}

		if res {
			cartitems, err := c.CartRepo.GetCartItems(cart.UserID)

			cart= UpdateValues(*cart,cartitems)

			err = c.CartRepo.UpdateCart(cart)
			if err != nil {
				return NewRestError(http.StatusBadRequest, os.Getenv("ISSUE_ON_UPDATE_CART"), nil)
			}
			return nil
		}
	}

	item.Quantity = uint(*req.Quantity)
	item.TotalPrice = float64(item.Quantity) * product.Price

	err = c.CartRepo.Update(item)
	if err != nil {
		return NewRestError(http.StatusBadRequest, os.Getenv("UPDATE_CHECK_PRODUCT_ISSUE"), nil)
	}

	cartitems, err := c.CartRepo.GetCartItems(cart.UserID)

	cart= UpdateValues(*cart,cartitems)

	err = c.CartRepo.UpdateCart(cart)
	if err != nil {
		return NewRestError(http.StatusBadRequest, os.Getenv("ISSUE_ON_UPDATE_CART"), nil)
	}

	return nil
}

// Delete checks products, cart and delete item.
func (c *CartService) Delete(productid string, userid string) (bool, error) {
	var cart *Cart
	_, err := c.ProductRepo.GetProductById(productid)
	if err != nil {
		return false, NewRestError(http.StatusBadRequest, os.Getenv("CHECK_USER_CART_PRODUCT_ISSUE"), err.Error())
	}

	cart, err = c.CartRepo.CreateCart(ResponseToCart(userid))
	if err != nil {
		return false, NewRestError(http.StatusBadRequest, os.Getenv("CREATE_CART_ISSUE"), err.Error())
	}

	item, bool := c.CartRepo.FindCartItemByID(productid, userid)
	if !bool {
		return false, NewRestError(http.StatusBadRequest, os.Getenv("CART_HASNT_PRODUCT"), nil)
	}

	res,err:= c.CartRepo.Delete(item)
	if err != nil {
		return false,err
	}

	if res {
		cartitems, err := c.CartRepo.GetCartItems(cart.UserID)

		cart= UpdateValues(*cart,cartitems)

		err = c.CartRepo.UpdateCart(cart)
		if err != nil {
			return false, NewRestError(http.StatusBadRequest, os.Getenv("ISSUE_ON_UPDATE_CART"), nil)
		}
		return true, nil
	}
	return false,err
}
