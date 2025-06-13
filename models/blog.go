package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model

	Title   string `json:"title" gorm:"type:varchar(255);not null"`
	Slug    string `json:"slug" gorm:"type:varchar(255);unique;not null"`
	Summary string `json:"summary" gorm:"type:text;not null"`
	Content string `json:"content" gorm:"type:text;not null"`
	Image   string `json:"image,omitempty" gorm:"type:varchar(255);default:null"`
	Status  string `json:"status,omitempty" gorm:"type:varchar(100);default:null"`
}
