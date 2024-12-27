package routes

import (
	"net/http"
	"order_service_saga/internal/handler"

	"github.com/gin-gonic/gin"
)

func RouteConfig(router *gin.Engine, orderHandler *handler.OrderHandler) {
	orderGroup := router.Group("/api/orders")
	{
		orderGroup.GET("/", orderHandler.GetAllOrders)
		orderGroup.POST("/", orderHandler.CreateOrder)
		orderGroup.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "hello world",
			})
		})
	}
}
