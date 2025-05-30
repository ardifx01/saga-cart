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
	reader    *kafka.Reader
	publisher *ProductPublisher
	db        *gorm.DB
}

func NewProductConsumer(reader *kafka.Reader, publisher *ProductPublisher, db *gorm.DB) *ProductConsumer {
	return &ProductConsumer{
		reader:    reader,
		publisher: publisher,
		db:        db,
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
		case "stock-reserved":
			log.Printf("Message received (stock-reserved): value=%s\n", string(msg.Value))
			processDecreaseQtyProduct(c.db, c.publisher, msg.Value)
		case "stock-failed":
			log.Printf("Message received (stock-failed): value=%s\n", string(msg.Value))

			var stockFailed struct {
				ProductID  int `json:"product_id"`
				JumlahBeli int `json:"jumlah_beli"`
			}
			json.Unmarshal(msg.Value, &stockFailed)

			var product domain.Product
			err = c.db.First(&product, stockFailed.ProductID).Error
			if err != nil {
				log.Printf("product was not found (consumer) : %v", err.Error())
				return
			}

			product.Quantity += stockFailed.JumlahBeli

			if err := c.db.Save(&product).Error; err != nil {
				log.Printf("failed to update product: %v", err.Error())
				return
			}

			log.Println("succesfully update product stock (stock-failed)")
		}
	}
}

func processDecreaseQtyProduct(db *gorm.DB, publisher *ProductPublisher, message []byte) {
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

	pub_msg := struct {
		ID           int     `json:"id"`
		ProductID    int     `json:"product_id"`
		Amount       float64 `json:"amount"`
		HargaProduct float64 `json:"harga_product"`
		Qty          int     `json:"qty"`
		JumlahBeli   int     `json:"jumlah_beli"`
	}{
		ID:           orderCreated.ID,
		ProductID:    orderCreated.ProductID,
		Amount:       orderCreated.Amount,
		HargaProduct: product.Price,
		Qty:          product.Quantity,
		JumlahBeli:   orderCreated.Qty,
	}

	pub_msg_payload, err := json.Marshal(pub_msg)
	if err != nil {
		log.Fatalf("error marshal publish message: %v", err.Error())
	}

	publisher.Publish("payment-request", pub_msg_payload)

	log.Println("succesfully update product")
}
