package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/threeforone/order-meal/server/global"
	"github.com/threeforone/order-meal/server/model"
	"github.com/threeforone/order-meal/server/model/dto"
	"github.com/threeforone/order-meal/server/service"
	"github.com/threeforone/order-meal/server/utils"
)

func Login(c *gin.Context) {
	var userReq = dto.UserLoginReq{}
	c.ShouldBind(&userReq)
	user, err := service.Login(userReq)
	if err != nil {
		ResponseError(c, CodeInvalidPassword)
		return
	}
	TokenNext(c, *user)
}

func SignUp(c *gin.Context) {
	var userReq = dto.UserSignUpReq{}
	c.ShouldBind(&userReq)
	err := service.SignUp(userReq)
	if err != nil {
		ResponseErrorWithMsg(c, err.Error())
		return
	}
	ResponseSuccess(c, nil)
}

// TokenNext 登录以后签发jwt
func TokenNext(c *gin.Context, user model.User) {
	status := user.Status
	if status == 0 {
		ResponseErrorWithCodeMsg(c, CodeInvalidToken, "待审核，请等待管理员审核")
		return
	}
	if status == 2 {
		ResponseErrorWithCodeMsg(c, CodeInvalidToken, "账号暂不可用")
		return
	}

	j := &utils.JWT{SigningKey: []byte(global.Conf.JWT.SigningKey)} // 唯一签名

	claims := j.CreateClaims(dto.BaseClaims{
		ID:           user.ID,
		Username:     user.Username,
		Admin:        user.Admin,
		SuperAdmin:   user.SuperAdmin,
		CompanyId:    user.CompanyID,
		DepartmentId: user.DepartmentID,
		Status:       user.Status,
	})
	token, err := j.CreateToken(claims)
	data := map[string]interface{}{
		"user":  claims,
		"token": token,
	}
	if err != nil {
		ResponseError(c, CodeInvalidToken)
		return
	}
	ResponseSuccess(c, data)
}

func GetUserListP(c *gin.Context) {
	req := queryToSearchReq(c)
	data, err := service.GetUserListP(*req)
	if err != nil {
		ResponseErrorWithMsg(c, err.Error())
		return
	}
	ResponseSuccess(c, data)
}

// userApi.GET("/detail/:id") //单个用户详情，根据id
func GetUserDetail(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		ResponseErrorWithMsg(c, err.Error())
		return
	}
	data, err := service.GetUserDetail(uint(userId))
	if err != nil {
		ResponseErrorWithMsg(c, err.Error())
		return
	}
	ResponseSuccess(c, data)
}

func ChangeUserStatus(c *gin.Context) {
	var req dto.UserStatusReq
	c.ShouldBind(&req)
	err := service.ChangeUserStatus(req)
	if err != nil {
		ResponseErrorWithMsg(c, err.Error())
		return
	}
	ResponseSuccess(c, nil)
}

// userApi.PUT("/admin")      //授予管理员
func ChangeUserAdmin(c *gin.Context) {

	var req dto.UserAdminReq
	c.ShouldBind(&req)
	err := service.ChangeUserAdmin(req)
	if err != nil {
		ResponseErrorWithMsg(c, err.Error())
		return
	}
	ResponseSuccess(c, nil)
}

// userApi.PUT("/reset")      //重置密码
func ResetUserPassword(c *gin.Context) {
	var req dto.UserResetReq
	c.ShouldBind(&req)
	err := service.ChangeUserPassword(req)
	if err != nil {
		ResponseErrorWithMsg(c, err.Error())
		return
	}
	ResponseSuccess(c, nil)
}

func queryToSearchReq(c *gin.Context) *dto.UserSearchReq {
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
	var status *int8
	if str, flag := c.GetQuery("status"); flag && str != "" {
		cvt, _ := strconv.Atoi(str)
		a := int8(cvt)
		status = &a
	}
	var admin *int8
	if str, flag := c.GetQuery("admin"); flag && str != "" {
		cvt, _ := strconv.Atoi(str)
		a := int8(cvt)
		admin = &a
	}
	var username string
	if str, flag := c.GetQuery("username"); flag && str != "" {
		username = str
	}
	var companyId *uint
	if str, flag := c.GetQuery("companyId"); flag && str != "" {
		cvt, _ := strconv.Atoi(str)
		a := uint(cvt)
		companyId = &a
	}
	var departmentId *uint
	if str, flag := c.GetQuery("departmentId"); flag && str != "" {
		cvt, _ := strconv.Atoi(str)
		a := uint(cvt)
		companyId = &a
	}
	req := &dto.UserSearchReq{
		PageParam: dto.PageParam{
			Page:     page,
			PageSize: pageSize,
		},
		Status:       status,
		Admin:        admin,
		Username:     username,
		CompanyId:    companyId,
		DepartmentId: departmentId,
	}
	return req
}
