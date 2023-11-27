package service

import (
	"errors"
	"job-portal-api/internal/model"
	"job-portal-api/internal/repository"

	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=companyService.go -destination=companyService_mock.go -package=service
type ComapnyService interface {
	AddingCompany(company model.AddCompany) (model.Company, error)
	ViewCompanyById(Id uint64) (model.Company, error)
	ViewAllCompanies() ([]model.Company, error)
}

func NewCompanyService(comapnyRepo repository.ComapnyRepo) (ComapnyService, error) {
	if comapnyRepo == nil {
		log.Info().Msg("comapny service cannot be nil")
		return nil, errors.New("company service cannot be nil")
	}
	return &Service{
		comapnayRepo: comapnyRepo,
	}, nil
}

func (s *Service) AddingCompany(company model.AddCompany) (model.Company, error) {

	companyData := model.Company{
		CompanyName: company.CompanyName,
		Address:     company.Address,
		Domain:      company.Domain,
	}

	companyData, err := s.comapnayRepo.CreateComapny(companyData)
	if err != nil {
		return model.Company{}, err
	}

	return companyData, nil
}

func (s *Service) ViewCompanyById(cId uint64) (model.Company, error) {
	companyData, err := s.comapnayRepo.GetCompanyByID(cId)
	if err != nil {
		return model.Company{}, err
	}
	return companyData, nil
}

func (s *Service) ViewAllCompanies() ([]model.Company, error) {

	companiesData, err := s.comapnayRepo.GetAllCompanies()
	if err != nil {
		return nil, errors.ErrUnsupported

	}

	return companiesData, nil
}
