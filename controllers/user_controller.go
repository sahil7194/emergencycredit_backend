package controllers

import (
	"backend/config"
	"backend/models"
	request "backend/request/user"
	response "backend/response/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaswdr/faker"
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

	var users []models.User
	config.DB.Find(&users)

	var userRes []response.UserResponse

	for _, user := range users {
		userRes = append(userRes, response.UserResponse{
			Slug:        user.Slug,
			Name:        user.Name,
			Email:       user.Email,
			Mobile:      user.Mobile,
			DateOfBirth: user.DateOfBirth,
			Gender:      user.Gender,
			Type:        user.Type,
			Pan:         user.Pan,
		})
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "User Index",
			"users":   userRes,
		},
	)
}

func UserShow(context *gin.Context) {

	slug := context.Param("slug")

	var user models.User

	if err := config.DB.Where("slug=?", slug).First(&user).Error; err != nil {
		context.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "User not Found",
			},
		)
	}

	response := response.UserResponse{
		Slug:        user.Slug,
		Name:        user.Name,
		Email:       user.Email,
		Mobile:      user.Mobile,
		DateOfBirth: user.DateOfBirth,
		Gender:      user.Gender,
		Type:        user.Type,
		Pan:         user.Pan,
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "User show",
			"user":    response,
		},
	)
}

func UserStore(context *gin.Context) {

	var req request.CreateUserRequest

	if errReq := context.ShouldBindJSON(&req); errReq != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": errReq.Error()})
		return
	}

	faker := faker.New()

	user := models.User{
		Slug:        faker.Internet().Slug(),
		Name:        req.Name,
		Password:    req.Password,
		Email:       req.Email,
		Mobile:      req.Mobile,
		DateOfBirth: req.DateOfBirth,
		Gender:      req.Gender,
		Type:        req.Type,
		Pan:         req.Pan,
	}

	result := config.DB.Create(&user)

	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	response := response.UserResponse{
		Slug:        user.Slug,
		Name:        user.Name,
		Email:       user.Email,
		Mobile:      user.Mobile,
		DateOfBirth: user.DateOfBirth,
		Gender:      user.Gender,
		Type:        user.Type,
		Pan:         user.Pan,
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "User Store",
			"user":    response,
		},
	)
}

func UserUpdate(context *gin.Context) {

	slug := context.Param("slug")

	var user models.User

	if err := config.DB.Where("slug=?", slug).First(&user).Error; err != nil {
		context.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "User not Found",
			},
		)
		return
	}

	var req request.UpdateUserRequest

	if errReq := context.ShouldBindJSON(&req); errReq != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": errReq.Error()})
		return
	}

	updateErr := config.DB.Model(&user).Updates(models.User{
		Password:    req.Password,
		Name:        req.Name,
		Email:       req.Email,
		Mobile:      req.Mobile,
		DateOfBirth: req.DateOfBirth,
		Gender:      req.Gender,
		Type:        req.Type,
		Pan:         req.Pan,
	}).Error

	if updateErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": updateErr.Error()})
		return
	}

	response := response.UserResponse{
		Slug:        user.Slug,
		Name:        user.Name,
		Email:       user.Email,
		Mobile:      user.Mobile,
		DateOfBirth: user.DateOfBirth,
		Gender:      user.Gender,
		Type:        user.Type,
		Pan:         user.Pan,
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "User Update Successfully",
			"user":    response,
		},
	)
}

func UserDestroy(context *gin.Context) {

	slug := context.Param("slug")

	var user models.User

	if err := config.DB.Where("slug=?", slug).First(&user).Error; err != nil {
		context.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "User not Found",
			},
		)
		return
	}

	config.DB.Delete(&user)
	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "User delete successfully ",
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
