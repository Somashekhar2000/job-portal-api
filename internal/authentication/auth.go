package authentication

import (
	"crypto/rsa"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type Key int

const AuthKey Key = 1

// auth stuct with fields private and public key
type Auth struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

//go:generate mockgen -source=auth.go -destination=auth_mock.go -package=service
// authentication interface to generate and validat etoken
type Authenticaton interface {
	GenerateToken(claims jwt.RegisteredClaims) (string, error)
	ValidateToken(token string) (jwt.RegisteredClaims, error)
}

// fator function taht returns authentication having struct
func NewAuth(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) (Authenticaton, error) {
	if privateKey == nil && publicKey == nil {
		return nil, errors.New("public and private are nil")
	}

	return &Auth{
		privateKey: privateKey,
		publicKey:  publicKey,
	}, nil
}
