package services

import (
	"errors"

	"project/internal/model"
	"project/internal/repository"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

// NewService creates a new Service instance with the given user and company repositories.
type Service struct {
	userRepo    repository.Users
	companyRepo repository.Company
}

// NewService creates a new Service instance with the given user and company repositories.
func NewService(userRepo repository.Users, companyRepo repository.Company) (*Service, error) {
	if userRepo == nil {
		return nil, errors.New("user/company repository not provided")
	}

	return &Service{userRepo: userRepo, companyRepo: companyRepo}, nil

}

// UsersService is an interface for user-related operations.
//
//go:generate mockgen -source=userService.go -destination=userservice_mock.go -package=services
type UsersService interface {
	UserSignup(userSignup model.UserSignup) (model.User, error)
	Userlogin(userLogin model.UserLogin) (jwt.RegisteredClaims, error)
}

// UserSignup registers a new user with the provided user signup details.
func (s *Service) UserSignup(userSignup model.UserSignup) (model.User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userSignup.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Msg("error occurred while hashing password")
		return model.User{}, errors.New("password hashing failed")
	}

	user := model.User{UserName: userSignup.UserName, Email: userSignup.Email, PasswordHash: string(hashedPassword)}

	createdUser, err := s.userRepo.CreateUser(user)
	if err != nil {
		log.Error().Err(err).Msg("failed to create user")
		return model.User{}, errors.New("user creation failed")
	}

	return createdUser, nil

}

// UserLogin attempts to authenticate a user based on login credentials.
func (s *Service) Userlogin(userLogin model.UserLogin) (jwt.RegisteredClaims, error) {

	fetchedUser, err := s.userRepo.FetchUserByEmail(userLogin.Email)
	if err != nil {
		log.Error().Err(err).Msg("couldnot find user")
		return jwt.RegisteredClaims{}, errors.New("user login failed")
	}

	err = bcrypt.CompareHashAndPassword([]byte(fetchedUser.PasswordHash), []byte(userLogin.Password))
	if err != nil {
		log.Error().Err(err).Msg("password of user incorrect")
		return jwt.RegisteredClaims{}, errors.New("user login failed")
	}

	// Create JWT claims for the authenticated user.
	claims := jwt.RegisteredClaims{
		Issuer:    "service project",
		Subject:   strconv.FormatUint(uint64(fetchedUser.ID), 10),
		Audience:  jwt.ClaimStrings{"users"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	return claims, nil

}
