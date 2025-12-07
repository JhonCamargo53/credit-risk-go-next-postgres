package asset

import (
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/ports"
)

type AssetService struct {
	repository ports.AssetRepository
}

func NewAssetService(repository ports.AssetRepository) *AssetService {
	return &AssetService{
		repository: repository,
	}
}

func (service *AssetService) GetAllAssets() ([]models.Asset, error) {
	return service.repository.FindAll()
}
