package models

import "time"

type Asset struct {
	ID             uint            `gorm:"primaryKey" json:"ID"`
	CreatedAt      time.Time       `json:"CreatedAt"`
	UpdatedAt      time.Time       `json:"UpdatedAt"`
	Name           string          `gorm:"unique;not null" json:"name"`
	Description    string          `gorm:"size:255;not null" json:"description"`
	CustomerAssets []CustomerAsset `gorm:"foreignKey:AssetID" json:"-"`
	Status         bool            `gorm:"default:true" json:"status"`
}
