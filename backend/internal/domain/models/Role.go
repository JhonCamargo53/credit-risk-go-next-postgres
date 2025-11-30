package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name   string `gorm:"unique;not null" json:"name"`
	Access int    `gorm:"not null" json:"access"`
	Status bool   `gorm:"default:true" json:"status"`
}
