package router

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func InitRoutes() {

	r := gin.Default()

	api := r.Group("/api")

	registerTestRoutes(api)

	registerGuestUserRoutes(api)

	registerUtilsRoutes(api)

	registerAuthRoutes(api)

	registerUserRoutes(api)

	registerAgentRoutes(api)

	registerCrmRoutes(api)

	var application_url_with_port strings.Builder

	application_url_with_port.WriteString(os.Getenv("APP_HOST"))
	application_url_with_port.WriteString(":")
	application_url_with_port.WriteString(os.Getenv("APP_PORT"))

	r.Run(application_url_with_port.String())
}
