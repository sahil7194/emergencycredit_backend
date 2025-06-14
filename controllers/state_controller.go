package controllers

import (
	"backend/config"
	"backend/models"
	request "backend/request/state"
	response "backend/response/state"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StateIndex(context *gin.Context) {

	var states []models.State

	config.DB.Find(&states)

	var stateResponse []response.StateResponse

	for _, state := range states {

		stateResponse = append(stateResponse, response.StateResponse{
			Id:   state.ID,
			Name: state.Name,
		})
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "states Index",
			"data":    stateResponse,
		},
	)
}

func StateIndexForUtils(context *gin.Context) {

	var states []models.State

	config.DB.Find(&states)

	var stateResponse []response.StateResponse

	for _, state := range states {

		stateResponse = append(stateResponse, response.StateResponse{
			Id:   state.ID,
			Name: state.Name,
		})
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "states Index",
			"data":    stateResponse,
		},
	)
}

func StateShow(context *gin.Context) {

	slug := context.Param("slug")

	var state models.State

	if err := config.DB.Where("id=?", slug).First(&state).Error; err != nil {

		context.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "state not Found",
			},
		)
	}

	statesResponse := response.StateResponse{
		Id:   state.ID,
		Name: state.Name,
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "State show",
			"data":    statesResponse,
		},
	)
}

func StateStore(context *gin.Context) {

	var req request.CreateStateRequest

	if errReq := context.ShouldBindBodyWithJSON(&req); errReq != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": errReq.Error()})
		return
	}

	state := models.State{
		Name: req.Name,
	}

	result := config.DB.Create(&state)

	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	response := response.StateResponse{
		Id:   state.ID,
		Name: state.Name,
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "state Store",
			"data":    response,
		},
	)
}

func StateUpdate(context *gin.Context) {

	slug := context.Param("slug")

	var State models.State

	if err := config.DB.Where("id=?", slug).First(&State).Error; err != nil {
		context.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "State not Found",
			},
		)
	}

	var req request.UpdateStateRequest

	if errReq := context.ShouldBindJSON(&req); errReq != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": errReq.Error()})
		return
	}

	updateErr := config.DB.Model(&State).Updates(models.State{
		Name: req.Name,
	}).Error

	if updateErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": updateErr.Error()})
		return
	}

	StatesResponse := response.StateResponse{
		Id:   State.ID,
		Name: State.Name,
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "State Update",
			"data":    StatesResponse,
		},
	)
}

func StateDestroy(context *gin.Context) {

	slug := context.Param("slug")

	var state models.State

	if err := config.DB.Where("id=?", slug).First(&state).Error; err != nil {
		context.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "state not Found",
			},
		)
		return
	}

	config.DB.Delete(&state)
	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "state delete successfully ",
		},
	)
}
