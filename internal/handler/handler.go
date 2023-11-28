package handler

import (
	"fmt"
	"job-portal-api/internal/authentication"
	"job-portal-api/internal/middleware"
	"job-portal-api/internal/service"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	serviceUser    service.UserService
	serviceComapny service.ComapnyService
	serviceJob     service.JobService
}

func SetupApi(auth authentication.Authenticaton, userService service.UserService, comapnyService service.ComapnyService, jobService service.JobService) *gin.Engine {

	router := gin.New()

	mid, err := middleware.NewMid(auth)
	if err != nil {
		log.Panic("middleware are not set")
	}

	userHandler, err := NewUserHandler(userService)
	if err != nil {
		log.Panic("user handlers are not set")
	}

	companyHandler, err := NewCompanyHandler(comapnyService)
	if err != nil {
		log.Panic("company handlers are not set")
	}

	jobHandler, err := NewJobHandler(jobService)
	if err != nil {
		log.Panic("job handlers are not set")
	}

	router.Use(mid.Log(), gin.Recovery())

	router.GET("/api/check", check)

	router.POST("/api/signup", userHandler.Signup)
	router.POST("/api/login", userHandler.login)

	router.POST("/api/create_comapny", mid.Authentication(companyHandler.AddCompany))
	router.GET("/api/get_company/:id", mid.Authentication(companyHandler.ViewCompanyByID))
	router.GET("/api/get_companies", mid.Authentication(companyHandler.ViewAllComapny))

	router.POST("/api/addjob/companyID/:id", mid.Authentication(jobHandler.CreateJobByCompanyID))
	router.GET("/api/get_job_by_company_id/:id", mid.Authentication(jobHandler.ViewJobByCompanyId))
	router.GET("/api/get_job_by_job_id/:id", mid.Authentication(jobHandler.ViewJobByJobID))
	router.GET("/api/get_jobs", mid.Authentication(jobHandler.ViewAllJobs))
	router.GET("/api/process_application", mid.Authentication(jobHandler.ProcessJobApplication))

	router.GET("/api/otp_genereation", userHandler.GeneratingOTP)

	return router
}

func check(c *gin.Context) {

	time.Sleep(time.Second * 3)
	select {
	case <-c.Request.Context().Done():
		fmt.Println("user not there")
		return
	default:
		c.JSON(http.StatusOK, gin.H{"msg": "statusOk"})

	}

}
