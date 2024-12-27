package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type OrderPublisher struct {
	writer *kafka.Writer
}

func NewOrderPublisher(writer *kafka.Writer) *OrderPublisher {
	return &OrderPublisher{
		writer: writer,
	}
}

func (p *OrderPublisher) Publish(topic string, message []byte) error {
	err := p.writer.WriteMessages(context.Background(), kafka.Message{
		Topic: topic,
		Value: message,
	})
	if err != nil {
		log.Printf("error publish data (order serice): %v", err.Error())
		return err
	}
	log.Printf("Message published to topic %s: %s\n", topic, message)
	return nil
}

func (p *OrderPublisher) Close() error {
	return p.writer.Close()
}
