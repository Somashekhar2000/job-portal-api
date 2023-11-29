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

type OTPVerification struct {
	EmailID          string `json:"emailID" validate:"required"`
	OTP              string `json:"otp" validate:"required"`
	NewPassword      string `json:"new_password" validate:"required"`
	ConfirmPasswrord string `json:"confirm_password" validate:"required"`
}
