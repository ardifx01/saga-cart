package routes

import (
	"net/http"
	"product_service_saga/internal/handler"

	"github.com/gin-gonic/gin"
)

func RouteConfig(router *gin.Engine, productHandler *handler.ProductHandler) {
	productGroup := router.Group("/api/products")
	{
		productGroup.GET("/", productHandler.GetAll)
		productGroup.GET("/:id", productHandler.GetProduct)
		productGroup.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "hello world",
			})
		})
	}
}
