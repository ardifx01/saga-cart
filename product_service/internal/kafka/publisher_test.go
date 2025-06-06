package kafka_test

import (
	"errors"
	"product_service_saga/internal/kafka"
	mock_kafka "product_service_saga/internal/mocks/kafka"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestProductPublisher_Publish(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockWriter := mock_kafka.NewMockIKafka(ctrl)
	publisher := kafka.NewProductPublisher(mockWriter)

	t.Run("Sucess publish kafka", func(t *testing.T) {
		mockWriter.EXPECT().WriteMessages(gomock.Any(), gomock.Any()).Return(nil)

		err := publisher.Publish("topic", make([]byte, 0))
		assert.NoError(t, err)
	})

	t.Run("Error publish kafka", func(t *testing.T) {
		mockWriter.EXPECT().WriteMessages(gomock.Any(), gomock.Any()).Return(errors.New("error"))

		err := publisher.Publish("topic", make([]byte, 0))
		assert.Error(t, err)
	})
}

func TestProductPublisher_Close(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockWriter := mock_kafka.NewMockIKafka(ctrl)
	publisher := kafka.NewProductPublisher(mockWriter)

	t.Run("Sucess close kafka", func(t *testing.T) {
		mockWriter.EXPECT().Close().Return(nil)

		err := publisher.Close()
		assert.NoError(t, err)
	})
}
