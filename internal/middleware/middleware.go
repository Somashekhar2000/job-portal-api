package middleware

import (
	"fmt"
	"job-portal-api/internal/authentication"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Mid struct {
	auth authentication.Authenticaton
}

type Middleware interface {
	Authentication(next gin.HandlerFunc) gin.HandlerFunc
	Log() gin.HandlerFunc
}

func NewMid(auth authentication.Authenticaton) (Middleware, error) {
	if auth == nil {
		log.Info().Msg("authencatiomn is nil")
		return nil, fmt.Errorf("error authentication is nil")
	}
	return &Mid{
		auth: auth,
	}, nil
}
