package models

import "gorm.io/gorm"

type DocumentType struct {
	gorm.Model
	Code        string `gorm:"unique;not null" json:"code"`
	Description string `gorm:"not null" json:"description"`
	Status      bool   `gorm:"default:true" json:"status"`
}
