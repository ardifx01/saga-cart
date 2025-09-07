package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"product_service_saga/internal/contracts"
	"product_service_saga/internal/domain"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService contracts.ProductServiceContract
	esClient       *elasticsearch.TypedClient
}

func NewProductHandler(productService contracts.ProductServiceContract, esClient *elasticsearch.TypedClient) *ProductHandler {
	return &ProductHandler{
		productService: productService,
		esClient:       esClient,
	}
}

func (h *ProductHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "1"))

	list_products, total, err := h.productService.GetProductsPaginate(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    nil,
			"message": fmt.Sprintf("error get products (handler): %v", err.Error()),
		})
	}

	total_page := (total + int64(pageSize) - 1) / int64(pageSize)

	c.JSON(http.StatusOK, gin.H{
		"data":       list_products,
		"page":       page,
		"page_size":  pageSize,
		"total":      total,
		"total_page": total_page,
	})
}

func (h *ProductHandler) IndexAllProducts() error {
	products, err := h.productService.GetProducts()
	if err != nil {
		return err
	}

	for _, prod := range products {
		_, err := h.esClient.Index("products").Id(fmt.Sprintf("%d", prod.ID)).Request(prod).Do(context.Background())
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *ProductHandler) SearchProducts(c *gin.Context) {
	query := c.Query("q")

	res, err := h.esClient.Search().
		Index("products").
		Query(&types.Query{
			MultiMatch: &types.MultiMatchQuery{
				Query:  query,
				Fields: []string{"name", "description"},
			},
		}).
		Do(c.Request.Context())

	if err != nil {
		log.Printf("Elasticsearch error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	products := make([]*domain.Product, 0, len(res.Hits.Hits))
	for _, hit := range res.Hits.Hits {
		var p domain.Product
		if err := json.Unmarshal(hit.Source_, &p); err != nil {
			log.Printf("Failed to unmarshal hit: %v", err)
			continue
		}
		products = append(products, &p)
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
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
