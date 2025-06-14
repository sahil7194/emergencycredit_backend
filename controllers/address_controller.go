package controllers

import (
	"backend/config"
	"backend/models"
	request "backend/request/check-cibil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CibilSaveAddressDetails(context *gin.Context) {

	var request request.SaveAddressDetails

	if errReq := context.ShouldBindJSON(&request); errReq != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": errReq.Error()})
		return
	}

	saveAddress := models.Address{
		Address: request.Address,
		PinCode: request.PinCode,
		CityID:  request.CityID,
		StateID: request.StateID,
		UserID:  "1",
	}

	config.DB.Create(&saveAddress)

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
