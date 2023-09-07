package user

import (
	"github.com/gin-gonic/gin"
	"github.com/threeforone/order-meal/server/controller"
	"github.com/threeforone/order-meal/server/middlewares"
)

func InitRouter(r *gin.RouterGroup) {
	r.POST("/login", controller.Login)
	r.POST("/signUp", controller.SignUp)
	userApi := r.Group("user")
	userApi.Use(middlewares.JWTAuthMiddleware())
	{
		api1 := userApi.Group("").Use(middlewares.JWTAuthAdmin())
		{
			api1.GET("/list", controller.GetUserListP)
			api1.GET("/detail/:id", controller.GetUserDetail)
			api1.PUT("/status", controller.ChangeUserStatus)
			api1.PUT("/reset", controller.ResetUserPassword)
		}
		api2 := userApi.Group("").Use(middlewares.JWTAuthAdmin()).Use(middlewares.JWTAuthSuperAdmin())
		{
			api2.PUT("/admin", controller.ChangeUserAdmin)
		}

	}

}
