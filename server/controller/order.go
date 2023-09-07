package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/threeforone/order-meal/server/service"
)

func GetOrderTime(c *gin.Context) {
	data := service.GetOrderTime()
	ResponseSuccess(c, &data)
}

func GetOrderStatus(c *gin.Context) {
	data := service.GetOrderStatus(c.GetUint("userId"))
	ResponseSuccess(c, &data)
}

func Order(c *gin.Context) {
	err := service.OrderMeal(c.GetUint("userId"), c.GetUint("companyId"), c.GetUint("departmentId"))
	if err != nil {
		ResponseErrorWithMsg(c, err.Error())
		return
	}
	ResponseSuccess(c, nil)
}
