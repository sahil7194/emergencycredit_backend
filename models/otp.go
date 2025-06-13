package models

import "gorm.io/gorm"

type Otp struct {
	gorm.Model

	Type   string `json:"type" gorm:"type:varchar(255);not null"`
	Value  string `json:"value" gorm:"type:varchar(255);not null"`
	OTP    string `json:"otp" gorm:"type:varchar(255);not null"`
	Status string `json:"status" gorm:"type:varchar(100);not null"`
}
