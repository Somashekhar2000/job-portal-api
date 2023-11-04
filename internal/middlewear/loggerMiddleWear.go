package middlewear

import (
	"context"
	"project/internal/auth"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type Middlewear struct {
	a *auth.Auth
}

func NewMiddleWear(a *auth.Auth) (Middlewear, error) {
	return Middlewear{a: a}, nil
}

type key string

const TraceIdKey key = "1"

func (m Middlewear) Log() gin.HandlerFunc {

	return func(c *gin.Context) {
		traceId := uuid.NewString()
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, TraceIdKey, traceId)
		req := c.Request.WithContext(ctx)
		c.Request = req

		log.Info().Str("traceId", traceId).Msg("in log file")
		defer log.Logger.Info().Msg("request processing complete")
		c.Next()

	}

}
