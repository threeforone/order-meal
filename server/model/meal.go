package model

import "gorm.io/gorm"

type Meal struct {
	gorm.Model
	UserId       uint
	CompanyID    uint `gorm:"default:null"`
	DepartmentID uint `gorm:"default:null"`
	DeletedBy    uint `gorm:"default:null"`
}
