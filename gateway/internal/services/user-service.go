package services

import (
	"context"
	config "gateway/configs"
	userv1 "gateway/gen/user"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func GetUser(c *gin.Context) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	grpcClient := userv1.NewUserClient(conn)

	res, err := grpcClient.GetProfileUser(ctx, &userv1.GetProfileUserRequest{Name: c.Param("username")})
	if err != nil {
		log.Fatalf("error calling FetchUser: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"user": res,
	})
	// rdbClient := GetRDBClient()
	// user, err := rdbClient.Get(username)
	// if err == redis.Nil {
	// 	user = "patrickphan@gmail.com"
	// 	rdbClient.Set(username, user, 5*time.Minute)
	// 	//	fmt.Println("Redis client set")
	// } else if err != nil {
	// 	c.JSON(500, gin.H{
	// 		"error": err.Error(),
	// 	})
	// }
	// c.JSON(200, gin.H{
	// 	"user": user,
	// })
	// fmt.Println("user", user)

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()

	// resp, err := clien
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
