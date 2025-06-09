package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SchemeIndex(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Scheme Index",
		},
	)
}

func SchemeShow(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Scheme show",
		},
	)
}

func SchemeStore(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Scheme Store",
		},
	)
}

func SchemeUpdate(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Scheme Update",
		},
	)
}

func SchemeDestroy(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Scheme Destroy",
		},
	)
}
