package kafka_test

import (
	"errors"
	"payment_service_saga/internal/kafka"
	mock_kafka "payment_service_saga/internal/mocks/kafka"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestPaymentPublisher_Publish(t *testing.T) {
	ctrl := gomock.NewController(t)
	kafkaWriterMock := mock_kafka.NewMockIKafka(ctrl)
	paymentPublisher := kafka.NewPaymentPublisher(kafkaWriterMock)

	t.Run("Success publish data to topic", func(t *testing.T) {
		kafkaWriterMock.EXPECT().WriteMessages(gomock.Any(), gomock.Any()).Return(nil)

		err := paymentPublisher.Publish("test", []byte{})
		assert.NoError(t, err)
	})

	t.Run("Error publish data to topic", func(t *testing.T) {
		kafkaWriterMock.EXPECT().WriteMessages(gomock.Any(), gomock.Any()).Return(errors.New("error"))

		err := paymentPublisher.Publish("test", []byte{})
		assert.Error(t, err)
	})
}

func TestPaymentPublisher_Close(t *testing.T) {
	ctrl := gomock.NewController(t)
	kafkaWriterMock := mock_kafka.NewMockIKafka(ctrl)
	paymentPublisher := kafka.NewPaymentPublisher(kafkaWriterMock)

	t.Run("Success close kafka writer", func(t *testing.T) {
		kafkaWriterMock.EXPECT().Close().Return(nil)

		err := paymentPublisher.Close()
		assert.NoError(t, err)
	})

	t.Run("Error close kafka writer", func(t *testing.T) {
		kafkaWriterMock.EXPECT().Close().Return(errors.New("error"))

		err := paymentPublisher.Close()
		assert.Error(t, err)
	})
}
