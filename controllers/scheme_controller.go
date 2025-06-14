package controllers

import (
	"backend/config"
	"backend/models"
	request "backend/request/scheme"
	response "backend/response/scheme"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jaswdr/faker"
)

func SchemeIndex(context *gin.Context) {

	var schemes []models.Scheme

	config.DB.Find(&schemes)

	var schemesResponse []response.SchemeResponse

	for _, scheme := range schemes {

		schemesResponse = append(schemesResponse, response.SchemeResponse{
			Summary:         scheme.Summary,
			Slug:            scheme.Slug,
			Title:           scheme.Title,
			Description:     scheme.Description,
			Image:           scheme.Image,
			RedirectionLink: scheme.RedirectionLink,
			MinInterestRate: scheme.MinInterestRate,
			MaxInterestRate: scheme.MaxInterestRate,
			MinCIBIL:        scheme.MinCibil,
			MaxCIBIL:        scheme.MaxCibil,
			MinTenure:       scheme.MinTenure,
			MaxTenure:       scheme.MaxTenure,
			MinAmount:       scheme.MinAmount,
			MaxAmount:       scheme.MaxAmount,
			Status:          scheme.Status,
		})
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Schemes Index",
			"data":    schemesResponse,
		},
	)
}

func SchemeShow(context *gin.Context) {
	slug := context.Param("slug")

	var scheme models.Scheme

	if err := config.DB.Where("slug=?", slug).First(&scheme).Error; err != nil {

		context.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "scheme not Found",
			},
		)
	}

	schemesResponse := response.SchemeResponse{
		Summary:         scheme.Summary,
		Slug:            scheme.Slug,
		Title:           scheme.Title,
		Description:     scheme.Description,
		Image:           scheme.Image,
		RedirectionLink: scheme.RedirectionLink,
		MinInterestRate: scheme.MinInterestRate,
		MaxInterestRate: scheme.MaxInterestRate,
		MinCIBIL:        scheme.MinCibil,
		MaxCIBIL:        scheme.MaxCibil,
		MinTenure:       scheme.MinTenure,
		MaxTenure:       scheme.MaxTenure,
		MinAmount:       scheme.MinAmount,
		MaxAmount:       scheme.MaxAmount,
		Status:          scheme.Status,
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "scheme show",
			"data":    schemesResponse,
		},
	)
}

func SchemeStore(context *gin.Context) {

	var req request.CreateSchemeRequest

	if errReq := context.ShouldBindBodyWithJSON(&req); errReq != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": errReq.Error()})
		return
	}

	faker := faker.New()

	baseURL := "https://picsum.photos/id/"
	randomID := generateRandomNumber(1, 1000)
	imageURL := baseURL + strconv.Itoa(randomID) + "/370/240"

	scheme := models.Scheme{
		Slug:            faker.Internet().Slug(),
		Image:           imageURL,
		Title:           req.Title,
		Summary:         req.Summary,
		Description:     req.Description,
		MaxAmount:       req.MaxAmount,
		MinAmount:       req.MinAmount,
		MaxCibil:        req.MaxCibil,
		MinCibil:        req.MinCibil,
		MaxInterestRate: req.MaxAmount,
		MinInterestRate: req.MinInterestRate,
		MaxTenure:       req.MaxTenure,
		MinTenure:       req.MinTenure,
		RedirectionLink: req.RedirectionLink,
		BankId:          req.BankId,
	}

	result := config.DB.Create(&scheme)

	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	response := response.SchemeResponse{
		Summary:         scheme.Summary,
		Slug:            scheme.Slug,
		Title:           scheme.Title,
		Description:     scheme.Description,
		Image:           scheme.Image,
		RedirectionLink: scheme.RedirectionLink,
		MinInterestRate: scheme.MinInterestRate,
		MaxInterestRate: scheme.MaxInterestRate,
		MinCIBIL:        scheme.MinCibil,
		MaxCIBIL:        scheme.MaxCibil,
		MinTenure:       scheme.MinTenure,
		MaxTenure:       scheme.MaxTenure,
		MinAmount:       scheme.MinAmount,
		MaxAmount:       scheme.MaxAmount,
		Status:          scheme.Status,
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Scheme Store",
			"data":    response,
		},
	)
}

func SchemeUpdate(context *gin.Context) {

	slug := context.Param("slug")

	var scheme models.Scheme

	if err := config.DB.Where("slug=?", slug).First(&scheme).Error; err != nil {
		context.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "scheme not Found",
			},
		)
	}

	var req request.UpdateSchemeRequest

	if errReq := context.ShouldBindJSON(&req); errReq != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": errReq.Error()})
		return
	}

	updateErr := config.DB.Model(&scheme).Updates(models.Scheme{
		Title:           req.Title,
		Summary:         req.Summary,
		Description:     req.Description,
		MaxAmount:       req.MaxAmount,
		MinAmount:       req.MinAmount,
		MaxCibil:        req.MaxCibil,
		MinCibil:        req.MinCibil,
		MaxInterestRate: req.MaxAmount,
		MinInterestRate: req.MinInterestRate,
		MaxTenure:       req.MaxTenure,
		MinTenure:       req.MinTenure,
		RedirectionLink: req.RedirectionLink,
		BankId:          req.BankId,
	}).Error

	if updateErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": updateErr.Error()})
		return
	}

	schemesResponse := response.SchemeResponse{
		Summary:         scheme.Summary,
		Slug:            scheme.Slug,
		Title:           scheme.Title,
		Description:     scheme.Description,
		Image:           scheme.Image,
		RedirectionLink: scheme.RedirectionLink,
		MinInterestRate: scheme.MinInterestRate,
		MaxInterestRate: scheme.MaxInterestRate,
		MinCIBIL:        scheme.MinCibil,
		MaxCIBIL:        scheme.MaxCibil,
		MinTenure:       scheme.MinTenure,
		MaxTenure:       scheme.MaxTenure,
		MinAmount:       scheme.MinAmount,
		MaxAmount:       scheme.MaxAmount,
		Status:          scheme.Status,
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "scheme Update",
			"data":    schemesResponse,
		},
	)
}

func SchemeDestroy(context *gin.Context) {

	slug := context.Param("slug")

	var scheme models.Scheme

	if err := config.DB.Where("slug=?", slug).First(&scheme).Error; err != nil {
		context.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "scheme not Found",
			},
		)
		return
	}

	config.DB.Delete(&scheme)
	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "scheme delete successfully ",
		},
	)
}
