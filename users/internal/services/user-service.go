package services

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	config "github.com/phankieuphu/ecom-user/configs"
	userPb "github.com/phankieuphu/ecom-user/gen/user/v1"
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
	fmt.Println("req", req)
	return &userPb.GetProfileUserResponse{
		Message: "Hello " + req.GetName(),
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

	return &userPb.RegisterUserResponse{ID: user.ID.String(), Message: "Success"}, nil
}
