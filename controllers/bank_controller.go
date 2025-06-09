package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BankIndex(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Bank Index",
		},
	)
}

func BankShow(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Bank show",
		},
	)
}

func BankStore(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Bank Store",
		},
	)
}

func BankUpdate(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Bank Update",
		},
	)
}

func BankDestroy(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Bank Destroy",
		},
	)
}
