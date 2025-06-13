package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model

	Address string  `json:"address" gorm:"type:text;not null"`
	PinCode string  `json:"pin_code" gorm:"type:varchar(20);not null"`
	CityID  *string `json:"city_id,omitempty" gorm:"type:varchar(255);default:null"`
	StateID *string `json:"state_id,omitempty" gorm:"type:varchar(255);default:null"`
	UserID  *string `json:"user_id,omitempty" gorm:"type:varchar(255);default:null"`
}
