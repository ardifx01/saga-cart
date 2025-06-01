package service_test

import (
	"errors"
	"product_service_saga/internal/domain"
	"product_service_saga/internal/service"
	mock_contracts "product_service_saga/mocks/contracts"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockProdRepo := mock_contracts.NewMockProductRepoContract(ctrl)
	prodService := service.NewProductService(mockProdRepo)

	t.Run("Get All Products", func(t *testing.T) {
		products := []*domain.Product{
			{
				ID:   1,
				Name: "Celana",
			},
		}
		mockProdRepo.EXPECT().GetProducts().Return(products, nil)

		res, err := prodService.GetProducts()
		assert.NoError(t, err)
		assert.Len(t, res, 1)
		assert.Equal(t, "Celana", res[0].Name)
	})

	t.Run("Error get products", func(t *testing.T) {
		mockErr := errors.New("database error")
		mockProdRepo.EXPECT().GetProducts().Return(nil, mockErr)

		res, err := prodService.GetProducts()
		assert.Nil(t, res)
		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
	})
}

func TestFindByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockProdRepo := mock_contracts.NewMockProductRepoContract(ctrl)
	prodService := service.NewProductService(mockProdRepo)

	t.Run("Success get product by id", func(t *testing.T) {
		prod := &domain.Product{
			ID: 1,
		}
		mockProdRepo.EXPECT().FindByID(1).Return(prod, nil)

		res, err := prodService.FindByID(1)
		assert.NoError(t, err)
		assert.Equal(t, prod.ID, res.ID)
	})

	t.Run("Fail get product by id", func(t *testing.T) {
		mockProdRepo.EXPECT().FindByID(1).Return(nil, errors.New("prod not found"))

		_, err := prodService.FindByID(1)
		assert.ErrorContains(t, err, "prod not found")
	})
}
