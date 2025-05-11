package models

type User struct {
	BaseModel
	Email       string `gorm:"size:255;not null;unique" validate:"required,email" `
	Password    string `gorm:"size:255;null"`
	Address     string `gorm:"size:255; null"`
	PhoneNumber string `gorm:"type:char(10);not null"`
	Username    string `gorm:"size:50; unique, not null" validate:"required"`
	DisplayName string `gorm:"type:varchar(255);null"`
	Gender      bool   `gorm:"default:true"`
}

func (u *User) TableName() string {
	return "users"
}

type UserRegisterValidate struct {
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	Username    string `json:"username" binding:"required"`
	Address     string `json:"address" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	DisplayName string `json:"display_name"`
}
