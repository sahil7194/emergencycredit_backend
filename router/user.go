package router

import (
	"backend/controllers"
	"backend/middlewares"

	"github.com/gin-gonic/gin"
)

func registerUserRoutes(rg *gin.RouterGroup) {

	user := rg.Group("/user")

	user.Use(middlewares.AuthMiddleWare())
	{
		user.GET("profile", controllers.ShowAuthUserProfile)
		user.PUT("profile", controllers.UpdateAuthUserProfile)
		user.GET("application-history", controllers.AuthUserApplicationHistory)
		user.POST("apply", controllers.ApplicationStore)

	}

}
