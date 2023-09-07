package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/threeforone/order-meal/server/frontend"
	companyRouter "github.com/threeforone/order-meal/server/router/company"
	mealRouter "github.com/threeforone/order-meal/server/router/meal"
	orderRouter "github.com/threeforone/order-meal/server/router/order"
	userRouter "github.com/threeforone/order-meal/server/router/user"
)

func Routers(g *gin.Engine) {
	frontend.LoadStaticR(g)
	v := g.Group("/api/v1")
	companyRouter.InitRouter(v)
	userRouter.InitRouter(v)
	orderRouter.InitRouter(v)
	mealRouter.InitRouter(v)
}
