package model

import "gorm.io/gorm"

type UserSignup struct {
	UserName string `json:"username" validate:"required"`
	EmailID  string `json:"emailID" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	gorm.Model
	UserName string `json:"username"`
	EmailID  string `json:"emailID" gorm:"unique"`
	Password string `json:"-"`
}

type UserLogin struct {
	EmailID  string `json:"emailID" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ChangePassword struct {
	EmailID string `json:"emailID" validate:"required"`
	DOB     string `json:"dob"`
}
