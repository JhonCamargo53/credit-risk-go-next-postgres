package models

import (
	"time"

	"gorm.io/gorm"
)

type CreditRequest struct {
	ID              uint           `gorm:"primaryKey" json:"ID"`
	CreatedAt       time.Time      `json:"CreatedAt"`
	UpdatedAt       time.Time      `json:"UpdatedAt"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
	Amount          float64        `json:"amount"`
	TermMonths      int            `json:"termMonths"`
	CustomerID      uint           `gorm:"not null" json:"customerId"`
	Customer        Customer       `json:"-"`
	ProductType     string         `json:"productType"`
	CreditStatusID  uint           `gorm:"not null" json:"creditStatusId"`
	CreditStatus    CreditStatus   `json:"-"`
	RiskScore       float64        `gorm:"default:0" json:"riskScore"`
	RiskCategory    string         `json:"riskCategory"`
	RiskExplanation string         `json:"riskExplanation" gorm:"type:TEXT"`
}
