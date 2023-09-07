package company

import (
	"github.com/gin-gonic/gin"
	"github.com/threeforone/order-meal/server/controller"
)

func InitRouter(r *gin.RouterGroup) {
	companyApi := r.Group("company")
	{
		companyApi.GET("", controller.GetCompanyList)
	}

	departmentApi := r.Group("department")
	{
		departmentApi.GET("", controller.GetDepartmentByCompanyId)
	}
}
