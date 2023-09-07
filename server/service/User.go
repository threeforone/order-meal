package service

import (
	"errors"
	"fmt"

	"github.com/threeforone/order-meal/server/global"
	"github.com/threeforone/order-meal/server/model"
	"github.com/threeforone/order-meal/server/model/dto"
	"github.com/threeforone/order-meal/server/utils"
	"gorm.io/gorm"
)

func SignUp(userReq dto.UserSignUpReq) error {
	user := model.User{
		Username:     userReq.Username,
		Password:     utils.BcryptHash(userReq.Password),
		CompanyID:    userReq.CompanyId,
		DepartmentID: userReq.DepartmentId,
	}
	existUser := make([]model.User, 0)
	err := global.DB.Model(&model.User{}).Where("username = ?", user.Username).First(&existUser).Error
	fmt.Print(err)
	if len(existUser) != 0 {
		return errors.New("用户名已存在请更换")
	}
	return global.DB.Model(&user).Create(&user).Error
}

func Login(userReq dto.UserLoginReq) (*model.User, error) {
	user := model.User{}
	err := global.DB.Model(&user).Where("username = ?", userReq.Username).First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(userReq.Password, user.Password); !ok {
			return &user, errors.New("密码错误")
		}
	}
	return &user, err
}

func changeUserStatus(id uint, status int) error {
	user := &model.User{Model: gorm.Model{ID: id}}
	return global.DB.Model(user).UpdateColumn("status", status).Error
}

func GetUserListP(req dto.UserSearchReq) (data dto.PageResp, err error) {
	query := global.DB.Model(&model.User{})
	if req.Status != nil {
		query = query.Where("status = ?", *req.Status)
	}
	if req.Admin != nil {
		query.Where("admin = ? ", *req.Admin)
	}
	if req.CompanyId != nil {
		query.Where("company_id=?", *req.CompanyId)
	}
	if req.DepartmentId != nil {
		query.Where("company_id=?", *req.DepartmentId)
	}
	if req.Username != "" {
		query.Where("username like concat('%',?,'%')", req.Username)
	}
	if err = query.Count(&data.Total).Error; err != nil {
		return
	}
	if req.Page > 0 && req.PageSize > 0 {
		query = query.Limit(int(req.PageSize)).Offset((int(req.PageSize * (req.Page - 1))))
	}
	var dtos []model.User
	if err = query.Order("created_at DESC").Preload("Company").Preload("Department").Find(&dtos).Error; err != nil {
		return
	}
	data.Data = make([]any, len(dtos))

	for i, dto := range dtos {
		data.Data[i] = any(dto)
	}
	return
}

func GetUserDetail(id uint) (user model.User, err error) {
	err = global.DB.Model(&user).Where("id = ?", id).First(&user).Error
	return
}

func ChangeUserStatus(req dto.UserStatusReq) error {
	return changeUserStatus(req.Id, int(req.Status))
}

func ChangeUserAdmin(req dto.UserAdminReq) error {

	user := &model.User{Model: gorm.Model{ID: req.Id}}
	return global.DB.Model(user).UpdateColumn("admin", req.Admin).Error
}

func ChangeUserPassword(req dto.UserResetReq) error {
	user := &model.User{Model: gorm.Model{ID: req.Id}}
	return global.DB.Model(user).UpdateColumn("password", utils.BcryptHash(req.Password)).Error
}
