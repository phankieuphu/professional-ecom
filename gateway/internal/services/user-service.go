package services

import (
	"context"
	config "gateway/configs"
	userv1 "gateway/gen/user/v1"
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
		log.Println("Error calling GetProfileUser:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get user profile",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": res.User,
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

type UserRegisterRequest struct {
	Username    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Address     string `json:"address" binding:"required"`
}

func RegisterUser(c *gin.Context) {
	//
	var user UserRegisterRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	gRPCClient := userv1.NewUserClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	registerUser, registerUserError := gRPCClient.RegisterUser(ctx, &userv1.RegisterUserRequest{
		User: &userv1.UserModels{
			Username:    user.Username,
			Password:    user.Password,
			Address:     user.Address,
			PhoneNumber: user.PhoneNumber,
			Email:       user.Email,
		},
	})
	if registerUserError != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": registerUserError.Error()})
		return
	}
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
	c.JSON(http.StatusCreated, gin.H{
		"message": registerUser.Status.Message,
	})
}

type UpdateUserRequest struct {
	Password    string `json:"password"`
	Address     string `json:"address"`
	DisplayName string `json:"display_name"`
}

func UpdateUser(c *gin.Context) {
	var user UpdateUserRequest
	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	gRPCClient := userv1.NewUserClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	updateUser, err := gRPCClient.UpdateUser(ctx, &userv1.UpdateUserRequest{

		Password:    &user.Password,
		Address:     &user.Address,
		DisplayName: &user.DisplayName,
		Username:    username,
	})
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": updateUser.Status.Message,
		"user":    updateUser.User})
}
