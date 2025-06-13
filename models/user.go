package models

import (
	"gorm.io/gorm"
)

type User struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Password    string `json:"-"`
	Email       string `json:"email" gorm:"unique"`
	Mobile      string `json:"mobile" gorm:"unique"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
	Type        string `json:"type"`
	Pan         string `json:"pan"`
	gorm.Model
}
