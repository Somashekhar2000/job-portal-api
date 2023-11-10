package services

import (
	"errors"
	"project/internal/model"
	"sync"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// CompanyService is an interface for managing companies and jobs.
//

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
func (s *Service) JobCreate(newJob model.NewJobRequest, id uint64) (model.Response, error) {

	app := model.Job{
		CompanyId:           id,
		JobTitle:            newJob.JobTitle,
		Salary:              newJob.Salary,
		MinimumNoticePeriod: newJob.MinimumNoticePeriod,
		MaximumNoticePeriod: newJob.MaximumNoticePeriod,
		Budget:              newJob.Budget,
		JobDescription:      newJob.JobDescription,
		MinExperience:       newJob.MinExperience,
	}
	for _, v := range newJob.QualificationIDs {
		tempData := model.Qualification{
			Model: gorm.Model{
				ID: v,
			},
		}
		app.Qualifications = append(app.Qualifications, tempData)
	}
	for _, v := range newJob.LocationIDs {
		tempData := model.Location{
			Model: gorm.Model{
				ID: v,
			},
		}
		app.Locations = append(app.Locations, tempData)
	}
	for _, v := range newJob.SkillIDs {
		tempData := model.Skill{
			Model: gorm.Model{
				ID: v,
			},
		}
		app.Skills = append(app.Skills, tempData)
	}
	for _, v := range newJob.WorkModeIDs {
		tempData := model.WorkMode{
			Model: gorm.Model{
				ID: v,
			},
		}
		app.WorkModes = append(app.WorkModes, tempData)
	}
	for _, v := range newJob.ShiftIDs {
		tempData := model.Shift{
			Model: gorm.Model{
				ID: v,
			},
		}
		app.Shifts = append(app.Shifts, tempData)
	}
	for _, v := range newJob.JobTypeIDs {
		tempData := model.JobType{
			Model: gorm.Model{
				ID: v,
			},
		}
		app.JobTypes = append(app.JobTypes, tempData)
	}
	jobData, err := s.companyRepo.PostJob(app)
	if err != nil {
		return model.Response{}, err
	}
	return jobData, nil
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

func (s *Service) ProcessJobApplications(appData []model.NewUserApplication) ([]model.NewUserApplication, error) {
	var wg = new(sync.WaitGroup)
	ch := make(chan model.NewUserApplication)
	var finalData []model.NewUserApplication

	for _, v := range appData {
		wg.Add(1)
		go func(v model.NewUserApplication) {
			defer wg.Done()

			val, err := s.companyRepo.GetJobsByJobId(int(v.ID))
			if err != nil {
				return
			}
			check, value, err := s.Compare(v, val)
			if err != nil {
				return
			}
			if check {
				ch <- value
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

	return finalData, nil
}

func (s *Service) Compare(appData model.NewUserApplication, jobData model.Job) (bool, model.NewUserApplication, error) {
	matchingConditions := 0
	totalConditions := 8
	if appData.Jobs.Experience >= jobData.MinExperience {
		matchingConditions++
	}

	if appData.Jobs.NoticePeriod >= jobData.MinimumNoticePeriod {
		matchingConditions++
	}

	for _, v := range appData.Jobs.WorkModeIDs {
		for _, v1 := range jobData.WorkModes {
			if v == v1.ID {
				matchingConditions++
				break
			}
		}
	}

	for _, v := range appData.Jobs.JobTypeIDs {
		for _, v1 := range jobData.JobTypes {
			if v == v1.ID {
				matchingConditions++
				break
			}
		}
	}

	for _, v := range appData.Jobs.Location {
		for _, v1 := range jobData.Locations {
			if v == v1.ID {
				matchingConditions++
				break
			}
		}
	}

	for _, v := range appData.Jobs.Qualifications {
		for _, v1 := range jobData.Qualifications {
			if v == v1.ID {
				matchingConditions++
				break
			}
		}
	}

	for _, v := range appData.Jobs.Skills {
		for _, v1 := range jobData.Skills {
			if v == v1.ID {
				matchingConditions++
				break
			}
		}
	}

	for _, v := range appData.Jobs.Shift {
		for _, v1 := range jobData.Shifts {
			if v == v1.ID {
				matchingConditions++
				break
			}
		}
	}

	if matchingConditions*2 >= totalConditions {
		return true, appData, nil
	}

	return false, model.NewUserApplication{}, nil
}
