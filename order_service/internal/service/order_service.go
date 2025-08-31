package service

import (
	"context"
	"encoding/json"
	"log"
	"order_service_saga/internal/cache"
	"order_service_saga/internal/contracts"
	"order_service_saga/internal/domain"
)

type IOrderPublisher interface {
	Publish(topic string, message []byte) error
}

type OrderService struct {
	orderRepo contracts.OrderRepoContract
	publisher IOrderPublisher
	cache     cache.ICache
}

func NewOrderService(orderRepo contracts.OrderRepoContract, publisher IOrderPublisher, cache cache.ICache) contracts.OrderServiceContract {
	return &OrderService{
		orderRepo: orderRepo,
		publisher: publisher,
		cache:     cache,
	}
}

func (s *OrderService) GerOrders() (*[]domain.Order, error) {
	order_list, err := s.orderRepo.GerOrders()
	if err != nil {
		log.Printf("error while get all orders: %v", err.Error())
		return nil, err
	}
	return order_list, nil
}

func (s *OrderService) CreateOrder(order domain.Order) (*domain.Order, error) {
	order_created, err := s.orderRepo.CreateOrder(order)
	if err != nil {
		log.Printf("error while create order: %v", err.Error())
		return nil, err
	}

	message_byte, _ := json.Marshal(order_created)
	s.cache.Delete(context.Background(), "order-list")

	go s.publisher.Publish("stock-reserved", message_byte)

	return order_created, nil
}
