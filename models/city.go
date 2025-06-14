package models

import "gorm.io/gorm"

type City struct {
	gorm.Model

	Name    string `json:"name" gorm:"type:varchar(255);not null"`
	StateID string `json:"state_id,omitempty" gorm:"type:varchar(255);default:null"`
}
