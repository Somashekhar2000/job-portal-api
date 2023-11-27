package handler

import (
	"encoding/json"
	"errors"
	"job-portal-api/internal/authentication"
	"job-portal-api/internal/middleware"
	"job-portal-api/internal/model"
	"job-portal-api/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

type JobHandler interface {
	CreateJobByCompanyID(c *gin.Context)
	ViewJobByCompanyId(c *gin.Context)
	ViewJobByJobID(c *gin.Context)
	ViewAllJobs(c *gin.Context)
	ProcessJobApplication(c *gin.Context)
}

func NewJobHandler(serviceJob service.JobService) (JobHandler, error) {
	if serviceJob == nil {
		log.Info().Msg("job service cannot be nil")
		return nil, errors.New("job service cannot be nil")
	}

	return &Handler{
		serviceJob: serviceJob,
	}, nil
}

func (h *Handler) CreateJobByCompanyID(c *gin.Context) {

	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middleware.TraceIDKey).(string)
	if !ok {
		log.Info().Msg("missing trace id")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error ": http.StatusText(http.StatusInternalServerError)})
		return
	}

	_, ok = ctx.Value(authentication.AuthKey).(jwt.RegisteredClaims)
	if !ok {
		log.Info().Str("trace Id : ", traceId).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error ": http.StatusText(http.StatusUnauthorized)})
		return
	}

	id := c.Param("id")

	cId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Error().Err(err).Str("trace id : ", traceId).Msg("error in parsing id")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error ": http.StatusText(http.StatusBadRequest)})
		return
	}

	var jobData model.NewJobs

	err = json.NewDecoder(c.Request.Body).Decode(&jobData)
	if err != nil {
		log.Error().Err(err).Str("Trace id : ", traceId).Msg("error in decoding")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error ": http.StatusText(http.StatusBadRequest)})
		return
	}

	validate := validator.New()

	err = validate.Struct(jobData)
	if err != nil {
		log.Error().Err(err).Str("tacr id : ", traceId).Msg("error in validating job")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error ": http.StatusText(http.StatusBadRequest)})
		return
	}

	jodResponse, err := h.serviceJob.CreateJobByCompanyId(jobData, uint(cId))
	if err != nil {
		log.Error().Err(err).Str("trace id :", traceId).Msg("error in job creation")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error ": http.StatusText(http.StatusInternalServerError)})
		return
	}

	c.JSON(http.StatusOK, jodResponse)

}

func (h *Handler) ViewJobByCompanyId(c *gin.Context) {

	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middleware.TraceIDKey).(string)
	if !ok {
		log.Info().Msg("missing trace id")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	_, ok = ctx.Value(authentication.AuthKey).(jwt.RegisteredClaims)
	if !ok {
		log.Info().Str("trace id : ", traceId).Msg("login unsuccessful")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}

	id := c.Param("id")

	cID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Error().Err(err).Str("trace id : ", traceId).Msg("error in parsing id")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	jobData, err := h.serviceJob.ViewJobByCompanyID(uint(cID))
	if err != nil {
		log.Error().Err(err).Str("trace id : ", traceId)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	c.JSON(http.StatusOK, jobData)
}

func (h *Handler) ViewJobByJobID(c *gin.Context) {

	ctx := c.Request.Context()

	traceID, ok := ctx.Value(middleware.TraceIDKey).(string)
	if !ok {
		log.Info().Msg("trace id missing")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	_, ok = ctx.Value(authentication.AuthKey).(jwt.RegisteredClaims)
	if !ok {
		log.Info().Str("trace id : ", traceID).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}

	id := c.Param("id")

	jID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Error().Err(err).Str("trace id : ", traceID).Msg("error invalid job id")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	jobData, err := h.serviceJob.ViewJobByJobID(uint(jID))
	if err != nil {
		log.Error().Err(err).Str("trace id : ", traceID)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	c.JSON(http.StatusOK, jobData)
}

func (h *Handler) ViewAllJobs(c *gin.Context) {

	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middleware.TraceIDKey).(string)
	if !ok {
		log.Info().Msg("error missing trace id")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	_, ok = ctx.Value(authentication.AuthKey).(jwt.RegisteredClaims)
	if !ok {
		log.Info().Str("trace id : ", traceId).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}

	jobsData, err := h.serviceJob.ViewAllJobs()
	if err != nil {
		log.Error().Err(err).Str("tracr id : ", traceId)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}
	c.JSON(http.StatusOK, jobsData)
}

func (h *Handler) ProcessJobApplication(c *gin.Context) {

	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middleware.TraceIDKey).(string)
	if !ok {
		log.Info().Msg("missing trace id")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	_, ok = ctx.Value(authentication.AuthKey).(jwt.RegisteredClaims)
	if !ok {
		log.Info().Str("tracr id : ", traceId).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}

	var applications []model.NewUserApplication

	err := json.NewDecoder(c.Request.Body).Decode(&applications)
	if err != nil {
		log.Error().Err(err).Str("trace id : ", traceId).Msg("error in decoding")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	validate := validator.New()
	// log.Debug().Interface("body", applications).Msg("request body")
	err = validate.Struct(applications)
	if err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			log.Error().Err(err).Str("trace id : ", traceId).Msg("error in validaing")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
			return
		}
	}

	jobApplication := h.serviceJob.ProcessApplication(applications)
	if jobApplication == nil {
		log.Info().Str("trace id : ", traceId).Msg("all applications rejected")
		c.JSON(http.StatusBadRequest, gin.H{"error all applications rejected ": http.StatusText(http.StatusBadRequest)})
		return
	}

	c.JSON(http.StatusOK, jobApplication)

}
