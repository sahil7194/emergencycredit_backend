package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CibilSavePersonalDetails(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "save personal details",
		},
	)
}

func UserIndex(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "User Index",
		},
	)
}

func UserShow(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "User show",
		},
	)
}

func UserStore(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "User Store",
		},
	)
}

func UserUpdate(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "User Update",
		},
	)
}

func UserDestroy(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "User Destroy",
		},
	)
}

func ShowAuthUserProfile(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "show auth user user profile",
		},
	)
}

func UpdateAuthUserProfile(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "update auth user profile",
		},
	)
}
