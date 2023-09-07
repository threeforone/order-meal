package service

import (
	"github.com/threeforone/order-meal/server/global"
	"github.com/threeforone/order-meal/server/model"
)

func GetDepartmentByCompanyId(companyId uint) (data []model.DepartmentResp, err error) {

	err = global.DB.Model(&model.Department{}).Select("id,name,company_id").Where("company_id = ?", companyId).Find(&data).Error
	return
}
