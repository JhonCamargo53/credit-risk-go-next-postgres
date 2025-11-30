package models

import "gorm.io/gorm"

type CreditStatus struct {
	gorm.Model
	Name   string `gorm:"not null;unique" json:"name"`
	Status bool   `gorm:"default:true" json:"status"`
}
