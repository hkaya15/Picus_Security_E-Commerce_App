package model

import (
	api "github.com/hkaya15/PicusSecurity/Final_Project/pkg/api/model"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/product/model"
)

func ResponseToCartItem(c *api.CartItem, p *ProductBase, userId string) *CartsItem {
	totalItemPrice := CalculateItemPrice(uint(*c.Quantity), p.Price)
	return &CartsItem{
		CartID:     userId,
		Product:    *p,
		ProductID:  *c.ProductID,
		Quantity:   uint(*c.Quantity),
		TotalPrice: totalItemPrice,
	}
}

func CalculateItemPrice(quantity uint, price float64) float64 {
	return float64(quantity) * price
}

func ResponseToCart(userId string) *Cart {
	return &Cart{
		UserID: userId,
		CompleteOrder: false,
	}
}

func ResponseAPICart(cart *Cart)*api.APICart{

	items := make([]*api.CartItemToResponse, 0)
	for _, v := range cart.Items{
		items = append(items,ItemResponse(v))
	}
	totalPrice:=CalculateCartTotalPrice(items)
	return &api.APICart{
		UserID: cart.UserID,
		Cartitems: items,
		Cartlength: int64(len(items)),
		CompleteOrder: *&cart.CompleteOrder,
		TotalPrice: totalPrice,
	}
}



func ItemResponse(b *CartsItem) *api.CartItemToResponse{
	product:=ProductToResponse(&b.Product)
	return &api.CartItemToResponse{
		CartID: b.CartID,
		Product: product,
		ProductID: b.ProductID,
		Quantity: int64(b.Quantity),
		TotalPrice: b.TotalPrice,
	}
}

func CalculateCartTotalPrice(items []*api.CartItemToResponse)float64{
	var sum float64
	for i := 0; i < len(items); i++ {
		sum=sum+items[i].TotalPrice
	}
	return sum
}
