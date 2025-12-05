package models

import "time"

type DocumentType struct {
	ID          uint      `gorm:"primaryKey" json:"ID"`
	CreatedAt   time.Time `json:"CreatedAt"`
	UpdatedAt   time.Time `json:"UpdatedAt"`
	Code        string    `gorm:"unique;not null" json:"code"`
	Description string    `gorm:"not null" json:"description"`
	Status      bool      `gorm:"default:true" json:"status"`
}
