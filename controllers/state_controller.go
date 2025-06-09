package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StateIndex(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "State Index",
		},
	)
}

func StateIndexForUtils(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "State Index for utils",
		},
	)
}

func StateShow(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "State show",
		},
	)
}

func StateStore(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "State Store",
		},
	)
}

func StateUpdate(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "State Update",
		},
	)
}

func StateDestroy(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "State Destroy",
		},
	)
}
