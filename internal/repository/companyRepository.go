package repository

import (
	"errors"
	"job-portal-api/internal/model"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

//go:generate mockgen -source=companyRepository.go -destination=companyRepository_mock.go -package=repository
type ComapnyRepo interface {
	CreateComapny(company model.Company) (model.Company, error)
	GetCompanyByID(cID uint64) (model.Company, error)
	GetAllCompanies() ([]model.Company, error)
}

func NewCompanyRepo(db *gorm.DB) (ComapnyRepo, error) {
	if db == nil {
		log.Info().Msg("database cannot be nil")
		return nil, errors.New("database cannot be nil ")
	}
	return &Repo{
		db: db,
	}, nil
}

func (r *Repo) CreateComapny(company model.Company) (model.Company, error) {

	output := r.db.Create(&company)

	if output.Error != nil {
		log.Error().Err(output.Error).Msg("error in creating company table")
		return model.Company{}, errors.New("error in creating table")
	}

	return company, nil
}

func (r *Repo) GetCompanyByID(cID uint64) (model.Company, error) {

	var companydata model.Company
	output := r.db.Where("id = ?", cID).First(&companydata)
	if output.Error != nil {
		log.Error().Err(output.Error).Msg("error company id does not exists")
		return model.Company{}, errors.New("error company id does not exists")
	}
	return companydata, nil
}

func (r *Repo) GetAllCompanies() ([]model.Company, error) {

	var companiesData []model.Company

	output := r.db.Find(&companiesData)
	if output.Error != nil {
		log.Error().Err(output.Error).Msg("error while fetching companies data")
		return nil, errors.New("error while fetching companies data")
	}

	return companiesData, nil
}
