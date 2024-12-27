package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type OrderConsumer struct {
	reader *kafka.Reader
}

func NewOrderConsumer(reader *kafka.Reader) *OrderConsumer {
	return &OrderConsumer{
		reader: reader,
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
		log.Printf("Message received: key=%s, value=%s\n", string(msg.Key), string(msg.Value))
	}
}
