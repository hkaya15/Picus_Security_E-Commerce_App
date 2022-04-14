package service

import (
	"errors"
	"net/http"
	"os"

	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/cart/model"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/cart/repository"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/order/model"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/order/repository"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/product/repository"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/errors"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/jwt"
)

type OrderService struct {
	OrderRepo   *OrderRepository
	CartRepo    *CartRepository
	ProductRepo *ProductRepository
}

func NewOrderService(o *OrderRepository, crt *CartRepository, pr *ProductRepository) *OrderService {
	return &OrderService{OrderRepo: o, CartRepo: crt, ProductRepo: pr}
}

func (o *OrderService) Migrate() {
	o.OrderRepo.Migrate()
}

func (o *OrderService) CompleteOrder(user *AccessTokenDetails) (*Order,error) {
	cart, err := o.CartRepo.CreateCart(ResponseToCart(user.UserID))
	if err != nil {
		return nil,NewRestError(http.StatusBadRequest, os.Getenv("CREATE_CART_ISSUE"), err.Error())
	}
	fullcart, err := o.CartRepo.GetCartList(cart)
	if err != nil {
		return nil,NewRestError(http.StatusBadRequest, os.Getenv("GET_CART_ISSUE"), err.Error())
	}
	if len(fullcart.Items) == 0 {
		return nil, errors.New(os.Getenv("CART_EMPTY_FAIL"))
	}
	order,err:= o.OrderRepo.CompleteOrder(NewOrder(fullcart))
	if err!=nil{
		return nil, NewRestError(http.StatusBadRequest, os.Getenv("ORDER_ISSUE"), err.Error())
	}
	for _,v:=range fullcart.Items{
		 res,err := o.CartRepo.Delete(v)
		 if err != nil {
			return nil, errors.New(os.Getenv("DELETE_ITEM_ISSUE"))
		}
		if res==true{
			continue
		}
	}
	return order,nil
	
}
