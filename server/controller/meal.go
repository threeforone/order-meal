package controller

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/threeforone/order-meal/server/service"
)

func GetDayStatistics(c *gin.Context) {
	var page int = 1
	if str, flag := c.GetQuery("page"); flag && str != "" {
		cvt, _ := strconv.Atoi(str)
		page = cvt
	}
	var pageSize int = 10
	if str, flag := c.GetQuery("pageSize"); flag && str != "" {
		cvt, _ := strconv.Atoi(str)
		pageSize = cvt
	}

	var companyId uint
	if str, flag := c.GetQuery("companyId"); flag && str != "" {
		cvt, _ := strconv.Atoi(str)
		a := uint(cvt)
		companyId = a
	}
	var departmentId uint
	if str, flag := c.GetQuery("departmentId"); flag && str != "" {
		cvt, _ := strconv.Atoi(str)
		a := uint(cvt)
		companyId = a
	}
	createdAt := time.Now()
	data, err := service.GetMealDayStatistics(page, pageSize, companyId, departmentId, &createdAt)
	if err != nil {
		ResponseErrorWithMsg(c, err.Error())
		return
	}
	ResponseSuccess(c, data)
}

func GetStatistics(c *gin.Context) {
	var page int = 1
	if str, flag := c.GetQuery("page"); flag && str != "" {
		cvt, _ := strconv.Atoi(str)
		page = cvt
	}
	var pageSize int = 10
	if str, flag := c.GetQuery("pageSize"); flag && str != "" {
		cvt, _ := strconv.Atoi(str)
		pageSize = cvt
	}

	var companyId uint
	if str, flag := c.GetQuery("companyId"); flag && str != "" {
		cvt, _ := strconv.Atoi(str)
		a := uint(cvt)
		companyId = a
	}
	var departmentId uint
	if str, flag := c.GetQuery("departmentId"); flag && str != "" {
		cvt, _ := strconv.Atoi(str)
		a := uint(cvt)
		companyId = a
	}
	var startTime *time.Time
	if str, flag := c.GetQuery("startTime"); flag && str != "" {
		cvt, _ := time.Parse("2006-01-02", str)
		startTime = &cvt
	}
	var endTime *time.Time
	if str, flag := c.GetQuery("endTime"); flag && str != "" {
		cvt, _ := time.Parse("2006-01-02", str)
		endTime = &cvt
	}
	data, err := service.GetMealStatistics(page, pageSize, companyId, departmentId, startTime, endTime)
	if err != nil {
		ResponseErrorWithMsg(c, err.Error())
		return
	}
	ResponseSuccess(c, data)
}
