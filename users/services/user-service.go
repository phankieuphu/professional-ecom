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
	"users/models"

	"github.com/gin-gonic/gin"
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
	for {
		select {
		case <-time.After(5 * time.Second):
			// time.Sleep(5 * time.Second)
			if u.Email != "" {
				fmt.Println("Not correct user")
				return errors.New("not correct user")
			}
			return nil
		case <-ctx.Done():
			return fmt.Errorf("email sending canceled: %v", ctx.Err())
		}
	}
}
func writeToLogService(u *models.User) error {
	fmt.Println("write log to service")
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
