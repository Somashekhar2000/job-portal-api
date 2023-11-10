package repository

import (
	"errors"
	"project/internal/model"

	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=jobDao.go -destination=companyrepository_mock.go -package=repository
type Company interface {
	CreateCompany(company model.Company) (model.Company, error)
	GetAllCompany() ([]model.Company, error)
	GetCompany(companyID int) (model.Company, error)
	PostJob(nj model.Job) (model.Response, error)
	GetJobs(companyID int) ([]model.Job, error)
	GetAllJobs() ([]model.Job, error)
	GetJobsByJobId(jobID int) (model.Job, error)
	FetchJobData(jid uint64) (model.Job, error)
}

// CreateCompany creates a new company record in the repository.
func (r *Repo) CreateCompany(companyData model.Company) (model.Company, error) {
	err := r.db.Create(&companyData).Error
	if err != nil {
		return model.Company{}, err
	}
	return companyData, nil
}

// GetAllCompany retrieves a list of companies from the repository.
func (r *Repo) GetAllCompany() ([]model.Company, error) {
	var companies []model.Company
	err := r.db.Find(&companies).Error
	if err != nil {
		return nil, err
	}

	return companies, nil
}

// GetCompanyByID retrieves a Company from the repository by its ID.
func (r *Repo) GetCompany(id int) (model.Company, error) {
	var companyModel model.Company
	id1 := uint64(id)
	err := r.db.Where("id = ?", id1).Find(&companyModel).Error
	if err != nil {
		return model.Company{}, err
	}
	return companyModel, nil

}

// CreateJob creates a new job in the repository.
func (r *Repo) PostJob(nj model.Job) (model.Response, error) {

	res := r.db.Create(&nj).Error
	if res != nil {
		log.Info().Err(res).Send()
		return model.Response{}, errors.New("job creation failed")
	}
	return model.Response{ID: uint64(nj.ID)}, nil
}

// GetJobsByID retrieves a list of jobs associated with a given ID from the repository.
func (r *Repo) GetJobs(id int) ([]model.Job, error) {
	var jobs []model.Job

	tx := r.db.Where("uid = ?", id)
	err := tx.Find(&jobs).Error
	if err != nil {
		return nil, err
	}
	return jobs, nil

}

// GetAllJobs retrieves all job records from the repository.
func (r *Repo) GetAllJobs() ([]model.Job, error) {
	var jobs []model.Job
	err := r.db.Preload("Locations").Preload("Skills").Preload("WorkModes").Preload("Qualifications").Preload("Shifts").Preload("JobTypes").Find(&jobs).Error
	if err != nil {
		return nil, err
	}

	return jobs, nil
}

// GetJobByID retrieves a job from the repository by its unique identifier.
func (r *Repo) GetJobsByJobId(id int) (model.Job, error) {
	var job model.Job

	tx := r.db.Preload("Comp").
		Preload("Locations").
		Preload("Skills").
		Preload("Qualifications").
		Preload("Shifts").Where("id = ?", id)
	err := tx.Find(&job).Error
	if err != nil {
		return model.Job{}, err
	}
	return job, nil

}

func (r *Repo) FetchJobData(jid uint64) (model.Job, error) {
	var j model.Job
	result := r.db.Preload("Comp").
		Preload("Locations").
		Preload("Skills").
		Preload("Qualifications").
		Preload("Shifts").
		Where("id = ?", jid).
		Find(&j)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return model.Job{}, result.Error
	}

	return j, nil
}
