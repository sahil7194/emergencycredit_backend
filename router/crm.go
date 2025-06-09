package router

import (
	"backend/controllers"
	"backend/middlewares"

	"github.com/gin-gonic/gin"
)

func registerCrmRoutes(rg *gin.RouterGroup) {

	crm := rg.Group("/crm")

	crm.Use(middlewares.CrmMiddleWare())
	{

		user := crm.Group("/users")
		{
			user.GET("/", controllers.UserIndex)
			user.POST("/", controllers.UserStore)
			user.GET("/:slug", controllers.UserShow)
			user.PUT("/:slug", controllers.UserUpdate)
			user.DELETE("/:slug", controllers.UserDestroy)
		}

		banks := crm.Group("/banks")
		{
			banks.GET("/", controllers.BankIndex)
			banks.POST("/", controllers.BankStore)
			banks.GET("/:slug", controllers.BankShow)
			banks.PUT("/:slug", controllers.BankUpdate)
			banks.DELETE("/:slug", controllers.BankDestroy)
		}

		states := crm.Group("/states")
		{
			states.GET("/", controllers.StateIndex)
			states.POST("/", controllers.StateStore)
			states.GET("/:slug", controllers.StateShow)
			states.PUT("/:slug", controllers.StateUpdate)
			states.DELETE("/:slug", controllers.StateDestroy)
		}

		city := crm.Group("/city")
		{
			city.GET("/", controllers.CityIndex)
			city.POST("/", controllers.CityStore)
			city.GET("/:slug", controllers.CityShow)
			city.PUT("/:slug", controllers.CityUpdate)
			city.DELETE("/:slug", controllers.CityDestroy)
		}

		schemes := crm.Group("/schemes")
		{
			schemes.GET("/", controllers.SchemeIndex)
			schemes.POST("/", controllers.SchemeStore)
			schemes.GET("/:slug", controllers.SchemeIndex)
			schemes.PUT("/:slug", controllers.SchemeUpdate)
			schemes.DELETE("/:slug", controllers.SchemeDestroy)
		}

		logs := crm.Group("/log")
		{
			logs.GET("/otp", controllers.OtpIndex)
			logs.GET("/cibil", controllers.CibilIndex)
			logs.GET("/applications", controllers.ApplicationIndex)

		}
	}

}
