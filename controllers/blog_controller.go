package controllers

import (
	"backend/config"
	"backend/models"
	request "backend/request/blog"
	response "backend/response/blog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jaswdr/faker"
)

func BlogIndex(context *gin.Context) {

	var blogs []models.Blog
	config.DB.Find(&blogs)

	var blogRes []response.BlogResponse

	for _, blog := range blogs {
		blogRes = append(blogRes, response.BlogResponse{
			Slug:    blog.Slug,
			Title:   blog.Title,
			Summary: blog.Summary,
			Content: blog.Content,
			Image:   blog.Image,
			Status:  blog.Status,
		})
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Blog Index",
			"blogs":   blogRes,
		},
	)
}

func BlogShow(context *gin.Context) {

	slug := context.Param("slug")

	var blog models.Blog

	if err := config.DB.Where("slug=?", slug).First(&blog).Error; err != nil {

		context.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "Blog not Found",
			},
		)
	}

	blogResponse := response.BlogResponse{
		Slug:    blog.Slug,
		Title:   blog.Title,
		Summary: blog.Summary,
		Content: blog.Content,
		Image:   blog.Image,
		Status:  blog.Status,
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "blog show",
			"data":    blogResponse,
		},
	)
}

func BlogStore(context *gin.Context) {

	var req request.CreateBlogRequest

	if errReq := context.ShouldBindBodyWithJSON(&req); errReq != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": errReq.Error()})
		return
	}

	faker := faker.New()

	baseURL := "https://picsum.photos/id/"
	randomID := generateRandomNumber(1, 1000)
	imageURL := baseURL + strconv.Itoa(randomID) + "/370/240"

	blog := models.Blog{
		Slug:    faker.Internet().Slug(),
		Image:   imageURL,
		Title:   req.Title,
		Summary: req.Summary,
		Content: req.Content,
	}

	result := config.DB.Create(&blog)

	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	response := response.BlogResponse{
		Slug:    blog.Slug,
		Title:   blog.Title,
		Summary: blog.Summary,
		Content: blog.Content,
		Image:   blog.Image,
		Status:  blog.Status,
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "blog Store",
			"data":    response,
		},
	)
}

func BlogUpdate(context *gin.Context) {

	slug := context.Param("slug")

	var blog models.Blog

	if err := config.DB.Where("slug=?", slug).First(&blog).Error; err != nil {
		context.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "blog not Found",
			},
		)
	}

	var req request.UpdateBlogRequest

	if errReq := context.ShouldBindJSON(&req); errReq != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": errReq.Error()})
		return
	}

	updateErr := config.DB.Model(&blog).Updates(models.Blog{
		Title:   req.Title,
		Summary: req.Summary,
		Content: req.Content,
	}).Error

	if updateErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": updateErr.Error()})
		return
	}

	blogsResponse := response.BlogResponse{
		Slug:    blog.Slug,
		Title:   blog.Title,
		Summary: blog.Summary,
		Content: blog.Content,
		Image:   blog.Image,
		Status:  blog.Status,
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "blog Update",
			"data":    blogsResponse,
		},
	)
}

func BlogDestroy(context *gin.Context) {

	slug := context.Param("slug")

	var blog models.Blog

	if err := config.DB.Where("slug=?", slug).First(&blog).Error; err != nil {
		context.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "blog not Found",
			},
		)
		return
	}

	config.DB.Delete(&blog)
	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "blog delete successfully ",
		},
	)
}
