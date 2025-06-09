package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApplicationIndex(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Application Index",
		},
	)
}

func ApplicationShow(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Application show",
		},
	)
}

func ApplicationStore(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Application Store",
		},
	)
}

func ApplicationUpdate(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Application Update",
		},
	)
}

func ApplicationDestroy(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Application Destroy",
		},
	)
}

func AuthUserApplicationHistory(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Auth User Application History",
		},
	)
}

func AgentReferenceHistory(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "agent reference history",
		},
	)
}
