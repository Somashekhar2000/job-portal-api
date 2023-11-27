package repository

import (
	"errors"
	"job-portal-api/internal/model"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

//go:generate mockgen -source=userRepository.go -destination=userRepository_mock.go -package=repository
type UserRepository interface {
	CreateUser(userData model.User) (model.User, error)
	CheckUser(email string) (model.User, error)
}

func NewUserRepo(db *gorm.DB) (UserRepository, error) {
	if db == nil {
		return nil, errors.New("database cannot be nil")
	}
	return &Repo{
		db: db,
	}, nil
}

func (r *Repo) CreateUser(userData model.User) (model.User, error) {

	output := r.db.Create(&userData)
	if output.Error != nil {
		log.Error().Err(output.Error).Msg("error in database, could not create user")
	}

	return userData, nil
}

func (r *Repo) CheckUser(email string) (model.User, error) {

	var userData model.User

	data := r.db.Where("email_id = ?", email).First(&userData)

	if data.Error != nil {
		log.Error().Err(data.Error).Msg("error email not found in database")
		return model.User{}, errors.New("error email not found")
	}

	return userData, nil
}
