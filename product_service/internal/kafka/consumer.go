package kafka

import (
	"context"
	"encoding/json"
	"log"
	"product_service_saga/events"
	"product_service_saga/internal/domain"

	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

type ProductConsumer struct {
	reader *kafka.Reader
	db     *gorm.DB
}

func NewProductConsumer(reader *kafka.Reader, db *gorm.DB) *ProductConsumer {
	return &ProductConsumer{
		reader: reader,
		db:     db,
	}
}

func (c *ProductConsumer) Consume() {
	defer c.reader.Close()
	for {
		msg, err := c.reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		switch msg.Topic {
		case "order-created":
			log.Printf("Message received (order-created): value=%s\n", string(msg.Value))
			processDecreaseQtyProduct(c.db, msg.Value)
		}
	}
}

func processDecreaseQtyProduct(db *gorm.DB, message []byte) {
	var orderCreated events.OrderCreatedEvent
	err := json.Unmarshal([]byte(message), &orderCreated)

	if err != nil {
		log.Printf("error unmarshal message: %v", err.Error())
	}

	var product domain.Product
	err = db.First(&product, orderCreated.ProductID).Error
	if err != nil {
		log.Printf("product was not found (consumer) : %v", err.Error())
		return
	}

	product.Quantity -= orderCreated.Qty

	if err := db.Save(&product).Error; err != nil {
		log.Printf("failed to update product: %v", err.Error())
		return
	}

	log.Println("succesfully update product")
}
