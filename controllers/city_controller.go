package controllers

import (
	"backend/config"
	"backend/models"
	request "backend/request/city"
	response "backend/response/city"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CityIndex(context *gin.Context) {

	var cities []models.City

	config.DB.Find(&cities)

	var cityResponse []response.CityResponse

	for _, city := range cities {

		cityResponse = append(cityResponse, response.CityResponse{
			Id:   city.ID,
			Name: city.Name,
		})
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "city Index",
			"data":    cityResponse,
		},
	)
}

func CityIndexFilterByStateId(context *gin.Context) {

	var cities []models.City

	state_id := context.Query("state_id")

	if state_id == "" {
		context.JSON(http.StatusOK, gin.H{"message": "state id required"})
		return
	}

	if err := config.DB.Where("state_id=?", state_id).Find(&cities).Error; err != nil {
		context.JSON(http.StatusOK, gin.H{"message": "not city found"})
		return
	}

	var cityResponse []response.CityResponse

	for _, city := range cities {

		cityResponse = append(cityResponse, response.CityResponse{
			Id:   city.ID,
			Name: city.Name,
		})
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "city Index",
			"data":    cityResponse,
		},
	)
}

func CityShow(context *gin.Context) {

	slug := context.Param("slug")

	var city models.City

	if err := config.DB.Where("id=?", slug).First(&city).Error; err != nil {

		context.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "city not Found",
			},
		)
	}

	cityResponse := response.CityResponse{
		Id:   city.ID,
		Name: city.Name,
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "city show",
			"data":    cityResponse,
		},
	)
}

func CityStore(context *gin.Context) {

	var req request.CreateCityRequest

	if errReq := context.ShouldBindBodyWithJSON(&req); errReq != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": errReq.Error()})
		return
	}

	city := models.City{
		Name:    req.Name,
		StateID: req.StateId,
	}

	result := config.DB.Create(&city)

	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	response := response.CityResponse{
		Id:   city.ID,
		Name: city.Name,
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "state Store",
			"data":    response,
		},
	)
}

func CityUpdate(context *gin.Context) {

	slug := context.Param("slug")

	var city models.City

	if err := config.DB.Where("id=?", slug).First(&city).Error; err != nil {
		context.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "city not Found",
			},
		)
	}

	var req request.UpdateCityRequest

	if errReq := context.ShouldBindJSON(&req); errReq != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": errReq.Error()})
		return
	}

	updateErr := config.DB.Model(&city).Updates(models.City{
		Name: req.Name,
	}).Error

	if updateErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": updateErr.Error()})
		return
	}

	cityResponse := response.CityResponse{
		Id:   city.ID,
		Name: city.Name,
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "city Update",
			"data":    cityResponse,
		},
	)
}

func CityDestroy(context *gin.Context) {

	slug := context.Param("slug")

	var city models.City

	if err := config.DB.Where("id=?", slug).First(&city).Error; err != nil {
		context.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "city not Found",
			},
		)
		return
	}

	config.DB.Delete(&city)
	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "city delete successfully ",
		},
	)
}
