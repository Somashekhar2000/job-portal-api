package handlers

import (
	"project/internal/auth"
	"project/internal/middlewear"
	"project/internal/services"

	"github.com/gin-gonic/gin"
)

func Api(a *auth.Auth, s *services.Service) *gin.Engine {
	r := gin.New()
	h, _ := NewHandler(a, s, s)
	//h := handler{a: a, us: s, cs: s}
	m, _ := middlewear.NewMiddleWear(a)
	r.Use(m.Log(), gin.Recovery())
	r.POST("/api/register", h.userSignup)
	r.POST("/api/login", h.userLogin)
	r.POST("/api/companies", m.Auth(h.companyCreation))
	r.GET("/api/companies", m.Auth(h.getAllCompany))
	r.GET(" /api/company/:company_id", m.Auth(h.getCompanyById))
	r.POST("/api/companies/:company_id/jobs", m.Auth(h.postJobByCompany))
	r.GET("/api/companies/:company_id/jobs", m.Auth(h.getJobByCompany))
	r.GET("/api/jobs", m.Auth(h.getAllJob))
	r.GET("/api/jobs/:job_id", m.Auth(h.getJobByJobId))
	return r
}
