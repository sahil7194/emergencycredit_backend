package models

import "gorm.io/gorm"

type Scheme struct {
	gorm.Model

	Summary         string `json:"summary" gorm:"type:text;not null"`
	Slug            string `json:"slug" gorm:"type:varchar(255);unique;not null"`
	Title           string `json:"title" gorm:"type:varchar(255);not null"`
	Description     string `json:"description" gorm:"type:text;not null"`
	Image           string `json:"image,omitempty" gorm:"type:varchar(255);default:null"`
	RedirectionLink string `json:"redirection_link" gorm:"type:text;not null"`

	MinInterestRate string `json:"min_interest_rate" gorm:"type:varchar(100)"`
	MaxInterestRate string `json:"max_interest_rate" gorm:"type:varchar(100)"`
	MinCIBIL        string `json:"min_cibil" gorm:"type:varchar(100)"`
	MaxCIBIL        string `json:"max_cibil" gorm:"type:varchar(100)"`
	MinTenure       string `json:"min_tenure" gorm:"type:varchar(100)"`
	MaxTenure       string `json:"max_tenure" gorm:"type:varchar(100)"`
	MinAmount       string `json:"min_amount" gorm:"type:varchar(100)"`
	MaxAmount       string `json:"max_amount" gorm:"type:varchar(100)"`
	Status          string `json:"status" gorm:"type:varchar(100)"`
}
