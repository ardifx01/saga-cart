package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type PaymentPublisher struct {
	writer *kafka.Writer
}

func NewPaymentPublisher(writer *kafka.Writer) *PaymentPublisher {
	return &PaymentPublisher{
		writer: writer,
	}
}

func (p *PaymentPublisher) Publish(topic string, message []byte) error {
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

func (p *PaymentPublisher) Close() error {
	return p.writer.Close()
}
