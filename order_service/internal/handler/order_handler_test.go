package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"order_service_saga/internal/domain"
	"order_service_saga/internal/handler"
	mock_contracts "order_service_saga/internal/mocks/contracts"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var ctlr *gomock.Controller
var serviceMock *mock_contracts.MockOrderServiceContract
var orderHandler *handler.OrderHandler

func TestOrderHandler_GetAllOrders(t *testing.T) {
	ctlr = gomock.NewController(t)
	serviceMock = mock_contracts.NewMockOrderServiceContract(ctlr)
	orderHandler = handler.NewOrderHandler(serviceMock)

	t.Run("Success get all orders", func(t *testing.T) {
		orders := &[]domain.Order{
			{
				ID:           1,
				CustomerName: "pepeg",
			},
		}

		serviceMock.EXPECT().GerOrders().Return(orders, nil)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/orders", orderHandler.GetAllOrders)

		req, _ := http.NewRequest(http.MethodGet, "/orders", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "pepeg")
	})

	t.Run("Gagal get orders", func(t *testing.T) {
		serviceMock.EXPECT().GerOrders().Return(nil, errors.New("error"))

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/orders", orderHandler.GetAllOrders)

		req, _ := http.NewRequest(http.MethodGet, "/orders", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "error")
	})
}

func TestOrderHandler_CreateOrder(t *testing.T) {
	ctlr = gomock.NewController(t)
	serviceMock = mock_contracts.NewMockOrderServiceContract(ctlr)
	orderHandler = handler.NewOrderHandler(serviceMock)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/orders", orderHandler.CreateOrder)

	t.Run("Error badrequest", func(t *testing.T) {
		request := `{
			"CustomerName": "aji",
		}`

		req, _ := http.NewRequest(http.MethodPost, "/orders", strings.NewReader(request))

		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Success Create Order", func(t *testing.T) {
		request := &domain.Order{
			CustomerName: "test",
			ProductID:    1,
			Qty:          2,
			Amount:       3,
			Status:       "Pending",
			OrderDate:    time.Now(),
		}

		payload, _ := json.Marshal(request)

		serviceMock.EXPECT().CreateOrder(gomock.Any()).Return(request, nil)

		req, _ := http.NewRequest(http.MethodPost, "/orders", strings.NewReader(string(payload)))

		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Contains(t, w.Body.String(), "test")
	})

	t.Run("Gagal Create Order karena order nil", func(t *testing.T) {
		request := &domain.Order{
			CustomerName: "test",
			ProductID:    1,
			Qty:          2,
			Amount:       3,
			Status:       "Pending",
			OrderDate:    time.Now(),
		}

		payload, _ := json.Marshal(request)

		serviceMock.EXPECT().CreateOrder(gomock.Any()).Return(nil, errors.New("error"))

		req, _ := http.NewRequest(http.MethodPost, "/orders", strings.NewReader(string(payload)))

		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "error error create order")
	})
}
