package services

import (
	"errors"
	"project/internal/model"

	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=companyService.go -destination=companyservice_mock.go -package=services
type CompanyService interface {
	CompanyCreate(nc model.CreateCompany) (model.Company, error)
	GetAllCompanies() ([]model.Company, error)
	GetCompanyById(id int) (model.Company, error)
	JobCreate(nj model.CreateJob, id uint64) (model.Job, error)
	GetJobsByCompanyId(id int) ([]model.Job, error)
	GetAllJobs() ([]model.Job, error)
	GetJobByJobId(id int) (model.Job, error)
}

func (s *Service) CompanyCreate(nc model.CreateCompany) (model.Company, error) {
	company := model.Company{CompanyName: nc.CompanyName, Adress: nc.Adress, Domain: nc.Domain}
	cu, err := s.c.CreateCompany(company)
	if err != nil {
		log.Error().Err(err).Msg("couldnot create company")
		return model.Company{}, errors.New("company creation failed")
	}

	return cu, nil
}

func (s *Service) GetAllCompanies() ([]model.Company, error) {

	AllCompanies, err := s.c.GetAllCompany()
	if err != nil {
		log.Error().Err(err).Msg("couldnot get companies")
		return nil, err
	}
	return AllCompanies, nil

}

func (s *Service) GetCompanyById(id int) (model.Company, error) {
	if id > 10 {
		return model.Company{}, errors.New("id cannnot be greater")
	}
	Companies, err := s.c.GetCompany(id)
	if err != nil {
		log.Error().Err(err).Msg("couldnot get company")
		return model.Company{}, err
	}
	return Companies, nil

}

func (s *Service) JobCreate(nj model.CreateJob, id uint64) (model.Job, error) {
	job := model.Job{JobTitle: nj.JobTitle, JobSalary: nj.JobSalary, Uid: id}
	cu, err := s.c.CreateJob(job)
	if err != nil {
		log.Error().Err(err).Msg("couldnot create job")
		return model.Job{}, errors.New("job creation failed")
	}

	return cu, nil
}

func (s *Service) GetJobsByCompanyId(id int) ([]model.Job, error) {
	if id > 10 {
		return nil, errors.New("id cannnot be greater")
	}
	AllCompanies, err := s.c.GetJobs(id)
	if err != nil {
		return nil, errors.New("job retreval failed")
	}
	return AllCompanies, nil
}

func (s *Service) GetAllJobs() ([]model.Job, error) {

	AllJobs, err := s.c.GetAllJobs()
	if err != nil {
		return nil, err
	}
	return AllJobs, nil

}

func (s *Service) GetJobByJobId(id int) (model.Job, error) {
	if id > 10 {
		return model.Job{}, errors.New("id cannnot be greater")
	}
	Companies, err := s.c.GetJobsByJobId(id)
	if err != nil {
		log.Error().Err(err).Msg("couldnot get company")
		return model.Job{}, err
	}
	return Companies, nil

}
