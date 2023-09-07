package service

import (
	"github.com/threeforone/order-meal/server/global"
	"github.com/threeforone/order-meal/server/model"
)

func GetCompanyList() (data []model.CompanyResp, err error) {
	err = global.DB.Model(&model.Company{}).Select("id,name").Order("id ASC").Find(&data).Error
	return
}
