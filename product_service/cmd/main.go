package main

import (
	"log"
	"product_service_saga/internal/db"
	"product_service_saga/internal/handler"
	"product_service_saga/internal/kafka"
	"product_service_saga/internal/repository"
	"product_service_saga/internal/routes"
	"product_service_saga/internal/service"
	"product_service_saga/pkg"

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
		log.Printf("error connect to postgreSQL: %v", err.Error())
	}

	// setup kafka
	kafkaWriter := pkg.ConnectKafkaWriter()
	kafkaReader := pkg.ConnectKafkaReader("stock-reserved", "stock-failed")

	productPublisher := kafka.NewProductPublisher(kafkaWriter)
	productConsumer := kafka.NewProductConsumer(kafkaReader, productPublisher, db)
	go productConsumer.Consume()
	// setup kafka

	productRepo := repository.NewProductRepo(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	routes.RouteConfig(r, productHandler)

	r.Run(":8081")
}
