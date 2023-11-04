package model

import (
	//"project/internal/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName     string `json:"name"  gorm:"unique"`
	Email        string `json:"email"  gorm:"unique"`
	PasswordHash string `json:"-" validate:"required"`
}

type UserSignup struct {
	UserName string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// func (m *UserSignup) Read(p []byte) (n int, err error) {
// 	if m.pos >= len(m.data) {
// 		return 0, io.EOF
// 	}
// 	n = copy(p, m.data[m.pos:])
// 	m.pos += n
// 	return n, nil
// }
