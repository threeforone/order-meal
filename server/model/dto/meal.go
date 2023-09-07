package dto

import "time"

type MealSearchReq struct {
	PageParam
	CreatedAt    time.Time `json:"createdAt"`
	CompanyID    uint      `json:"companyId"`
	DepartmentID uint      `json:"departmentId"`
	StartTime    time.Time `json:"startTime"`
	EndTime      time.Time `json:"endTime"`
}

type MealResp struct {
	ID             int64     `json:"id" db:"id"`
	UserId         int64     `json:"userId" db:"user_id"`
	Username       string    `json:"username" db:"username"`
	CompanyID      uint      `json:"companyId" db:"company_id"`
	CompanyName    string    `json:"companyName"`
	DepartmentID   uint      `json:"departmentId" db:"department_id"`
	DepartmentName string    `json:"departmentName"`
	CreatedAt      time.Time `json:"createdAt"`
}

type MealStatisticsResp struct {
	CompanyID      uint   `json:"companyId"`
	CompanyName    string `json:"companyName"`
	DepartmentID   uint   `json:"departmentId"`
	DepartmentName string `json:"departmentName"`
	StatisticsDate string `json:"statisticsDate"`
	Count          int    `json:"count"`
}
