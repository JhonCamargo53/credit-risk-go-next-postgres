package models

import "time"

type Role struct {
	ID        uint      `gorm:"primaryKey" json:"ID"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	Name      string    `gorm:"unique;not null" json:"name"`
	Access    int       `gorm:"not null" json:"access"`
	Status    bool      `gorm:"default:true" json:"status"`
}
