package services

import (
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/database"
	"gorm.io/gorm"
)

type AssetService struct {
	DB *gorm.DB
}

func (s *AssetService) GetAllAssets() ([]models.Asset, error) {
	var assets []models.Asset
	err := database.DB.Find(&assets).Error
	if err != nil {
		return nil, err
	}
	return assets, nil
}
