package router

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func registerAuthRoutes(rg *gin.RouterGroup) {

	rg.POST("/signin", controllers.SignIn)
	rg.POST("/signup", controllers.SignUp)
	rg.POST("/forget-password", controllers.ForgetPassword)
	rg.POST("/reset-password", controllers.ResetPassword)

}
