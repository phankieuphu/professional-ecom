package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	config "github.com/phankieuphu/ecom-user/configs"
	userPb "github.com/phankieuphu/ecom-user/gen/user/v1"
	user_constant "github.com/phankieuphu/ecom-user/internal/constant"
	"github.com/phankieuphu/ecom-user/internal/logger"
	"github.com/phankieuphu/ecom-user/internal/models"
)

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

type UserService struct {
	userPb.UnimplementedUserServer
}

func (s *UserService) GetProfileUser(ctx context.Context, req *userPb.GetProfileUserRequest) (*userPb.GetProfileUserResponse, error) {
	name := req.Name
	if name == "" {
		return &userPb.GetProfileUserResponse{
			Status: &userPb.ResponseStatus{
				Code:    user_constant.CodeMissingInformationError,
				Message: user_constant.MessageMissingInformationError,
			},
			User: nil,
		}, fmt.Errorf("missing username")
	}

	database := config.GetDb()

	user := &models.User{}

	if err := database.First(user, "username = ?", name).Error; err != nil {
		return &userPb.GetProfileUserResponse{
			Status: &userPb.ResponseStatus{
				Code:    user_constant.CodeUserNotFoundError,
				Message: user_constant.MessageUserNotFoundError,
			},
			User: nil,
		}, fmt.Errorf("user not found: %w", err)
	}
	go func() {
		kafka, err := NewKafkaProducer(config.GetKafkaBrokers())
		if err == nil {
			userCreate, err := json.Marshal(user)
			if err == nil {
				kafkaSendMessage := kafka.ProduceMessage(user_constant.KafkaTopicUserCreated, userCreate)
				if kafkaSendMessage != nil {
					fmt.Println("Error sending message to Kafka:", kafkaSendMessage)
				} else {
					fmt.Println("Message sent to Kafka successfully")
				}
			}
		}

		if err != nil {
			fmt.Println("Error creating Kafka producer:", err)
		}
	}()

	return &userPb.GetProfileUserResponse{
		Status: &userPb.ResponseStatus{
			Code:    user_constant.CodeUserFetchedSuccess,
			Message: user_constant.MessageUserFetchedSuccess,
		},
		User: &userPb.UserModels{
			Id:          user.ID.String(),
			Username:    user.Username,
			Email:       user.Email,
			Address:     user.Address,
			PhoneNumber: user.PhoneNumber,
			DisplayName: user.DisplayName,
		},
	}, nil
}
func (s *UserService) RegisterUser(ctx context.Context, req *userPb.RegisterUserRequest) (*userPb.RegisterUserResponse, error) {
	//  req models.UserRegisterValidate
	// if err := c.ShouldBindJSON(&req); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	reqUser := req.User
	fmt.Println("user", reqUser)
	errChan := make(chan error, 5)
	var wg sync.WaitGroup
	// if err := c.ShouldBindJSON(&models.User); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// }
	// hash user password with bcrypt
	user := &models.User{
		Email:       reqUser.Email,
		Password:    reqUser.Password,
		Username:    reqUser.Username,
		Address:     reqUser.Address,
		PhoneNumber: reqUser.PhoneNumber,
		DisplayName: "",
	}

	//ctx, cancel := context.WithTimeout(c.Request.Context(), 6*time.Second)
	// defer cancel()
	database := config.GetDb()

	userSaved := database.Save(user)
	fmt.Println("userSaved", userSaved)
	if userSaved.Error != nil {
		return nil, userSaved.Error
	}
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	if err := sendEmail(ctx, user); err != nil {
	// 		errChan <- fmt.Errorf("log error: %w", err)
	// 	}
	// }()

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
		return nil, fmt.Errorf("encountered errors: %s", strings.Join(errors, "; "))

	}

	return &userPb.RegisterUserResponse{Id: user.ID.String(), Status: &userPb.ResponseStatus{
		Code:    user_constant.CodeUserCreatedSuccess,
		Message: user_constant.MessageUserCreatedSuccess,
	}}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *userPb.UpdateUserRequest) (*userPb.UpdateUserResponse, error) {
	username := req.Username
	if username == "" {
		return &userPb.UpdateUserResponse{
			Status: &userPb.ResponseStatus{
				Code:    user_constant.CodeMissingInformationError,
				Message: user_constant.MessageMissingInformationError,
			},
		}, fmt.Errorf("missing user username")
	}
	database := config.GetDb()
	user := &models.User{}

	var mu sync.Mutex // Mutex to handle race condition

	mu.Lock()
	defer mu.Unlock()

	if err := database.First(user, "username = ?", username).Error; err != nil {
		return &userPb.UpdateUserResponse{
			Status: &userPb.ResponseStatus{
				Code:    user_constant.CodeUserNotFoundError,
				Message: user_constant.MessageUserNotFoundError,
			},
		}, fmt.Errorf("user not found: %v", err)
	}

	// Only update fields if they are provided (not nil)
	updateFields := make(map[string]interface{})

	if req.Address != nil && *req.Address != "" {
		updateFields["address"] = *req.Address
	}
	if req.DisplayName != nil && *req.DisplayName != "" {
		updateFields["display_name"] = *req.DisplayName
	}
	if req.Password != nil && *req.Password != "" {
		updateFields["password"] = *req.Password
	}

	if len(updateFields) > 0 {
		if err := database.Model(user).Where("username = ?", username).Updates(updateFields).Error; err != nil {
			return &userPb.UpdateUserResponse{
				Status: &userPb.ResponseStatus{
					Code:    user_constant.CodeUserNotFoundError,
					Message: "failed to update user",
				},
			}, fmt.Errorf("failed to update user: %v", err)
		}
	}

	// delete cache user if exits

	return &userPb.UpdateUserResponse{
		Status: &userPb.ResponseStatus{
			Code:    user_constant.CodeUserUpdatedSuccess,
			Message: user_constant.MessageUserUpdatedSuccess,
		},
	}, nil
}
