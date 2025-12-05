package models

import (
	"time"

	"gorm.io/gorm"
)

type DocumentType struct {
	ID          uint           `gorm:"primaryKey" json:"ID"`
	CreatedAt   time.Time      `json:"CreatedAt"`
	UpdatedAt   time.Time      `json:"UpdatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Code        string         `gorm:"unique;not null" json:"code"`
	Description string         `gorm:"not null" json:"description"`
	Status      bool           `gorm:"default:true" json:"status"`
}
