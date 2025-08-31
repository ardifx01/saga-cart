package main

import (
	"log"
	"order_service_saga/internal/cache"
	"order_service_saga/internal/db"
	"order_service_saga/internal/handler"
	"order_service_saga/internal/kafka"
	"order_service_saga/internal/repository"
	"order_service_saga/internal/routes"
	"order_service_saga/internal/service"
	"order_service_saga/pkg"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	// CORS

	db, err := db.ConnectPostgre()
	if err != nil {
		log.Fatalf("error create connection to db: %v", err.Error())
	}

	orderCache := cache.NewRedis()
	orderRepo := repository.NewOrderRepo(db, orderCache)

	// setup kafka
	kafkaWriter := pkg.ConnectKafkaWriter()
	kafkaConsumer := pkg.ConnectKafkaReader("order-failed", "order-succes")
	defer kafkaWriter.Close()
	orderPublisher := kafka.NewOrderPublisher(kafkaWriter)
	orderConsumer := kafka.NewOrderConsumer(kafkaConsumer, orderRepo)
	go orderConsumer.Consume()
	// setup kafka

	orderService := service.NewOrderService(orderRepo, orderPublisher, orderCache)
	orderHandler := handler.NewOrderHandler(orderService)

	routes.RouteConfig(r, orderHandler)

	r.Run(":8082")
}
