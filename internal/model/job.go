package model

import "gorm.io/gorm"

// Company represents a company entity.
type Company struct {
	gorm.Model
	CompanyName string `json:"companyName"  gorm:"unique"`
	Adress      string `json:"companyAdress" validate:"required"`
	Domain      string `json:"domain" validate:"required"`
}

// CreateCompany is used to create a new company.
type CreateCompany struct {
	CompanyName string `json:"companyName" validate:"required"`
	Adress      string `json:"companyAdress" validate:"required"`
	Domain      string `json:"domain" validate:"required"`
}

// Job represents a job entity.
type Job struct {
	gorm.Model
	JobTitle            string          `json:"job_title" validate:"required"`
	Salary              string          `json:"sal" validate:"required"`
	CompanyId           uint64          `json:"cid" validate:"required"`
	Comp                Company         `gorm:"ForeignKey:CompanyId"`
	MinimumNoticePeriod uint64          `json:"min_np" validate:"required"`
	MaximumNoticePeriod uint64          `json:"max_np" validate:"required"`
	Budget              float64         `json:"budget" validate:"required"`
	JobDescription      string          `json:"job_desc" validate:"required"`
	MinExperience       float64         `json:"min_exp" validate:"required"`
	MaxExperience       float64         `json:"max_exp" validate:"required"`
	Locations           []Location      `gorm:"many2many:job_locations;"`
	Skills              []Skill         `gorm:"many2many:job_skills;"`
	WorkModes           []WorkMode      `gorm:"many2many:job_work_modes;"`
	Qualifications      []Qualification `gorm:"many2many:job_qualifications;"`
	Shifts              []Shift         `gorm:"many2many:job_shifts;"`
	JobTypes            []JobType       `gorm:"many2many:job_jobtypes;"`
}
type Location struct {
	gorm.Model
	State string `json:"state" validate:"required"`
}
type Skill struct {
	gorm.Model
	Skillsets string `json:"skillsets" validate:"required"`
}
type WorkMode struct {
	gorm.Model
	Mode string `json:"work_mode" validate:"required"`
}
type Qualification struct {
	gorm.Model
	Degree string `json:"degree" validate:"required"`
}
type Shift struct {
	gorm.Model
	ShiftType string `json:"shift_type" validate:"required"`
}
type JobType struct {
	gorm.Model
	Typeofjob string `json:"job_type" validate:"required"`
}

type NewJobRequest struct {
	JobTitle            string  `json:"jobTitle" validate:"required"`
	Salary              string  `json:"sal" validate:"required"`
	MinimumNoticePeriod uint64  `json:"minNp" validate:"required"`
	MaximumNoticePeriod uint64  `json:"maxNp" validate:"required"`
	Budget              float64 `json:"budget" validate:"required"`
	JobDescription      string  `json:"jobDesc" validate:"required"`
	MinExperience       float64 `json:"minExp" validate:"required"`
	MaxExperience       float64 `json:"maxExp" validate:"required"`
	LocationIDs         []uint
	SkillIDs            []uint
	WorkModeIDs         []uint
	QualificationIDs    []uint
	ShiftIDs            []uint
	JobTypeIDs          []uint
}
type Response struct {
	ID uint64
}

type RequestFromUser struct {
	NoticePeriod   uint64  `json:"noticePeriod" validate:"required"`
	Location       []uint  `json:"location" `
	Skills         []uint  `json:"technologyStack" `
	Experience     float64 `json:"experience" validate:"required"`
	Qualifications []uint  `json:"qualifications"`
	Shift          []uint  `json:"shifts"`
	WorkModeIDs    []uint  `json:"workmode"`
	JobTypeIDs     []uint  `json:"jobtype"`
}
type NewUserApplication struct {
	Name string          `json:"name"`
	Age  string          `json:"age"`
	ID   uint64          `json:"jid"`
	Jobs RequestFromUser `json:"job_application"`
}
