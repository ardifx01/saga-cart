package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type IKafka interface {
	WriteMessages(ctx context.Context, msgs ...kafka.Message) error
	Close() error
}

type ProductPublisher struct {
	writer IKafka
}

func NewProductPublisher(writer IKafka) *ProductPublisher {
	return &ProductPublisher{
		writer: writer,
	}
}

func (p *ProductPublisher) Publish(topic string, message []byte) error {
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

func (p *ProductPublisher) Close() error {
	return p.writer.Close()
}
