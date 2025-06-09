package router

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func registerGuestUserRoutes(rg *gin.RouterGroup) {

	rg.GET("/blogs", controllers.BlogIndex)
	rg.GET("/blogs/:slug", controllers.BlogShow)

	rg.GET("/schemes", controllers.SchemeIndex)
	rg.GET("/schemes/:slug", controllers.SchemeShow)

}
