package kafka

import (
	"context"
	"encoding/json"
	"log"
	"order_service_saga/internal/contracts"

	"github.com/segmentio/kafka-go"
)

type OrderConsumer struct {
	reader    *kafka.Reader
	orderRepo contracts.OrderRepoContract
}

func NewOrderConsumer(reader *kafka.Reader, orderRepo contracts.OrderRepoContract) *OrderConsumer {
	return &OrderConsumer{
		reader:    reader,
		orderRepo: orderRepo,
	}
}

func (c *OrderConsumer) Consume() {
	defer c.reader.Close()
	for {
		msg, err := c.reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		log.Printf("Message received (order-service): key=%s, value=%s\n", string(msg.Key), string(msg.Value))
		// "order-failed", "order-succes"

		var message_response struct {
			ID          int  `json:"id"`
			StatusOrder bool `json:"status"`
		}

		err = json.Unmarshal(msg.Value, &message_response)
		if err != nil {
			log.Fatalf("error while unmarshal messsage (order-service): %v", err.Error())
		}

		switch msg.Topic {
		case "order-failed":
			c.orderRepo.UpdateOrderStatus(message_response.ID, message_response.StatusOrder)
		case "order-succes":
			c.orderRepo.UpdateOrderStatus(message_response.ID, message_response.StatusOrder)
		}
	}
}
