package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CrmMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("auth_token")

		if token != "crm_auth_token" {

			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

	}
}
