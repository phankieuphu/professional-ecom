package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
	"users/config"
	"users/logger"
	"users/models"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegisterRequest struct {
	Email string `json:"email" binding:"required,email"`
}

func RegisterUser(c *gin.Context) {
	var req models.UserRegisterValidate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	errChan := make(chan error, 5)
	var wg sync.WaitGroup
	// if err := c.ShouldBindJSON(&models.User); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// }
	// hash user password with bcrypt
	user := &models.User{
		Email:       req.Email,
		Password:    req.Password,
		Username:    req.Username,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
		DisplayName: req.DisplayName,
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 6*time.Second)
	defer cancel()
	database := config.GetDb()

	userSaved := database.Save(user)
	fmt.Println("userSaved", userSaved)
	if userSaved.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": userSaved.Error.Error(),
		})
		return
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := sendEmail(ctx, user); err != nil {
			errChan <- fmt.Errorf("log error: %w", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := writeToLogService(user); err != nil {
			errChan <- fmt.Errorf("write log service error %w", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := cacheUser(ctx, user); err != nil {
			errChan <- fmt.Errorf("error cache user %w", err)

		}
	}()
	go func() {
		wg.Wait()
		close(errChan)
	}()
	var errors []string
	for err := range errChan {
		if err != nil {
			errors = append(errors, err.Error())
		}
	}

	if len(errors) > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "partial failure", "errors": errors})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "registered", "user": user.ID})
}

func sendEmail(ctx context.Context, u *models.User) error {
	logger := logger.NewConsoleLogger()

	for {
		select {
		case <-time.After(5 * time.Second):
			// time.Sleep(5 * time.Second)
			if u.Email != "" {
				logger.Info("email service", "user", u)
			}
			logger.Info("Send email to user success", "user", u)
			return nil
		case <-ctx.Done():
			return fmt.Errorf("email sending canceled: %v", ctx.Err())
		}
	}
}
func writeToLogService(u *models.User) error {
	logger := logger.NewConsoleLogger()
	logger.Info("log service", "user", u)
	return nil
}

func cacheUser(ctx context.Context, u *models.User) error {
	fmt.Println("set cache to service")
	for {
		select {
		case <-time.After(2 * time.Second):
			log.Printf("[cache] User cached: %d", u.ID)
			return nil
		case <-ctx.Done():
			return fmt.Errorf("cache canceled: %v", ctx.Err())

		}
	}
}

func GetUserProfile(c *gin.Context) {
	logger := logger.NewConsoleLogger()
	// who is view profile
	// get request user
	username, isOk := c.Params.Get("username")
	if !isOk {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid params",
		})
	}
	database := config.GetDb()

	var user models.User
	if err := database.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "User not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// send notification to user
	topic := "ProfileViewedEvent"
	kafkaProducer, err := NewKafkaProducer(config.GetKafkaBrokers())
	if err != nil {

	}
	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(fmt.Sprintf("User %s viewed profile of %s", username, user.Username)),
	}

	if err = kafkaProducer.Produce(message, nil); err != nil {
		logger.Error("Send message to kafka failed", user)
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
