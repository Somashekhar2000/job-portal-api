package authentication

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

// func to generate token
func (a *Auth) GenerateToken(claims jwt.RegisteredClaims) (string, error) {
	//create new token
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	//signing token with private key
	token, err := tkn.SignedString(a.privateKey)
	if err != nil {
		log.Info().Msg("error in generating token")
		return "", fmt.Errorf("error in generating token : %w", err)
	}
	return token, nil
}

// func to validate the given token
func (a *Auth) ValidateToken(token string) (jwt.RegisteredClaims, error) {

	var rc jwt.RegisteredClaims

	//parse the token with registered claims
	tkn, err := jwt.ParseWithClaims(token, &rc, func(t *jwt.Token) (interface{}, error) {
		return a.publicKey, nil
	})

	if err != nil {
		log.Info().Msg("error in parsing the token")
		return jwt.RegisteredClaims{}, fmt.Errorf("error in parsing token : %w ", err)
	}

	//check if token valid or not
	if !tkn.Valid {
		log.Info().Msg("token invalid")
		return jwt.RegisteredClaims{}, errors.New("invalid token")
	}

	return rc, nil
}
