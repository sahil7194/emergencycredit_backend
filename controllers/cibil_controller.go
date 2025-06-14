package controllers

import (
	"backend/config"
	"backend/models"
	request "backend/request/check-cibil"
	response "backend/response/cibil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jaswdr/faker"
)

func CibilCheck(context *gin.Context) {

	var req request.CheckCibil

	if errReq := context.ShouldBindJSON(&req); errReq != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": errReq.Error()})
		return
	}

	faker := faker.New()

	var user models.User

	if err := config.DB.Where("slug=?", req.Slug).First(&user).Error; err != nil {
		context.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "User not Found",
			},
		)

		return
	} else {

		cibil := models.Cibil{
			Slug:    faker.Internet().Slug(),
			Score:   "123",
			Vendor:  "abc",
			UserID:  strconv.Itoa(int(user.ID)),
			Name:    user.Name,
			Email:   user.Name,
			Mobile:  user.Mobile,
			PanCard: user.Pan,
		}

		result := config.DB.Create(&cibil)

		if result.Error != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		response := response.CibilResponse{
			Slug:    cibil.Slug,
			Score:   cibil.Score,
			UserId:  cibil.UserID,
			Name:    cibil.Name,
			Email:   cibil.Email,
			Mobile:  cibil.Mobile,
			PanCard: cibil.PanCard,
		}

		context.JSON(
			http.StatusOK,
			gin.H{
				"success": true,
				"message": "data found successfully",
				"cibil":   response,
			},
		)
	}
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
