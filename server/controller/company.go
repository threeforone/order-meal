package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/threeforone/order-meal/server/service"
)

func GetCompanyList(c *gin.Context) {
	data, err := service.GetCompanyList()
	if err != nil {
		ResponseErrorWithData(c, CodeFailure, err.Error(), data)
		return
	}
	ResponseSuccess(c, data)

}
