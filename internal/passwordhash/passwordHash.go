package passwordhash

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func HashingPassword(password string) (string, error) {
	if password == "" {
		log.Error().Msg("error : empty password ")
		return "", errors.New("error in hashing password ")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("error in hashing password")
		return "", fmt.Errorf("error in hashing password : %w", err)
	}
	return string(hashedPassword), nil
}

func CheckingHashPassword(password string, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Error().Err(err).Msg("error wrong password")
		return errors.New("wrong password")
	}
	return nil
}

func OTPGeneration() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%v", rand.Intn(90000)+10000)
}
