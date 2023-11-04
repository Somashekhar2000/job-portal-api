package model

import (
	"gorm.io/gorm"
)

// User represents a user in the system.
type User struct {
	gorm.Model
	UserName     string `json:"name"  gorm:"unique"`
	Email        string `json:"email"  gorm:"unique"`
	PasswordHash string `json:"-" validate:"required"`
}

// UserSignup represents user registration data.
type UserSignup struct {
	UserName string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// UserLogin represents user login data.
type UserLogin struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
