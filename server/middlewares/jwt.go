package middlewares

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/threeforone/order-meal/server/controller"
	"github.com/threeforone/order-meal/server/utils"
)

const (
	ContextUserIDKey = "userKey"
)

var (
	ErrorUserNotLogin = errors.New("当前用户未登录")
)

// JWTAuthMiddleware 基于JWT的认证中间件--客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controller.ResponseErrorWithCodeMsg(c, controller.CodeInvalidToken, "请求头缺少Auth Token")
			c.Abort()
			return
		}
		claims, err := utils.NewJWT().ParseToken(authHeader)

		if err != nil {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}

		status := claims.Status
		if status == 0 {
			controller.ResponseErrorWithCodeMsg(c, controller.CodeInvalidToken, "待审核，请等待管理员审核")
			c.Abort()
			return
		}
		if status == 2 {
			controller.ResponseErrorWithCodeMsg(c, controller.CodeInvalidToken, "账号暂不可用")
			c.Abort()
			return
		}
		c.Set("userInfo", claims.BaseClaims)
		c.Set("userId", claims.BaseClaims.ID)
		c.Set("admin", claims.BaseClaims.Admin)
		c.Set("superAdmin", claims.BaseClaims.SuperAdmin)
		c.Set("companyId", claims.BaseClaims.CompanyId)
		c.Set("departmentId", claims.BaseClaims.DepartmentId)
	}
}

func JWTAuthAdmin() func(c *gin.Context) {
	return func(c *gin.Context) {
		fmt.Println("---------------------")
		fmt.Println(c.Request.URL)
		superAdminCvt, _ := c.Get("superAdmin")
		superAdmin := superAdminCvt.(int8)
		if superAdmin == 0 {
			adminCvt, _ := c.Get("admin")
			admin := adminCvt.(int8)
			if admin == 0 {
				controller.ResponseErrorWithMsg(c, "权限不够")
				c.Abort()
				return
			}
		}

	}
}

func JWTAuthSuperAdmin() func(c *gin.Context) {
	return func(c *gin.Context) {
		superAdminCvt, _ := c.Get("superAdmin")
		superAdmin := superAdminCvt.(int8)
		if superAdmin == 0 {
			controller.ResponseErrorWithMsg(c, "请联系超级管理员")
			c.Abort()
			return
		}
	}
}
