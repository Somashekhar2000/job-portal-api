package repository

import (
	"project/internal/model"
)

//go:generate mockgen -source=jobDao.go -destination=companyrepository_mock.go -package=repository
type Company interface {
	CreateCompany(company model.Company) (model.Company, error)
	GetAllCompany() ([]model.Company, error)
	GetCompany(companyID int) (model.Company, error)
	CreateJob(job model.Job) (model.Job, error)
	GetJobs(companyID int) ([]model.Job, error)
	GetAllJobs() ([]model.Job, error)
	GetJobsByJobId(jobID int) (model.Job, error)
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
	query := r.db.Where("id = ?", id1)
	err := query.Find(&companyModel).Error
	if err != nil {
		return model.Company{}, err
	}
	return companyModel, nil

}

// CreateJob creates a new job in the repository.
func (r *Repo) CreateJob(job model.Job) (model.Job, error) {
	err := r.db.Create(&job).Error
	if err != nil {
		return model.Job{}, err
	}
	return job, nil
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
	err := r.db.Find(&jobs).Error
	if err != nil {
		return nil, err
	}

	return jobs, nil
}

// GetJobByID retrieves a job from the repository by its unique identifier.
func (r *Repo) GetJobsByJobId(id int) (model.Job, error) {
	var job model.Job

	tx := r.db.Where("id = ?", id)
	err := tx.Find(&job).Error
	if err != nil {
		return model.Job{}, err
	}
	return job, nil

}
