package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignIn(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "signin",
		},
	)
}

func SignUp(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "signup ",
		},
	)
}

func ForgetPassword(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "forget password",
		},
	)
}

func ResetPassword(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "rest password",
		},
	)
}
