package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name"`
	RoleId   uint   `gorm:"not null;index" json:"roleId"`
	Role     Role   `gorm:"foreignKey:RoleId;references:ID" json:"-"`
	Email    string `gorm:"not null;unique" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Status   bool   `gorm:"default:true" json:"status"`
}
