package ports

import "github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"

type CustomerAssetRepository interface {
	FindAll(creditRequestID *uint) ([]models.CustomerAsset, error)
	FindByID(id uint) (*models.CustomerAsset, error)
	CountByCreditRequestID(creditRequestID uint) (int64, error)
	Create(ca *models.CustomerAsset) error
	Update(id uint, data *models.CustomerAsset) (*models.CustomerAsset, error)
	Delete(id uint) error
}
