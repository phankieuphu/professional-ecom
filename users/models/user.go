package models

import "github.com/google/uuid"

type User struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	Email       string    `gorm:"size:255;not null;unique" validate:"required,email"`
	Password    string    `gorm:"size:255;null"`
	Address     string    `gorm:"size:255; null"`
	PhoneNumber string    `gorm:"type:char(10);not null"`
	Username    string    `gorm:"size:50; unique, not null" validate:"required"`
}

func (u *User) TableName() string {
	return "user"
}
