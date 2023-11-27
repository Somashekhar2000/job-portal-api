package repository

import (
	"errors"
	"job-portal-api/internal/model"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

//go:generate mockgen -source=jobRepository.go -destination=jobRepository_mock.go -package=repository
type JobRepository interface {
	CreateJob(jodData model.Job) (model.Response, error)
	GetJobByCompanyID(cID uint) ([]model.Job, error)
	GetJobByJobID(cID uint) (model.Job, error)
	GetAllJobs() ([]model.Job, error)
}

func NewJobRepo(db *gorm.DB) (JobRepository, error) {
	if db == nil {
		log.Info().Msg("database cannot be nil")
		return nil, errors.New("data base cannot be nil")
	}

	return &Repo{
		db: db,
	}, nil
}

func (r *Repo) CreateJob(jobData model.Job) (model.Response, error) {

	output := r.db.Create(&jobData)

	if output.Error != nil {
		log.Error().Err(output.Error).Msg("error in creating job table")
		return model.Response{}, errors.New("could not create job")
	}

	return model.Response{
		Id: jobData.ID,
	}, nil
}

func (r *Repo) GetJobByCompanyID(cID uint) ([]model.Job, error) {

	var jobData []model.Job

	output := r.db.Preload("Company").Preload("Location").Preload("TechnologyStack").Preload("Qualifications").Preload("Shift").Preload("Jobtype").Where("cid = ?", cID).Find(&jobData)
	if output.Error != nil || output.RowsAffected == 0 {
		log.Error().Err(output.Error).Msg("error ivalid company id")
		return nil, errors.New("invalid company id")
	}

	return jobData, nil
}

func (r *Repo) GetJobByJobID(jID uint) (model.Job, error) {

	var jobData model.Job

	output := r.db.Preload("Company").Preload("Location").Preload("TechnologyStack").Preload("Qualifications").Preload("Shift").Preload("Jobtype").Where("id = ?", jID).First(&jobData)
	if output.Error != nil {
		log.Error().Err(output.Error).Msg("error in job id")
		return model.Job{}, errors.New("couls not find the job")
	}

	return jobData, nil
}

func (r *Repo) GetAllJobs() ([]model.Job, error) {

	var jobData []model.Job

	output := r.db.Preload("Company").Preload("Location").Preload("TechnologyStack").Preload("Qualifications").Preload("Shift").Preload("Jobtype").Find(&jobData)

	if output.Error != nil || output.RowsAffected == 0 {
		log.Error().Err(output.Error).Msg("error while retriving job data")
		return nil, errors.New("error while getting all jobs")
	}

	return jobData, nil
}
