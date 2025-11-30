package migrations

import (
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"gorm.io/gorm"
)

func AutoMigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.DocumentType{},
		&models.Asset{},
		&models.Customer{},
		&models.CreditRequest{},
		&models.CustomerAsset{},
		&models.Role{},
	)
}
