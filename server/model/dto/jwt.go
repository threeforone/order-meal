package dto

import "github.com/golang-jwt/jwt/v4"

type BaseClaims struct {
	ID           uint   `json:"id"`
	Username     string `json:"username"`
	Admin        int8   `json:"admin"`
	SuperAdmin   int8   `json:"superAdmin"`
	DepartmentId uint   `json:"departmentId"`
	CompanyId    uint   `json:"companyId"`
	Status       int8   `json:"status"`
}

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}
