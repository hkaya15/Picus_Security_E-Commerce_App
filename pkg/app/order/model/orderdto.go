package model

import (
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"

	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/api/model"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/cart/model"
)


func NewOrder(userid string, list []OrderItem) *Order {
	var orderPrice float64
	for _, item := range list {
		orderPrice += item.Product.Price
	}
	return &Order{
		ID:         uuid.NewString(),
		UserID:     userid,
		OrderItems: list,
		OrderPrice: orderPrice,
		OrderDate:  time.Now(),
		IsCanceled: false,
	}
}

func OrderToAPI(o Order) *OrderAPI {
	orderItems := make([]*OrderItemResponse, 0)
	for _, item := range o.OrderItems {
		orderItems = append(orderItems, ItemToOrderItemOrderList(&item))
	}

	return &OrderAPI{
		Userid:     o.UserID,
		Orderdate:  strfmt.DateTime(o.OrderDate),
		Iscanceled: o.IsCanceled,
		Orderprice: int64(o.OrderPrice),
		Orderitems: orderItems,
	}
}

func ItemToOrderItemOrderList(o *OrderItem) *OrderItemResponse {
	return &OrderItemResponse{
		Productid: o.ProductID,
		Userid: o.OrderID,
		Quantity: uint64(o.Quantity),
	}
}

func GetAllOrderToAPI(o []Order) *[]OrderAPI {
	var ordertoapi *OrderAPI
	orderapilist := make([]OrderAPI, 0)
	for _, v := range o {
		ordertoapi = OrderToAPI(v)
		orderapilist = append(orderapilist, *ordertoapi)
	}

	return &orderapilist
}


func NewOrderItem(userıd string, c CartsItem) *OrderItem {
	return &OrderItem{
		OrderID:   userıd,
		ProductID: c.Product.ProductId,
		UserID:    userıd,
		Product: &c.Product,
		Quantity: c.Quantity,
	}
}