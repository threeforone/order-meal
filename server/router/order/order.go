package order

import (
	"github.com/gin-gonic/gin"
	"github.com/threeforone/order-meal/server/controller"
	"github.com/threeforone/order-meal/server/middlewares"
)

func InitRouter(r *gin.RouterGroup) {
	orderApi := r.Group("order")
	orderApi.Use(middlewares.JWTAuthMiddleware())
	{
		orderApi.GET("/time", controller.GetOrderTime)
		orderApi.GET("/status", controller.GetOrderStatus)
		orderApi.POST("", controller.Order)
	}

}
