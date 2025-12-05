package models

import "time"

type CreditStatus struct {
	ID        uint      `gorm:"primaryKey" json:"ID"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	Name      string    `gorm:"not null;unique" json:"name"`
	Status    bool      `gorm:"default:true" json:"status"`
}
