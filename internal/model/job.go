package model

import "gorm.io/gorm"

// Company represents a company entity.
type Company struct {
	gorm.Model
	CompanyName string `json:"company_name"  gorm:"unique"`
	Adress      string `json:"company_adress" validate:"required"`
	Domain      string `json:"domain" validate:"required"`
}

// CreateCompany is used to create a new company.
type CreateCompany struct {
	CompanyName string `json:"company_name" validate:"required"`
	Adress      string `json:"company_adress" validate:"required"`
	Domain      string `json:"domain" validate:"required"`
}

// Job represents a job entity.
type Job struct {
	gorm.Model
	JobTitle  string  `json:"job_title" validate:"required"`
	JobSalary string  `json:"job_salary" validate:"required"`
	Company   Company `gorm:"ForeignKey:uid"`
	Uid       uint64  `JSON:"uid, omitempty"`
}

// CreateJob is used to create a new job.
type CreateJob struct {
	JobTitle  string `json:"job_title" validate:"required"`
	JobSalary string `json:"job_salary" validate:"required"`
}
