package middleware

import (
	"context"
	"errors"
	"job-portal-api/internal/authentication"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Key string

const TraceIDKey Key = "1"

func (m *Mid) Authentication(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := c.Request.Context()

		traceID, ok := ctx.Value(TraceIDKey).(string)

		if !ok {
			log.Info().Msg("traceID is not present in the context")
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error ": http.StatusText(http.StatusInternalServerError)})
			return
		}

		authHeader := c.Request.Header.Get("Authorization")
		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			err := errors.New("authorization header formate is invalid no proper header : bearer <token>")
			log.Error().Err(err).Str("trace id : ", traceID).Send()
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error ": http.StatusText(http.StatusInternalServerError)})
			return
		}

		claims, err := m.auth.ValidateToken(parts[1])
		if err != nil {
			log.Error().Err(err).Str("Trace id : ", traceID).Send()
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error ": http.StatusText(http.StatusInternalServerError)})
			return
		}

		ctx = context.WithValue(ctx, authentication.AuthKey, claims)

		req := c.Request.WithContext(ctx)

		c.Request = req
		next(c)

	}

}
