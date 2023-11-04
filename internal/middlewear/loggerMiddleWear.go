package middlewear

import (
	"context"
	"project/internal/auth"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

// Middleware represents the middleware structure.
type Middlewear struct {
	a *auth.Auth
}

// NewMiddleware creates a new instance of Middleware.
func NewMiddleWear(a *auth.Auth) (Middlewear, error) {
	return Middlewear{a: a}, nil
}

type contextKey string

const TraceIdKey contextKey = "1"

// Log returns a Gin middleware for logging.
func (m Middlewear) Log() gin.HandlerFunc {

	return func(c *gin.Context) {
		// Generate a new trace ID.
		traceID := uuid.NewString()
		ctx := c.Request.Context()

		// Store the trace ID in the context.
		ctx = context.WithValue(ctx, TraceIdKey, traceID)
		req := c.Request.WithContext(ctx)
		c.Request = req

		log.Info().Str("traceID", traceID).Msg("Logging in the file")
		defer log.Logger.Info().Msg("Request processing complete")
		c.Next()

	}

}
