package models

import (
	"time"

	"gorm.io/gorm"
)

type CreditStatus struct {
	ID        uint           `gorm:"primaryKey" json:"ID"`
	CreatedAt time.Time      `json:"CreatedAt"`
	UpdatedAt time.Time      `json:"UpdatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Name      string         `gorm:"not null;unique" json:"name"`
	Status    bool           `gorm:"default:true" json:"status"`
}
