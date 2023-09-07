package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UUID         string     `json:"uuid" gorm:"size:191"`
	Username     string     `json:"username" gorm:"size:255"`
	Password     string     `json:"-" gorm:"size:255"`
	CompanyID    uint       `json:"companyId" gorm:"default:null"`
	DepartmentID uint       `json:"departmentId" gorm:"default:null"`
	Status       int8       `json:"status" gorm:"default:0;comment:'0 待审核，1已审核，2 离职'"`
	Admin        int8       `json:"admin" gorm:"default:0;comment:'是否是管理员'"`
	SuperAdmin   int8       `json:"superAdmin" gorm:"default:0;comment:'是否是超级管理员'"`
	Company      Company    `json:"company" gorm:"foreignKey:CompanyID"`
	Department   Department `json:"department" gorm:"foreignKey:DepartmentID"`
}

func (User) TableName() string {
	return "sys_user"
}
