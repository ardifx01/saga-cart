package service

import (
	"encoding/json"
	"log"
	"order_service_saga/internal/contracts"
	"order_service_saga/internal/domain"
	"order_service_saga/internal/kafka"
)

type OrderService struct {
	orderRepo contracts.OrderRepoContract
	publisher *kafka.OrderPublisher
}

func NewOrderService(orderRepo contracts.OrderRepoContract, publisher *kafka.OrderPublisher) contracts.OrderServiceContract {
	return &OrderService{
		orderRepo: orderRepo,
		publisher: publisher,
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

	go s.publisher.Publish("stock-reserved", message_byte)

	return order_created, nil
}
