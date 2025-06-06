package service_test

import (
	"errors"
	"order_service_saga/internal/domain"
	mock_contracts "order_service_saga/internal/mocks/contracts"
	mock_service "order_service_saga/internal/mocks/order"
	"order_service_saga/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetOrders(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock_contracts.NewMockOrderRepoContract(ctrl)
	mockPublisher := mock_service.NewMockIOrderPublisher(ctrl)
	orderService := service.NewOrderService(mockRepo, mockPublisher)

	t.Run("Sucess get orders", func(t *testing.T) {
		orders := []domain.Order{
			{
				ID: 1,
			},
		}

		mockRepo.EXPECT().GerOrders().Return(&orders, nil)

		res, err := orderService.GerOrders()
		assert.NoError(t, err)
		assert.Len(t, *res, 1)
	})

	t.Run("Error get all orders", func(t *testing.T) {
		mockRepo.EXPECT().GerOrders().Return(nil, errors.New("error"))

		_, err := orderService.GerOrders()
		assert.Error(t, err)
	})
}

func TestCreateOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock_contracts.NewMockOrderRepoContract(ctrl)
	mockPublisher := mock_service.NewMockIOrderPublisher(ctrl)
	orderService := service.NewOrderService(mockRepo, mockPublisher)

	t.Run("Success create order", func(t *testing.T) {
		newOrder := domain.Order{
			ID: 1,
		}

		mockRepo.EXPECT().CreateOrder(gomock.Any()).Return(&newOrder, nil)
		mockPublisher.EXPECT().Publish("stock-reserved", gomock.Any()).Return(nil)

		res, err := orderService.CreateOrder(newOrder)
		assert.NoError(t, err)
		assert.Equal(t, newOrder.ID, res.ID)
	})

	t.Run("Error create order", func(t *testing.T) {
		mockRepo.EXPECT().CreateOrder(gomock.Any()).Return(nil, errors.New("error"))

		_, err := orderService.CreateOrder(domain.Order{
			ID: 1,
		})
		assert.Error(t, err)
	})
}
