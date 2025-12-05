package models

import "time"

type Customer struct {
	ID             uint            `gorm:"primaryKey" json:"ID"`
	CreatedAt      time.Time       `json:"CreatedAt"`
	UpdatedAt      time.Time       `json:"UpdatedAt"`
	Name           string          `json:"name"`
	Email          string          `json:"email"`
	PhoneNumber    string          `json:"phoneNumber"`
	DocumentNumber string          `json:"documentNumber"`
	DocumentTypeId uint            `json:"documentTypeId"`
	DocumentType   DocumentType    `gorm:"foreignKey:DocumentTypeId;references:ID" json:"-"`
	MonthlyIncome  float64         `gorm:"not null" json:"monthlyIncome"`
	CreatedByID    uint            `json:"createdById"`
	CreatedBy      User            `gorm:"foreignKey:CreatedByID;references:ID" json:"-"`
	Status         bool            `gorm:"default:true" json:"status"`
	CreditRequests []CreditRequest `gorm:"foreignKey:CustomerID" json:"-"`
}
