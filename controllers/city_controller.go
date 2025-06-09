package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CityIndex(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "City Index",
		},
	)
}

func CityIndexFilterByStateId(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "City Index filter by State Id",
		},
	)
}

func CityShow(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "City show",
		},
	)
}

func CityStore(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "City Store",
		},
	)
}

func CityUpdate(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "City Update",
		},
	)
}

func CityDestroy(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "City Destroy",
		},
	)
}
