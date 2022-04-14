package model

import (
	"time"

	"github.com/google/uuid"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/cart/model"
)

func ResponseToOrderItem(c *CartsItem) *OrderItem {
	return &OrderItem{
		OrderID: uuid.New().String(),
		ProductID: c.ProductID,
		Product: &c.Product,
	}
}


func NewOrder(c *Cart) *Order {
	var orderPrice float64
	orderItems := make([]*OrderItem, 0)
	for _, item := range c.Items {
		orderPrice += item.TotalPrice*float64(item.Quantity)
		orderItems = append(orderItems,ResponseToOrderItem(item))
	}
	return &Order{
		UserID: c.UserID,
		OrderItems: orderItems,
		OrderPrice: orderPrice,
		OrderDate: time.Now(),
		IsCanceled: false,
	}
}
