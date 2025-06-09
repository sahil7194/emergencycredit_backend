package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CibilSaveAddressDetails(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "save personal details",
		},
	)
}

func AddressIndex(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Address Index",
		},
	)
}

func AddressShow(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Address show",
		},
	)
}

func AddressStore(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Address Store",
		},
	)
}

func AddressUpdate(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Address Update",
		},
	)
}

func AddressDestroy(context *gin.Context) {

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Address Destroy",
		},
	)
}
