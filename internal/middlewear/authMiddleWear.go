package middlewear

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

const TokenIdKey key = "2"

// Auth is a middleware function that handles authentication.
func (m Middlewear) Auth(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := c.Request.Context()
		traceID, ok := ctx.Value(TraceIdKey).(string)
		if !ok {
			if !ok {
				log.Error().Msg("TraceID not a string")
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				return
			}
		}
		log.Info().Str("traceID", traceID).Msg("In authentication")

		authHeader := c.Request.Header.Get("Authorization")

		parts := strings.Split(authHeader, " ")
		// Checking the format of the Authorization header
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			// If the header format doesn't match required format, log and send an error
			err := errors.New("expected authorization header format: Bearer <token>")
			log.Error().Err(err).Str("Trace ID", traceID).Send()
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		claims, err := m.a.ValidateToken(parts[1])
		if err != nil {
			log.Info().Msg("Authentication failed")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		ctx = context.WithValue(ctx, TokenIdKey, claims)
		req := c.Request.WithContext(ctx)
		c.Request = req

		next(c)

	}
}
