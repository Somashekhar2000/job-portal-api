package service

import (
	"context"
	"errors"
	"fmt"
	"job-portal-api/internal/authentication"
	"job-portal-api/internal/cache"
	"job-portal-api/internal/model"
	"job-portal-api/internal/passwordhash"
	"job-portal-api/internal/repository"
	"net/smtp"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=userService.go -destination=userService_mock.go -package=service
type UserService interface {
	UserSignup(userSignup model.UserSignup) (model.User, error)
	Userlogin(userSignin model.UserLogin) (string, error)
	OTPGeneration(userdetails model.ChangePassword) (string, error)
	ValidatingOTP(userDetails model.OTPVerification) error
}

func NewUserService(userRepo repository.UserRepository, a authentication.Authenticaton, r cache.Caching) (UserService, error) {
	if userRepo == nil {
		return nil, errors.New("user Repo cannot be nil")
	}
	return &Service{
		userRepo:       userRepo,
		authentication: a,
		rdb:            r,
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

func (s *Service) OTPGeneration(userdetails model.ChangePassword) (string, error) {
	userData, err := s.userRepo.CheckUser(userdetails.EmailID)
	if err != nil {
		return "", err
	}

	otp := passwordhash.OTPGeneration()

	ctx := context.Background()
	err = s.rdb.AddOTP(ctx, userData.EmailID, otp)
	if err != nil {
		return "", err
	}

	from := "somashekharm4159@gmail.com"
	password := "dcbb xdwl zpyn aneh"

	to := userdetails.EmailID

	smtpServer := "smtp.gmail.com"
	smtpPort := 587

	message := fmt.Sprintf("Subject: Test Email\n\nOTP to reset password is : %v", otp)

	auth := smtp.PlainAuth("", from, password, smtpServer)

	smtpAddr := fmt.Sprintf("%s:%d", smtpServer, smtpPort)
	err = smtp.SendMail(smtpAddr, auth, from, []string{to}, []byte(message))
	if err != nil {
		log.Err(err).Msg("error while sending mail")
		return "", err
	}

	return otp, nil
}

func (s *Service) ValidatingOTP(userDetails model.OTPVerification) error {

	ctx := context.Background()

	val, err := s.rdb.GetOTP(ctx, userDetails.EmailID)
	if err != nil {
		return err
	}

	if val != userDetails.OTP {
		log.Error().Msg("otp miss match")
		return errors.New("OTP miss match")
	}

	if userDetails.NewPassword != userDetails.ConfirmPasswrord {
		log.Error().Msg("error password miss match")
		return errors.New("password miss match")
	}

	userData, err := s.userRepo.CheckUser(userDetails.EmailID)
	if err != nil {
		return err
	}

	hashPassword, err := passwordhash.HashingPassword(userDetails.ConfirmPasswrord)
	if err != nil {
		return err
	}

	userData.Password = hashPassword

	err = s.userRepo.UpdatePassword(userData)
	if err != nil {
		return err
	}

	return nil

}
