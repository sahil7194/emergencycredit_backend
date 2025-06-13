package models

import "gorm.io/gorm"

type Bank struct {
	Slug       string `json:"slug" gorm:"unique"`
	Name       string `json:"name"`
	VendorCode string `json:"vendor_code"`
	Details    string `json:"details"`
	Image      string `json:"image"`
	gorm.Model
}
