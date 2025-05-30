package handler

import (
	"fmt"
	"net/http"
	"order_service_saga/internal/contracts"
	"order_service_saga/internal/domain"

	"time"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService contracts.OrderServiceContract
}

type OrderRequest struct {
	CustomerName string  `json:"customer_name"`
	ProductID    int     `json:"product_id"`
	Qty          int     `json:"qty"`
	Amount       float64 `json:"amount"`
}

func NewOrderHandler(orderService contracts.OrderServiceContract) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

func (h *OrderHandler) GetAllOrders(c *gin.Context) {
	order_list, err := h.orderService.GerOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("error get all orders: %v", err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, order_list)
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req OrderRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("error create order (request): %v", err.Error()),
		})
		return
	}

	order := domain.Order{
		CustomerName: req.CustomerName,
		ProductID:    req.ProductID,
		Qty:          req.Qty,
		Amount:       req.Amount,
		Status:       "Pending",
		OrderDate:    time.Now(),
	}

	order_created, err := h.orderService.CreateOrder(order)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("error error create order: %v", err.Error()),
		})
		return
	}

	c.JSON(http.StatusCreated, order_created)
}
