package controllers

import (
	"backend/request/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignIn(context *gin.Context) {

	var request auth.LoginRequest

	if err := context.ShouldBindBodyWithJSON(&request); err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
