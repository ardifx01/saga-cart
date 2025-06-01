package handler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"product_service_saga/internal/domain"
	"product_service_saga/internal/handler"
	mock_contracts "product_service_saga/mocks/contracts"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_contracts.NewMockProductServiceContract(ctrl)
	handler := handler.NewProductHandler(mockService)

	t.Run("Success get products", func(t *testing.T) {
		expectedProducts := []*domain.Product{
			{ID: 1, Name: "Product A"},
		}

		mockService.EXPECT().GetProducts().Return(expectedProducts, nil)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/products", handler.GetAll)

		req, _ := http.NewRequest(http.MethodGet, "/products", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Product A")
	})

	t.Run("Gagal get products", func(t *testing.T) {

		mockService.EXPECT().GetProducts().Return(nil, errors.New("error get products"))

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/products", handler.GetAll)

		req, _ := http.NewRequest(http.MethodGet, "/products", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "error get products")
	})
	
}