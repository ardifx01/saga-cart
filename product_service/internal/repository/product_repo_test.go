package repository_test

import (
	"errors"
	"product_service_saga/internal/domain"
	mock_repository "product_service_saga/internal/mocks/repository"
	"product_service_saga/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestProudctRepo_GetProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockGorm := mock_repository.NewMockIGorm(ctrl)
	productRepo := repository.NewProductRepo(mockGorm)

	t.Run("Sucess get all products", func(t *testing.T) {
		products := []*domain.Product{
			{ID: 1, Name: "Product 1"},
			{ID: 2, Name: "Product 2"},
		}

		mockGorm.
			EXPECT().
			Find(gomock.Any()).
			DoAndReturn(func(dest interface{}, _ ...interface{}) *gorm.DB {
				ptr := dest.(*[]*domain.Product)
				*ptr = products
				return &gorm.DB{Error: nil}
			})

		result, err := productRepo.GetProducts()
		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, "Product 1", result[0].Name)
	})

	t.Run("Error get products", func(t *testing.T) {
		mockGorm.EXPECT().Find(gomock.Any()).Return(&gorm.DB{Error: errors.New("error")})

		_, err := productRepo.GetProducts()
		assert.Error(t, err)
	})
}

func TestProudctRepo_FindByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockGorm := mock_repository.NewMockIGorm(ctrl)
	productRepo := repository.NewProductRepo(mockGorm)

	t.Run("Success get product by id", func(t *testing.T) {
		product := &domain.Product{ID: 1, Name: "Test Product"}

		mockGorm.
			EXPECT().
			First(gomock.Any(), gomock.Any()).
			DoAndReturn(func(dest interface{}, _ ...interface{}) *gorm.DB {
				// dest is *domain.Product
				ptr := dest.(**domain.Product)
				*ptr = product
				return &gorm.DB{Error: nil}
			})

		result, err := productRepo.FindByID(1)
		assert.NoError(t, err)
		assert.Equal(t, "Test Product", result.Name)
	})

	t.Run("Error get product by id", func(t *testing.T) {
		mockGorm.EXPECT().First(gomock.Any(), gomock.Any()).Return(&gorm.DB{Error: errors.New("error")})

		_, err := productRepo.FindByID(1)
		assert.Error(t, err)
	})
}
