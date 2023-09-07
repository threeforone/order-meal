package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/threeforone/order-meal/server/service"
)

func GetDepartmentByCompanyId(c *gin.Context) {
	companyId, _ := strconv.ParseUint(c.Query("companyId"), 10, 64)
	data, err := service.GetDepartmentByCompanyId(uint(companyId))
	if err != nil {
		ResponseErrorWithData(c, CodeFailure, err.Error(), data)
		return
	}
	ResponseSuccess(c, &data)

}
