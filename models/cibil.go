package models

import "gorm.io/gorm"

type Cibil struct {
	gorm.Model

	Slug    string `json:"slug" gorm:"type:varchar(255);not null"`
	Score   string `json:"score" gorm:"type:varchar(100);not null"`
	Vendor  string `json:"vendor" gorm:"type:varchar(255);not null"`
	UserID  string `json:"user_id" gorm:"type:varchar(255);default:null"`
	Name    string `json:"name" gorm:"type:varchar(255);default:null"`
	Email   string `json:"email" gorm:"type:varchar(255);default:null"`
	Mobile  string `json:"mobile" gorm:"type:varchar(20);default:null"`
	PanCard string `json:"pan_card" gorm:"type:varchar(20);not null"`
}
