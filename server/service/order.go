package service

import (
	"errors"
	"time"

	"github.com/threeforone/order-meal/server/global"
	"github.com/threeforone/order-meal/server/initialize"
	"github.com/threeforone/order-meal/server/model"
	"github.com/threeforone/order-meal/server/model/dto"
)

func ChangeOrderTime() {
	req := new(dto.OrderTimeReq)
	initialize.ChangeOrderTime(req.StartTime.Format("15:04"), req.EndTime.Format("15:04"))
}

func CheckInTime() bool {
	currentTime := time.Now()
	currentDate := currentTime.Format("2006-01-02")
	startTimeString := currentDate + " " + global.OrderT.StartTime + ":00"
	endTimeString := currentDate + " " + global.OrderT.EndTime + ":59"
	startTime, _ := time.ParseInLocation("2006-01-02 15:04:05", startTimeString, time.Local)
	endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", endTimeString, time.Local)
	return currentTime.After(startTime) && currentTime.Before(endTime)
}

func GetOrderTime() global.OrderTime {
	return *global.OrderT
}

func GetOrderStatus(userId uint) int {
	var meal *[]model.Meal
	global.DB.Model(&model.Meal{}).Where("user_id = ?", userId).Where("DATE(created_at) = DATE(?)", time.Now()).First(&meal)
	return len(*meal)
}
func OrderMeal(userId, companyId, departmentId uint) error {
	flag := CheckInTime()
	if !flag {
		return errors.New("不在点餐时间内")
	}
	num := GetOrderStatus(userId)
	if num != 0 {
		return errors.New("已经点餐了")
	}
	meal := model.Meal{
		UserId:       userId,
		CompanyID:    companyId,
		DepartmentID: departmentId,
	}
	return global.DB.Model(&model.Meal{}).Create(&meal).Error

}
