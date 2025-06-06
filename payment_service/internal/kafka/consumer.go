package kafka

import (
	"context"
	"encoding/json"
	"log"
	"payment_service_saga/events"

	"github.com/segmentio/kafka-go"
)

type IKafkaConsumer interface {
	ReadMessage(ctx context.Context) (kafka.Message, error)
	Close() error
}

type IPaymentPublisher interface {
	Publish(topic string, message []byte) error
}

type PaymentConsumer struct {
	reader    IKafkaConsumer
	publisher IPaymentPublisher
}

func NewPaymentConsumer(reader IKafkaConsumer, publisher IPaymentPublisher) *PaymentConsumer {
	return &PaymentConsumer{
		reader:    reader,
		publisher: publisher,
	}
}

func (c *PaymentConsumer) Consume() {
	defer c.reader.Close()
	for {
		msg, err := c.reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		switch msg.Topic {
		case "payment-request":
			log.Printf("Message received (payment-request): key=%s, value=%s\n", string(msg.Key), string(msg.Value))

			var stockReservedEvent events.StockReservedEvent

			err := json.Unmarshal(msg.Value, &stockReservedEvent)

			if err != nil {
				log.Fatalf("error unmarshalling data consumer (payment-service): %v", err.Error())
			}

			statusOrdered := struct {
				ID          int  `json:"id"`
				StatusOrder bool `json:"status"`
			}{
				ID:          stockReservedEvent.ID,
				StatusOrder: stockReservedEvent.Amount < stockReservedEvent.HargaProduct,
			}

			status_msg, _ := json.Marshal(statusOrdered)

			if stockReservedEvent.Amount < stockReservedEvent.HargaProduct {
				fail_msg, _ := json.Marshal(stockReservedEvent)

				c.publisher.Publish("order-failed", status_msg)
				c.publisher.Publish("stock-failed", fail_msg)
			} else {
				c.publisher.Publish("order-succes", status_msg)
			}
		}
	}
}
