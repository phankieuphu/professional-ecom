package services

import (
	"fmt"
	"gateway/config"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaProducer struct {
	producer *kafka.Producer
}

func KafkaConsumer() {
	c, err := kafka.NewConsumer(config.LoadKafkaConsumerConfig())
	if err != nil {
		log.Fatalf("Error connect to kafka")
	}
	defer c.Close()
	listTopics := config.LoadListKafkaConsumerTopic()
	c.SubscribeTopics(listTopics, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Received: %s\n", string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

func NewKafkaProducer(brokers string) (*KafkaProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
	})
	if err != nil {
		return nil, err
	}

	return &KafkaProducer{
		producer: p,
	}, nil
}
