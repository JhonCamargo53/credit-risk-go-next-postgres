package ports

import "github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"

type AssetRepository interface {
	FindByID(id uint) (*models.Asset, error)
	FindAll() ([]models.Asset, error)
}
