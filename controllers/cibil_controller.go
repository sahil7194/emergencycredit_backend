package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CibilCheck(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Cibil Check",
		},
	)
}

func CibilIndex(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Cibil Index",
		},
	)
}

func CibilShow(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Cibil show",
		},
	)
}

func CibilStore(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Cibil Store",
		},
	)
}

func CibilUpdate(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Cibil Update",
		},
	)
}

func CibilDestroy(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Cibil Destroy",
		},
	)
}
