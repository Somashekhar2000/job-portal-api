package model

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	CompanyName string `json:"companyName" validate:"required"`
	Address     string `json:"address"`
	Domain      string `json:"domain"`
}

type AddCompany struct {
	CompanyName string `json:"companyName" validate:"required"`
	Address     string `json:"address"`
	Domain      string `json:"domain"`
}
