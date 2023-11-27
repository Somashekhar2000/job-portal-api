package service

import (
	"context"
	"encoding/json"
	"job-portal-api/internal/cache"
	"job-portal-api/internal/model"
	"job-portal-api/internal/repository"
	"sync"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

//go:generate mockgen -source=jobService.go -destination=jobService_mock.go -package=service
type JobService interface {
	CreateJobByCompanyId(jobdata model.NewJobs, cID uint) (model.Response, error)
	ViewJobByCompanyID(cID uint) ([]model.Job, error)
	ViewJobByJobID(jID uint) (model.Job, error)
	ViewAllJobs() ([]model.Job, error)
	ProcessApplication(applications []model.NewUserApplication) []model.NewUserApplication
}

func NewJobService(jobService repository.JobRepository, rdb cache.Caching) (JobService, error) {
	if jobService == nil {
		log.Info().Msg("jobservice cannot be nil")
	}
	return &Service{
		jobRepo: jobService,
		rdb:     rdb,
	}, nil
}

func (s *Service) CreateJobByCompanyId(jobDetails model.NewJobs, cID uint) (model.Response, error) {

	jobData := model.Job{
		Cid:             cID,
		Jobname:         jobDetails.Jobname,
		MinNoticePeriod: jobDetails.MinNoticePeriod,
		MaxNoticePeriod: jobDetails.MaxNoticePeriod,
		Description:     jobDetails.Description,
		MinExperience:   jobDetails.MinExperience,
		MaxExperience:   jobDetails.MaxExperience,
	}

	for _, v := range jobDetails.Jobtype {
		jobtype := model.JobType{
			Model: gorm.Model{
				ID: v,
			},
		}
		jobData.Jobtype = append(jobData.Jobtype, jobtype)
	}

	for _, v := range jobDetails.Location {
		jobLocation := model.Location{
			Model: gorm.Model{
				ID: v,
			},
		}
		jobData.Location = append(jobData.Location, jobLocation)
	}

	for _, v := range jobDetails.TechnologyStack {
		jobSkill := model.TechnologyStack{
			Model: gorm.Model{
				ID: v,
			},
		}
		jobData.TechnologyStack = append(jobData.TechnologyStack, jobSkill)
	}

	for _, v := range jobDetails.Qualifications {
		jobQualification := model.Qualification{
			Model: gorm.Model{
				ID: v,
			},
		}
		jobData.Qualifications = append(jobData.Qualifications, jobQualification)
	}

	for _, v := range jobDetails.Shift {
		jobShift := model.Shift{
			Model: gorm.Model{
				ID: v,
			},
		}
		jobData.Shift = append(jobData.Shift, jobShift)
	}

	responseData, err := s.jobRepo.CreateJob(jobData)
	if err != nil {
		return model.Response{}, err
	}

	return responseData, nil

}

func (s *Service) ViewJobByCompanyID(cID uint) ([]model.Job, error) {

	jobData, err := s.jobRepo.GetJobByCompanyID(cID)

	if err != nil {
		return nil, err
	}

	return jobData, nil
}

func (s *Service) ViewJobByJobID(jID uint) (model.Job, error) {

	jobData, err := s.jobRepo.GetJobByJobID(jID)
	if err != nil {
		return model.Job{}, err
	}
	return jobData, nil
}

func (s *Service) ViewAllJobs() ([]model.Job, error) {
	jobData, err := s.jobRepo.GetAllJobs()
	if err != nil {
		return nil, err
	}

	return jobData, nil
}

func (s *Service) ProcessApplication(applications []model.NewUserApplication) []model.NewUserApplication {
	ctx := context.Background()
	wg := new(sync.WaitGroup)
	ch := make(chan model.NewUserApplication)
	var finalData []model.NewUserApplication

	for _, v := range applications {
		wg.Add(1)
		go func(application model.NewUserApplication) {
			defer wg.Done()

			var jobData model.Job

			val, err := s.rdb.GetTheCacheData(ctx, application.Jid)

			if err != nil {
				jobDataFromDB, err := s.jobRepo.GetJobByJobID(application.Jid)
				if err != nil {
					log.Error().Err(err).Msg("invalid application job id does not exists")
					return
				}
				err = s.rdb.AddToTheCache(ctx, application.Jid, jobDataFromDB)
				if err != nil {
					return
				}
				jobData = jobDataFromDB

			} else {
				err = json.Unmarshal([]byte(val), &jobData)
				if err != nil {
					log.Error().Err(err).Msg("error in un marshaling")
					return
				}
			}
			check := CompareData(application, jobData)

			if check {
				ch <- application
			}

		}(v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		finalData = append(finalData, v)
	}

	return finalData
}

func CompareData(application model.NewUserApplication, jobData model.Job) bool {
	totalFields := 0
	matchedFields := 0

	totalFields++
	if application.Jobs.NoticePeriod >= jobData.MinNoticePeriod && application.Jobs.NoticePeriod <= int(jobData.MaxNoticePeriod) {
		matchedFields++
	}

	totalFields++
	if application.Jobs.Experience >= jobData.MinExperience && application.Jobs.Experience <= int(jobData.MaxExperience) {
		matchedFields++
	}

	count := 0
	totalFields++
	for _, v := range application.Jobs.Location {
		for _, v1 := range jobData.Location {
			if v == v1.ID {
				count++
			}
		}
	}
	if count != 0 {
		matchedFields++
	}

	count = 0
	totalFields++
	for _, v := range application.Jobs.TechnologyStack {
		for _, v1 := range jobData.TechnologyStack {
			if v == v1.ID {
				count++
			}
		}
	}
	if count != 0 {
		matchedFields++
	}

	count = 0
	totalFields++
	for _, v := range application.Jobs.Qualifications {
		for _, v1 := range jobData.Qualifications {
			if v == v1.ID {
				count++
			}
		}
	}
	if count != 0 {
		matchedFields++
	}

	count = 0
	totalFields++
	for _, v := range application.Jobs.Shift {
		for _, v1 := range jobData.Shift {
			if v == v1.ID {
				count++
			}
		}
	}
	if count != 0 {
		matchedFields++
	}

	count = 0
	totalFields++
	for _, v := range application.Jobs.Jobtype {
		for _, v1 := range jobData.Jobtype {
			if v == v1.ID {
				count++
			}
		}
	}
	if count != 0 {
		matchedFields++
	}

	if matchedFields*2 >= totalFields {
		return true
	}

	return false
}
