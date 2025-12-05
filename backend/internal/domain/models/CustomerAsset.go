package models

import "time"

type CustomerAsset struct {
	ID              uint          `gorm:"primaryKey" json:"ID"`
	CreatedAt       time.Time     `json:"CreatedAt"`
	UpdatedAt       time.Time     `json:"UpdatedAt"`
	CreditRequestID uint          `gorm:"not null" json:"creditRequestId"`
	CreditRequest   CreditRequest `gorm:"foreignKey:CreditRequestID" json:"-"`
	AssetID         uint          `gorm:"not null" json:"assetId"`
	CustomerID      uint          `gorm:"not null" json:"customerId"`
	MarketValue     float64       `gorm:"not null" json:"marketValue"`
	Description     string        `gorm:"size:255" json:"description"`
	Status          bool          `gorm:"default:true" json:"status"`
}
