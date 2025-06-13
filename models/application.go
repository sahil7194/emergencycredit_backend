package models

import "gorm.io/gorm"

type Application struct {
	gorm.Model

	ApplicationID string `json:"application_id" gorm:"type:varchar(255);not null"`
	Status        string `json:"status" gorm:"type:varchar(10);default:'0';not null"`
	Remarks       string `json:"remarks" gorm:"type:text;not null"`
	SchemeID      uint   `json:"scheme_id" gorm:"not null"`
	UserID        uint   `json:"user_id" gorm:"not null"`
	AgentID       *uint  `json:"agent_id,omitempty" gorm:"default:null"`
}
