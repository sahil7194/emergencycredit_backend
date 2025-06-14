package controllers

import (
	"backend/config"
	"backend/models"
	request "backend/request/bank"
	response "backend/response/bank"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jaswdr/faker"
)

func BankIndex(context *gin.Context) {

	var banks []models.Bank

	config.DB.Find(&banks)

	var banksResponse []response.BankResponse

	for _, bank := range banks {

		banksResponse = append(banksResponse, response.BankResponse{
			Slug:       bank.Slug,
			Name:       bank.Name,
			Details:    bank.Details,
			Image:      bank.Image,
			VendorCode: bank.VendorCode,
		})
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Bank Index",
			"data":    banksResponse,
		},
	)
}

func BankShow(context *gin.Context) {

	slug := context.Param("slug")

	var bank models.Bank

	if err := config.DB.Where("slug=?", slug).First(&bank).Error; err != nil {

		context.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "Bank not Found",
			},
		)
	}

	banksResponse := response.BankResponse{
		Slug:       bank.Slug,
		Name:       bank.Name,
		Details:    bank.Details,
		Image:      bank.Image,
		VendorCode: bank.VendorCode,
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Bank show",
			"data":    banksResponse,
		},
	)
}

func BankStore(context *gin.Context) {

	var req request.CreateBankRequest

	if errReq := context.ShouldBindBodyWithJSON(&req); errReq != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": errReq.Error()})
		return
	}

	faker := faker.New()

	baseURL := "https://picsum.photos/id/"
	randomID := generateRandomNumber(1, 1000)
	imageURL := baseURL + strconv.Itoa(randomID) + "/370/240"

	bank := models.Bank{
		Name:       req.Name,
		Slug:       faker.Internet().Slug(),
		Details:    req.Details,
		Image:      imageURL,
		VendorCode: faker.Internet().Slug(),
	}

	result := config.DB.Create(&bank)

	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	response := response.BankResponse{
		Slug:       bank.Slug,
		Name:       bank.Name,
		Details:    bank.Details,
		Image:      bank.Image,
		VendorCode: bank.VendorCode,
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Bank Store",
			"data":    response,
		},
	)
}

func BankUpdate(context *gin.Context) {

	slug := context.Param("slug")

	var bank models.Bank

	if err := config.DB.Where("slug=?", slug).First(&bank).Error; err != nil {
		context.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "Bank not Found",
			},
		)
	}

	var req request.UpdateBankRequest

	if errReq := context.ShouldBindJSON(&req); errReq != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": errReq.Error()})
		return
	}

	updateErr := config.DB.Model(&bank).Updates(models.Bank{
		Name:    req.Name,
		Details: req.Details,
	}).Error

	if updateErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": updateErr.Error()})
		return
	}

	banksResponse := response.BankResponse{
		Slug:       bank.Slug,
		Name:       bank.Name,
		Details:    bank.Details,
		Image:      bank.Image,
		VendorCode: bank.VendorCode,
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Bank Update",
			"data":    banksResponse,
		},
	)
}

func BankDestroy(context *gin.Context) {

	slug := context.Param("slug")

	var bank models.Bank

	if err := config.DB.Where("slug=?", slug).First(&bank).Error; err != nil {
		context.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "bank not Found",
			},
		)
		return
	}

	config.DB.Delete(&bank)
	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "bank delete successfully ",
		},
	)
}

func generateRandomNumber(min int, max int) int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return r.Intn(max-min+1) + min
}
