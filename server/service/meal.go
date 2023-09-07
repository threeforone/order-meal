package service

import (
	"time"

	"github.com/threeforone/order-meal/server/global"
	"github.com/threeforone/order-meal/server/model/dto"
)

func MealSign() {

}

// type MealSearchReq struct {
// 	PageParam
// 	CreatedAt    time.Time `json:"createdAt"`
// 	CompanyID    uint      `json:"companyId"`
// 	DepartmentID uint      `json:"departmentId"`
// 	StartTime    time.Time `json:"startTime"`
// 	EndTime      time.Time `json:"endTime"`
// }

func GetMealDayStatistics(page, pageSize int, companyId, departmentId uint, createdAt *time.Time) (data dto.PageResp, err error) {
	sqlStr := global.DB.Table("meal AS m").Select(`m.id,
	m.user_id AS user_id,
	u.username as username,
	m.company_id AS company_id,
	c.NAME AS company_name,
	m.department_id AS department_id,
	d.NAME AS department_name,
	m.created_at as created_at`).
		Joins("LEFT JOIN sys_user AS u ON m.user_id = u.id").
		Joins("LEFT JOIN sys_company AS c ON m.company_id = c.id").
		Joins("LEFT JOIN sys_department AS d ON m.department_id = d.id").
		Where("DATE(m.created_at) = DATE(?)", createdAt)
	if companyId != 0 {
		sqlStr = sqlStr.Where("m.company_id = ? ", companyId)
	}
	if departmentId != 0 {
		sqlStr = sqlStr.Where("m.department_id = ? ", departmentId)
	}
	sqlStr = sqlStr.Count(&data.Total).Limit(pageSize).Offset(pageSize * (page - 1)).Order("m.created_at DESC")
	var dtos []dto.MealResp
	if err = sqlStr.Find(&dtos).Error; err != nil {
		return
	}
	data.Data = make([]any, len(dtos))

	for i, dto := range dtos {
		data.Data[i] = any(dto)
	}
	return
}

func GetMealStatistics(page, pageSize int, companyId, departmentId uint, startTime, endTime *time.Time) (data dto.PageResp, err error) {
	sqlStr := global.DB.Table("meal AS m").Select(`
	m.company_id AS company_id,
	c.NAME AS company_name,
	m.department_id AS department_id,
	d.NAME AS department_name,
	DATE(m.created_at) as statistics_date,
	count(1) as count `).
		Joins("LEFT JOIN sys_user AS u ON m.user_id = u.id").
		Joins("LEFT JOIN sys_company AS c ON m.company_id = c.id").
		Joins("LEFT JOIN sys_department AS d ON m.department_id = d.id")
	if companyId != 0 {
		sqlStr = sqlStr.Where("m.company_id = ? ", companyId)
	}
	if departmentId != 0 {
		sqlStr = sqlStr.Where("m.department_id = ? ", departmentId)
	}
	if startTime != nil {
		sqlStr = sqlStr.Where("DATE(m.created_at)>= DATE(?)", startTime)
	}
	if endTime != nil {
		sqlStr = sqlStr.Where("DATE(m.created_at)<= DATE(?)", endTime)
	}
	sqlStr = sqlStr.Group("m.company_id,m.department_id,DATE(m.created_at)").Count(&data.Total).Limit(pageSize).Offset(pageSize * (page - 1)).Order("DATE(m.created_at) DESC")
	var dtos []dto.MealStatisticsResp
	if err = sqlStr.Find(&dtos).Error; err != nil {
		return
	}
	data.Data = make([]any, len(dtos))

	for i, dto := range dtos {
		data.Data[i] = any(dto)
	}
	return
}
