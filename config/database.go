package config

import (
	"backend/models"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var databaseError error

	DB, databaseError = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if databaseError != nil {
		panic("faild to connect database")
	}

	println("Database connected Successfully")

	//migration
	DB.AutoMigrate(
		models.User{},
		models.Address{},
		models.State{},
		models.City{},
		models.Bank{},
		models.Scheme{},
		models.Application{},
		models.Cibil{},
		models.Blog{},
		models.Otp{},
	)

}
