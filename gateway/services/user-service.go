package services

import (
	"gateway/config"
	"strings"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func GetUser(c *gin.Context) {
	username, isOk := c.Params.Get("username")
	if !isOk {
		c.JSON(400, gin.H{
			"error": "Please provider username",
		})
	}
	rdbClient := GetRDBClient()
	user, err := rdbClient.Get(username)
	if err == redis.Nil {
		user = "patrickphan@gmail.com"
		rdbClient.Set(username, user, 5*time.Minute)
		//	fmt.Println("Redis client set")
	} else if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"user": user,
	})
	// fmt.Println("user", user)
}

func RegisterUser(c *gin.Context) {
	//
	kafkaProducer, err := NewKafkaProducer(config.GetKafkaBrokers())
	if err != nil {
		c.JSON(500, gin.H{
			"error": strings.Split(err.Error(), ""),
		})
	}

	topic := "gateway"

	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte("Has user register"),
	}
	err = kafkaProducer.Produce(message, nil)
	if err != nil {
		c.JSON(500, gin.H{
			"error": strings.Split(err.Error(), ""),
		})
	}
}
