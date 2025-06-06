package kafka_test

import (
	"errors"
	"order_service_saga/internal/kafka"
	mock_kafka "order_service_saga/internal/mocks/kafka"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestOrderPublisher_Publish(t *testing.T) {
	ctlr := gomock.NewController(t)
	writerMock := mock_kafka.NewMockIKafkaWriter(ctlr)
	pub := kafka.NewOrderPublisher(writerMock)

	t.Run("Sucess publish data to topic", func(t *testing.T) {
		writerMock.EXPECT().WriteMessages(gomock.Any(), gomock.Any()).Return(nil)

		err := pub.Publish("test", []byte{})
		assert.NoError(t, err)
	})

	t.Run("Error publish data to topic", func(t *testing.T) {
		writerMock.EXPECT().WriteMessages(gomock.Any(), gomock.Any()).Return(errors.New("error"))

		err := pub.Publish("test", []byte{})
		assert.Error(t, err)
	})
}

func TestOrderPublisher_Close(t *testing.T) {
	ctlr := gomock.NewController(t)
	writerMock := mock_kafka.NewMockIKafkaWriter(ctlr)
	pub := kafka.NewOrderPublisher(writerMock)

	t.Run("Success close kafka writer", func(t *testing.T) {
		writerMock.EXPECT().Close().Return(nil)
		err := pub.Close()
		assert.NoError(t, err)
	})

	t.Run("Failed close kafka writer", func(t *testing.T) {
		writerMock.EXPECT().Close().Return(errors.New("error"))
		err := pub.Close()
		assert.Error(t, err)
	})
}
