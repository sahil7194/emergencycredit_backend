package main

import (
	"backend/config"
	"backend/router"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Fail to load Env")
	}

	appEnv := os.Getenv("APP_ENV")

	gin.SetMode(appEnv)

	var applicationName = os.Getenv("APP_NAME")

	config.ConnectDatabase()

	fmt.Println("Application Name:", applicationName)
}

func main() {

	// faker := faker.New()

	// appNumber := generateApplicationNumber("Union Bank", 1234)

	// fmt.Println(faker.Internet().Slug())
	// fmt.Println(appNumber)

	router.InitRoutes()
}

func generateApplicationNumber(bankName string, userId int) string {
	// Format current time as "YYYYMMDDHHMMSS"
	today := time.Now().Format("20060102150405")

	// Convert bank name to uppercase and take first 3 letters
	bankName = strings.ToUpper(bankName)
	if len(bankName) > 3 {
		bankName = bankName[:3]
	}

	// Construct the application number
	return fmt.Sprintf("AP%s%s%d", today, bankName, userId)
}
