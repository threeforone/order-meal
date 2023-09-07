package meal

import (
	"github.com/gin-gonic/gin"
	"github.com/threeforone/order-meal/server/controller"
	"github.com/threeforone/order-meal/server/middlewares"
)

func InitRouter(r *gin.RouterGroup) {
	mealApi := r.Group("/meal")
	mealApi.Use(middlewares.JWTAuthMiddleware()).Use(middlewares.JWTAuthAdmin())
	{
		mealApi.GET("/day", controller.GetDayStatistics)
		mealApi.GET("/statistics", controller.GetStatistics)
	}
}
