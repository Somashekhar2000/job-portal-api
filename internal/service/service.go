package service

import (
	"job-portal-api/internal/authentication"
	"job-portal-api/internal/cache"
	"job-portal-api/internal/repository"
)

type Service struct {
	userRepo       repository.UserRepository
	comapnayRepo   repository.ComapnyRepo
	jobRepo        repository.JobRepository
	authentication authentication.Authenticaton
	rdb cache.Caching
}
