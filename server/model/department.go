package model

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Name      string `json:"name"`
	CompanyId uint   `json:"companyId"`
}

func (Department) TableName() string {
	return "sys_department"
}

type DepartmentResp struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	CompanyId uint   `json:"companyId"`
}
