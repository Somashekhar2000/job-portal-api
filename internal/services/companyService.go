package services

import (
	"errors"
	"project/internal/model"

	"github.com/rs/zerolog/log"
)

// CompanyService is an interface for managing companies and jobs.
//
//go:generate mockgen -source=companyService.go -destination=companyservice_mock.go -package=services
type CompanyService interface {
	CompanyCreate(newCompany model.CreateCompany) (model.Company, error)
	GetAllCompanies() ([]model.Company, error)
	GetCompanyById(id int) (model.Company, error)
	JobCreate(newJob model.CreateJob, id uint64) (model.Job, error)
	GetJobsByCompanyId(companyID int) ([]model.Job, error)
	GetAllJobs() ([]model.Job, error)
	GetJobByJobId(jobID int) (model.Job, error)
}

// CompanyCreate creates a new company based on the provided company details.
func (s *Service) CompanyCreate(newCompany model.CreateCompany) (model.Company, error) {
	company := model.Company{CompanyName: newCompany.CompanyName, Adress: newCompany.Adress, Domain: newCompany.Domain}

	// Attempt to create the company in the repository.
	createdCompany, err := s.companyRepo.CreateCompany(company)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create the company")
		return model.Company{}, errors.New("company creation failed")
	}

	return createdCompany, nil
}

// GetAllCompanies retrieves a list of all companies from the service.
func (s *Service) GetAllCompanies() ([]model.Company, error) {

	companies, err := s.companyRepo.GetAllCompany()
	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve companies")
		return nil, err
	}
	return companies, nil

}

// GetCompanyByID retrieves a company by its ID from the service.
func (s *Service) GetCompanyById(id int) (model.Company, error) {
	if id > 10 {
		return model.Company{}, errors.New("ID should be 10 or less")
	}
	company, err := s.companyRepo.GetCompany(id)
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch company")
		return model.Company{}, err
	}
	return company, nil

}

// JobCreate creates a new job using the provided job details and associated user ID.
func (s *Service) JobCreate(newJob model.CreateJob, id uint64) (model.Job, error) {

	job := model.Job{JobTitle: newJob.JobTitle, JobSalary: newJob.JobSalary, Uid: id}
	createdJob, err := s.companyRepo.CreateJob(job)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create job")
		return model.Job{}, errors.New("job creation failed")
	}

	return createdJob, nil
}

// GetJobsByCompanyID retrieves a list of jobs associated with a company by its ID.
func (s *Service) GetJobsByCompanyId(companyID int) ([]model.Job, error) {
	if companyID > 10 {
		return nil, errors.New("company ID cannot be greater than 10")
	}
	jobs, err := s.companyRepo.GetJobs(companyID)
	if err != nil {
		return nil, errors.New("failed to retrieve jobs")
	}
	return jobs, nil
}

// GetAllJobs retrieves a list of all available jobs.
func (s *Service) GetAllJobs() ([]model.Job, error) {

	allJobs, err := s.companyRepo.GetAllJobs()
	if err != nil {
		return nil, errors.New("failed to retrieve jobs")
	}
	return allJobs, nil

}

// GetJobByID retrieves a job by its unique identifier.
func (s *Service) GetJobByJobId(id int) (model.Job, error) {
	if id > 10 {
		return model.Job{}, errors.New("ID cannot be greater than 10")
	}
	job, err := s.companyRepo.GetJobsByJobId(id)
	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve job from the repository")
		return model.Job{}, err
	}
	return job, nil

}
