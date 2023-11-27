package model

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	Company         Company           `json:"company" gorm:"ForeignKey:cid"`
	Cid             uint              `json:"cid"`
	Jobname         string            `json:"jobname" validate:"required"`
	MinNoticePeriod int               `json:"min_notice_period" validate:"required"`
	MaxNoticePeriod uint              `json:"max_notice_period" validate:"required"`
	Location        []Location        `json:"location" gorm:"many2many:job_location;"`
	TechnologyStack []TechnologyStack `json:"skills" gorm:"many2many:job_techstack;"`
	Description     string            `json:"description" validate:"required"`
	MinExperience   int               `json:"min_experience" validate:"required"`
	MaxExperience   uint              `json:"max_experience" validate:"required"`
	Qualifications  []Qualification   `json:"qualifications" gorm:"many2many:job_qualification;"`
	Shift           []Shift           `json:"shifts" gorm:"many2many:job_shift;" `
	Jobtype         []JobType         `json:"jobtype" gorm:"many2many:job_type;"`
}

type JobType struct {
	gorm.Model
	JobTypeName string `json:"jobtype"`
}

type Location struct {
	gorm.Model
	PlaceName string `json:"place_name"`
}

type TechnologyStack struct {
	gorm.Model
	StackName string `json:"stack_name"`
}

type Qualification struct {
	gorm.Model
	QualificationRequired string `json:"qualification_required"`
}

type Shift struct {
	gorm.Model
	ShiftType string `json:"shift_type"`
}

type NewJobs struct {
	Jobname         string `json:"jobName" validate:"required"`
	MinNoticePeriod int    `json:"minNoticePeriod" validate:"required"`
	MaxNoticePeriod uint   `json:"maxNoticePeriod" validate:"required"`
	Location        []uint `json:"location" `
	TechnologyStack []uint `json:"technologyStack" `
	Description     string `json:"description" validate:"required"`
	MinExperience   int    `json:"minExperience" validate:"required"`
	MaxExperience   uint   `json:"maxExperience" validate:"required"`
	Qualifications  []uint `json:"qualifications"`
	Shift           []uint `json:"shifts"`
	Jobtype         []uint `json:"jobtype"`
}

type Response struct {
	Id uint `json:"id"`
}

type NewUserApplication struct {
	Name string       `json:"name" validate:"required"`
	Age  string       `json:"age" validate:"required"`
	Jid  uint         `json:"jid" validate:"required"`
	Jobs Requestfield `json:"job_application"`
}

type Requestfield struct {
	NoticePeriod    int    `json:"noticePeriod" validate:"required"`
	Location        []uint `json:"location"`
	TechnologyStack []uint `json:"technologyStack"`
	Experience      int    `json:"experience" validate:"required"`
	Qualifications  []uint `json:"qualifications"`
	Shift           []uint `json:"shifts"`
	Jobtype         []uint `json:"jobtype"`
}
