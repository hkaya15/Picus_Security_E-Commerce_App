package repository

import (
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/order/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (o *OrderRepository) Migrate() {
	o.db.AutoMigrate(&Order{})
	o.db.AutoMigrate(&OrderItem{})
}

func (o *OrderRepository) CompleteOrder(order *Order) (*Order,error){
	zap.L().Debug("order.repo.complete", zap.Reflect("order", order))
	if err := o.db.Create(&order).Error; err != nil {
		zap.L().Error("order.repo.Create failed to create order", zap.Error(err))
		return nil, err
	}
	return order, nil
}
