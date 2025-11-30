package models

import "gorm.io/gorm"

type Asset struct {
	gorm.Model
	Name           string          `gorm:"unique;not null" json:"name"`
	Description    string          `gorm:"size:255;not null" json:"description"`
	CustomerAssets []CustomerAsset `gorm:"foreignKey:AssetID" json:"-"`
	Status         bool            `gorm:"default:true" json:"status"`
}
