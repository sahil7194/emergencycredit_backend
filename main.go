package main

import (
	"backend/router"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	initApplication()

	router.InitRoutes()
}

func initApplication() {

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Fail to load Env")
	}

	var applicationName = os.Getenv("APP_NAME")

	fmt.Println("Application Name:", applicationName)
}
