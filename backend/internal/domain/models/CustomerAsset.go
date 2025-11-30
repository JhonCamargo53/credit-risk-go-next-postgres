package models

import "gorm.io/gorm"

type CustomerAsset struct {
	gorm.Model
	CreditRequestID uint          `gorm:"not null" json:"creditRequestId"`
	CreditRequest   CreditRequest `gorm:"foreignKey:CreditRequestID" json:"-"`
	AssetID         uint          `gorm:"not null" json:"assetId"`
	CustomerID      uint          `gorm:"not null" json:"customerId"`
	MarketValue     float64       `gorm:"not null" json:"marketValue"`
	Description     string        `gorm:"size:255" json:"description"`
	Status          bool          `gorm:"default:true" json:"status"`
}
