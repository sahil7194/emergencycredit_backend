package models

import "gorm.io/gorm"

type State struct {
	gorm.Model

	Name string `json:"name" gorm:"type:varchar(255);not null"`
}
