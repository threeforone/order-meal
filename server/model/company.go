package model

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name string `json:"name"`
}

func (Company) TableName() string {
	return "sys_company"
}

type CompanyResp struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
