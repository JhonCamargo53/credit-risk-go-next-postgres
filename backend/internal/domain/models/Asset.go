package models

import (
	"time"

	"gorm.io/gorm"
)

type Asset struct {
	ID             uint            `gorm:"primaryKey" json:"ID"`
	CreatedAt      time.Time       `json:"CreatedAt"`
	UpdatedAt      time.Time       `json:"UpdatedAt"`
	DeletedAt      gorm.DeletedAt  `gorm:"index" json:"-"`
	Name           string          `gorm:"unique;not null" json:"name"`
	Description    string          `gorm:"size:255;not null" json:"description"`
	CustomerAssets []CustomerAsset `gorm:"foreignKey:AssetID" json:"-"`
	Status         bool            `gorm:"default:true" json:"status"`
}
