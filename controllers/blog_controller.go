package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BlogIndex(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Blog Index",
		},
	)
}

func BlogShow(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Blog show",
		},
	)
}

func BlogStore(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Blog Store",
		},
	)
}

func BlogUpdate(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Blog Update",
		},
	)
}

func BlogDestroy(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Blog Destroy",
		},
	)
}
