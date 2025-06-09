package router

import (
	"backend/controllers"
	"backend/middlewares"

	"github.com/gin-gonic/gin"
)

func registerAgentRoutes(rg *gin.RouterGroup) {

	agent := rg.Group("/agent")

	agent.Use(middlewares.AgentMiddleWare())
	{
		agent.GET("profile", controllers.ShowAuthUserProfile)
		agent.PUT("profile", controllers.UpdateAuthUserProfile)
		agent.GET("reference-history", controllers.AgentReferenceHistory)
	}

}
