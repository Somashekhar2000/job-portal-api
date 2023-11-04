package auth

import (
	"crypto/rsa"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

// Auth represents an authentication service with RSA keys for token generation and validation.
type Auth struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

// NewAuth creates a new Auth instance with the provided private and public keys.
func NewAuth(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) (*Auth, error) {

	if privateKey == nil || publicKey == nil {
		return nil, errors.New("privateKey and publicKey cannot be nil")
	}
	return &Auth{privateKey: privateKey, publicKey: publicKey}, nil
}

// ValidateToken validates a JWT token using the public key.
func (a *Auth) ValidateToken(token string) (jwt.RegisteredClaims, error) {
	var c jwt.RegisteredClaims

	// Parse the token with the registered claims and public key.
	tkn, err := jwt.ParseWithClaims(token, &c, func(token *jwt.Token) (interface{}, error) {
		return a.publicKey, nil
	})
	if err != nil {
		return jwt.RegisteredClaims{}, fmt.Errorf("failed to parse token: %w", err)
	}

	if !tkn.Valid {
		return jwt.RegisteredClaims{}, errors.New("invalid token")
	}
	return c, nil
}

// GenerateToken generates a JWT token with the specified claims using the private key.
func (a *Auth) GenerateToken(claims jwt.RegisteredClaims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenStr, err := token.SignedString(a.privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenStr, nil
}
