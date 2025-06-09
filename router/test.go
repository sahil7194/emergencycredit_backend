package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerTestRoutes(rg *gin.RouterGroup) {

	rg.GET("ping", func(contex *gin.Context) {

		contex.JSON(http.StatusOK, gin.H{"message": "server  working fine"})
	})

}
