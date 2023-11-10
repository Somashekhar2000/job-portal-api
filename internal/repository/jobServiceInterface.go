package repository

import model "project/internal/model"

//go:generate mockgen -source=jobServiceInterface.go -destination=jobServiceInterface_mock.go -package=repository
type CompanyService interface {
	CompanyCreate(newCompany model.CreateCompany) (model.Company, error)
	GetAllCompanies() ([]model.Company, error)
	GetCompanyById(id int) (model.Company, error)
	JobCreate(newJob model.NewJobRequest, id uint64) (model.Response, error)
	GetJobsByCompanyId(companyID int) ([]model.Job, error)
	GetAllJobs() ([]model.Job, error)
	GetJobByJobId(jobID int) (model.Job, error)
	ProcessJobApplications(appData []model.NewUserApplication) ([]model.NewUserApplication, error)
}
