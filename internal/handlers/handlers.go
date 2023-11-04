package handlers

import (
	"project/internal/auth"
	"project/internal/middlewear"
	"project/internal/services"

	"github.com/gin-gonic/gin"
)

// API creates and configures a new Gin Engine for the API.
func API(authService *auth.Auth, service *services.Service) *gin.Engine {

	router := gin.New()

	handler, _ := NewHandler(authService, service, service)

	middleware, _ := middlewear.NewMiddleWear(authService)

	// Use logging and recovery middleware
	router.Use(middleware.Log(), gin.Recovery())

	// Define API routes.
	router.POST("/api/register", handler.userSignup)
	router.POST("/api/login", handler.userLogin)
	router.POST("/api/companies", middleware.Auth(handler.companyCreation))
	router.GET("/api/companies", middleware.Auth(handler.getAllCompany))
	router.GET(" /api/company/:company_id", middleware.Auth(handler.getCompanyById))
	router.POST("/api/companies/:company_id/jobs", middleware.Auth(handler.postJobByCompany))
	router.GET("/api/companies/:company_id/jobs", middleware.Auth(handler.getJobByCompany))
	router.GET("/api/jobs", middleware.Auth(handler.getAllJob))
	router.GET("/api/jobs/:job_id", middleware.Auth(handler.getJobByJobId))

	return router
}
