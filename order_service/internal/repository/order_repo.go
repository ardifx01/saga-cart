package repository

import (
	"log"
	"order_service_saga/internal/contracts"
	"order_service_saga/internal/domain"

	"gorm.io/gorm"
)

type OrderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) contracts.OrderRepoContract {
	return &OrderRepo{
		db: db,
	}
}

func (o *OrderRepo) GerOrders() (*[]domain.Order, error) {
	var orderList []domain.Order
	err := o.db.Find(&orderList).Error
	if err != nil {
		log.Printf("error get all orders: %v", err.Error())
		return nil, err
	}
	return &orderList, nil
}

func (o *OrderRepo) CreateOrder(order domain.Order) (*domain.Order, error) {
	err := o.db.Create(&order).Error
	if err != nil {
		log.Printf("error create order: %v", err.Error())
		return nil, err
	}
	return &order, nil
}
