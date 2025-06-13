package models

import "gorm.io/gorm"

type Cibil struct {
	gorm.Model

	Slug    string  `json:"slug" gorm:"type:varchar(255);not null"`
	Score   string  `json:"score" gorm:"type:varchar(100);not null"`
	Vendor  string  `json:"vendor" gorm:"type:varchar(255);not null"`
	UserID  *string `json:"user_id,omitempty" gorm:"type:varchar(255);default:null"` // nullable
	Name    *string `json:"name,omitempty" gorm:"type:varchar(255);default:null"`    // nullable
	Email   *string `json:"email,omitempty" gorm:"type:varchar(255);default:null"`   // nullable
	Mobile  *string `json:"mobile,omitempty" gorm:"type:varchar(20);default:null"`   // nullable
	PanCard string  `json:"pan_card" gorm:"type:varchar(20);not null"`
}
