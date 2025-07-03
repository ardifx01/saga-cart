package repository

import (
	"context"
	"encoding/json"
	"log"
	"order_service_saga/internal/cache"
	"order_service_saga/internal/contracts"
	"order_service_saga/internal/domain"

	"gorm.io/gorm"
)

type OrderRepo struct {
	db    *gorm.DB
	cache *cache.Redis
}

func NewOrderRepo(db *gorm.DB, cache *cache.Redis) contracts.OrderRepoContract {
	return &OrderRepo{
		db:    db,
		cache: cache,
	}
}

func (o *OrderRepo) GerOrders() (*[]domain.Order, error) {
	var orderList []domain.Order
	ctx := context.Background()

	val, err := o.cache.Get(ctx, "order-list")
	if err == nil {
		json.Unmarshal([]byte(val), &orderList)
		return &orderList, nil
	}

	err = o.db.Find(&orderList).Error
	if err != nil {
		log.Printf("error get all orders: %v", err.Error())
		return nil, err
	}

	orderData, _ := json.Marshal(orderList)

	_ = o.cache.Set(ctx, "order-list", orderData, 1)

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

func (o *OrderRepo) UpdateOrderStatus(orderID int, status bool) error {
	var orderFind domain.Order
	result := o.db.First(&orderFind, orderID)
	if result.Error != nil {
		log.Printf("error while finding order to update status: %v", result.Error)
		return result.Error
	}

	if !status {
		orderFind.Status = "Berhasil"
	} else {
		orderFind.Status = "Gagal"
	}

	result = o.db.Save(&orderFind)
	if result.Error != nil {
		log.Printf("error while updating status order: %v", result.Error)
		return result.Error
	}

	return nil
}
