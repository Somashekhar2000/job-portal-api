package service

import (
	"errors"
	"job-portal-api/internal/authentication"
	"job-portal-api/internal/model"
	"job-portal-api/internal/passwordhash"
	"job-portal-api/internal/repository"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

//go:generate mockgen -source=userService.go -destination=userService_mock.go -package=service
type UserService interface {
	UserSignup(userSignup model.UserSignup) (model.User, error)
	Userlogin(userSignin model.UserLogin) (string, error)
	otpGeneration(userdetails model.ChangePassword) (string, error)
}

func NewUserService(userRepo repository.UserRepository, a authentication.Authenticaton) (UserService, error) {
	if userRepo == nil {
		return nil, errors.New("user Repo cannot be nil")
	}
	return &Service{
		userRepo:       userRepo,
		authentication: a,
	}, nil
}

func (s *Service) UserSignup(userData model.UserSignup) (model.User, error) {
	hashedPassword, err := passwordhash.HashingPassword(userData.Password)
	if err != nil {
		return model.User{}, err
	}

	userDetails := model.User{
		UserName: userData.UserName,
		EmailID:  userData.EmailID,
		Password: hashedPassword,
	}

	userDetails, err = s.userRepo.CreateUser(userDetails)
	if err != nil {
		return model.User{}, err
	}

	return userDetails, nil

}

func (s *Service) Userlogin(userSignin model.UserLogin) (string, error) {

	userData, err := s.userRepo.CheckUser(userSignin.EmailID)
	if err != nil {
		return "", err
	}

	err = passwordhash.CheckingHashPassword(userSignin.Password, userData.Password)
	if err != nil {
		return "", err
	}

	claims := jwt.RegisteredClaims{
		Issuer:    "job portal project",
		Subject:   strconv.FormatUint(uint64(userData.ID), 10),
		Audience:  jwt.ClaimStrings{"users"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token, err := s.authentication.GenerateToken(claims)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) otpGeneration(userdetails model.ChangePassword) (string, error) {
	userData, err := s.userRepo.CheckUser(userdetails.EmailID)
	if err != nil {
		return "", err
	}

}
