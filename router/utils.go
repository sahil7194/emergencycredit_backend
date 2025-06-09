package router

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func registerUtilsRoutes(rg *gin.RouterGroup) {

	rg.GET("/states", controllers.StateIndexForUtils)
	rg.GET("/city", controllers.CityIndexFilterByStateId)

	//routes for cibil

	rg.POST("/cibil/personal", controllers.CibilSavePersonalDetails)
	rg.POST("/cibil/request-otp", controllers.OtpRequest)
	rg.POST("/cibil/verify-otp", controllers.OtpVerify)
	rg.POST("/cibil/address", controllers.CibilSaveAddressDetails)
	rg.POST("/cibil/check", controllers.CibilCheck)

}
