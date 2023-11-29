package handler

import (
	"encoding/json"
	"errors"
	"job-portal-api/internal/middleware"
	"job-portal-api/internal/model"
	"job-portal-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=userHandler.go -destination=.mock/userHandler_mock.go -package=handler
type UserHandler interface {
	Signup(c *gin.Context)
	login(c *gin.Context)
	GeneratingOTP(c *gin.Context)
	VerifyOTP(c *gin.Context)
}

func NewUserHandler(serviceUser service.UserService) (UserHandler, error) {
	if serviceUser == nil {
		return nil, errors.New("userService Cannot be nil")
	}
	return &Handler{
		serviceUser: serviceUser,
	}, nil
}

func (h *Handler) Signup(c *gin.Context) {

	ctx := c.Request.Context()

	traceID, ok := ctx.Value(middleware.TraceIDKey).(string)
	if !ok {
		log.Info().Msg("missing Trace Id in context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error ": http.StatusText(http.StatusInternalServerError)})
		return
	}

	var userData model.UserSignup

	err := json.NewDecoder(c.Request.Body).Decode(&userData)
	if err != nil {
		log.Error().Err(err).Str("trace ID : ", traceID).Msg("error in decoding signup struct")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error ": http.StatusText(http.StatusBadRequest)})
		return
	}

	validate := validator.New()
	err = validate.Struct(userData)
	if err != nil {
		log.Error().Err(err).Str("trace ID :", traceID).Msg("error in validating sigup struct")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error ": http.StatusText(http.StatusBadRequest)})
		return
	}

	userdata, err := h.serviceUser.UserSignup(userData)
	if err != nil {
		log.Error().Err(err).Str("trace ID :", traceID).Msg("error in user sigup")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error ": http.StatusText(http.StatusInternalServerError)})
		return
	}

	c.JSON(http.StatusOK, userdata)

}

func (h *Handler) login(c *gin.Context) {
	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middleware.TraceIDKey).(string)
	if !ok {
		log.Info().Msg("missing trace ID")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error ": http.StatusText(http.StatusInternalServerError)})
		return
	}

	var userData model.UserLogin

	err := json.NewDecoder(c.Request.Body).Decode(&userData)
	if err != nil {
		log.Error().Err(err).Str("trace Id :", traceId).Msg("error in decoding")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error ": http.StatusText(http.StatusBadRequest)})
		return
	}

	validate := validator.New()
	err = validate.Struct(userData)
	if err != nil {
		log.Error().Err(err).Str("trace ID :", traceId).Msg("error in validating")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error ": http.StatusText(http.StatusBadRequest)})
		return
	}

	token, err := h.serviceUser.Userlogin(userData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error ": http.StatusText(http.StatusBadRequest)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token ": token})
}

func (h *Handler) GeneratingOTP(c *gin.Context) {
	ctx := c.Request.Context()

	traceID, ok := ctx.Value(middleware.TraceIDKey).(string)
	if !ok {
		log.Error().Err(errors.New("missing trace id")).Msg("error missing trace id")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	var forgotPassword model.ChangePassword

	err := json.NewDecoder(c.Request.Body).Decode(&forgotPassword)
	if err != nil {
		log.Err(err).Str("trace id :", traceID).Msg("error in decoding")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	validate := validator.New()

	err = validate.Struct(&forgotPassword)
	if err != nil {
		log.Err(err).Str("trace id :", traceID).Msg("error in validating")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	otp, err := h.serviceUser.OTPGeneration(forgotPassword)
	if err != nil || otp == "" {
		log.Error().Err(err).Msg("===============")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	c.JSON(http.StatusOK, otp)

}

func (h *Handler) VerifyOTP(c *gin.Context) {

	ctx := c.Request.Context()

	traceID, ok := ctx.Value(middleware.TraceIDKey).(string)
	if !ok {
		log.Err(errors.New("missing trace id")).Msg("add the missing trace id")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	var userData model.OTPVerification

	err := json.NewDecoder(c.Request.Body).Decode(&userData)
	if err != nil {
		log.Err(err).Str("trace id :", traceID).Msg("error in decoding the json body")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	validate := validator.New()

	err = validate.Struct(&userData)
	if err != nil {
		log.Err(err).Str("trace id :", traceID).Msg("error in validating the json body")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	err = h.serviceUser.ValidatingOTP(userData)
	if err != nil {
		log.Err(err).Str("trace id :", traceID).Msg("error validating otp")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"pass word reset ": " successful"})
}
