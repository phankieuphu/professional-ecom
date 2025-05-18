package services

import (
	"fmt"
	"log"

	config "github.com/phankieuphu/ecom-user/configs"

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

func (k *KafkaProducer) ProduceMessage(topic string, message []byte) error {
	if k.producer == nil {
		return fmt.Errorf("kafka producer is not initialized")
	}

	deliveryChan := make(chan kafka.Event)

	err := k.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
	}, deliveryChan)

	if err != nil {
		return fmt.Errorf("failed to produce message: %v", err)
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		return fmt.Errorf("failed to deliver message: %v", m.TopicPartition.Error)
	}

	log.Printf("Message produced to topic %s\n", topic)
	return nil
}
