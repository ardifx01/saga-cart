package handler

import (
	"fmt"
	"net/http"
	"product_service_saga/internal/contracts"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService contracts.ProductServiceContract
}

func NewProductHandler(productService contracts.ProductServiceContract) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (h *ProductHandler) GetAll(c *gin.Context) {
	list_products, err := h.productService.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    nil,
			"message": fmt.Sprintf("error get products (handler): %v", err.Error()),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": list_products,
	})
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	id_param, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("id is not valid: %v", err.Error()),
		})
		return
	}

	product, err := h.productService.FindByID(id_param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("product with id %v is not found: %v", id_param, err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}
