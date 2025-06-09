package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OtpRequest(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Otp Request",
		},
	)
}

func OtpVerify(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Otp verify",
		},
	)
}

func OtpIndex(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Otp Index",
		},
	)
}

func OtpShow(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Otp show",
		},
	)
}

func OtpStore(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Otp Store",
		},
	)
}

func OtpUpdate(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Otp Update",
		},
	)
}

func OtpDestroy(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Otp Destroy",
		},
	)
}
