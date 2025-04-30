package test

import (
	"gateway/services"
	"testing"
)

func TestNewKafkaProducer(t *testing.T) {
	broker := "localhost:9092"

	producer, err := services.NewKafkaProducer(broker)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if producer == nil {
		t.Fatal("Expected producer to be non-nil")
	}

	producer.Close()
}
