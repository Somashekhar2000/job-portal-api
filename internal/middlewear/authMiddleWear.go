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

func (m Middlewear) Auth(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		traceId, ok := ctx.Value(TraceIdKey).(string)
		if !ok {
			log.Error().Msg("tokenid not string")
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusInternalServerError})
			return
		}
		log.Info().Str("traceid", "traceId").Msg("in authentication")

		authHead := c.Request.Header.Get("Authorization")

		parts := strings.Split(authHead, " ")
		// Checking the format of the Authorization header
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			// If the header format doesn't match required format, log and send an error
			err := errors.New("expected authorization header format: Bearer <token>")
			log.Error().Err(err).Str("Trace Id", traceId).Send()
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		regClaims, err := m.a.ValidateToken(parts[1])
		if err != nil {
			log.Info().Msg("Auth failed")
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusUnauthorized})
		}

		ctx = context.WithValue(ctx, TokenIdKey, regClaims)
		req := c.Request.WithContext(ctx)
		c.Request = req

		next(c)

	}
}
