package main

import (
	"log"
	"order_service_saga/pkg"
	"payment_service_saga/internal/db"
	"payment_service_saga/internal/kafka"

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

	// Postgre
	_, err := db.ConnectPostgre()
	if err != nil {
		log.Fatalf("error create connection to db (payment-service): %v", err.Error())
	}
	// Postgre

	// setup kafka
	kafkaWriter := pkg.ConnectKafkaWriter()
	kafkaReader := pkg.ConnectKafkaReader("payment-request")
	paymentPublisher := kafka.NewPaymentPublisher(kafkaWriter)
	paymentConsumer := kafka.NewPaymentConsumer(kafkaReader, paymentPublisher)
	go paymentConsumer.Consume()
	// setup kafka

	r.Run(":8083")
}
