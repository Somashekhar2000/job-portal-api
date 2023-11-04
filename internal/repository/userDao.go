package repository

import (
	"errors"
	"project/internal/model"

	"gorm.io/gorm"
)

// Repo is a repository that interacts with the database.
type Repo struct {
	db *gorm.DB
}

// NewRepo creates a new repository with the provided database connection.
func NewRepo(db *gorm.DB) (*Repo, error) {
	if db == nil {
		return nil, errors.New("database connection is not provided")
	}

	return &Repo{db: db}, nil

}

// Users is an interface for user-related operations.

//go:generate mockgen -source=userDao.go -destination=userrepository_mock.go -package=repository
type Users interface {
	CreateUser(user model.User) (model.User, error)
	FetchUserByEmail(email string) (model.User, error)
}

// CreateUser creates a new user in the database.
func (r *Repo) CreateUser(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// FetchUserByEmail retrieves a user from the database by their email.
func (r *Repo) FetchUserByEmail(email string) (model.User, error) {
	var user model.User
	tx := r.db.Where("email=?", email).First(&user)
	if tx.Error != nil {
		return model.User{}, nil
	}
	return user, nil

}
