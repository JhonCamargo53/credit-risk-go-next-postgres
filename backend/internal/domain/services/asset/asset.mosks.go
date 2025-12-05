package asset

import (
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/ports"
)

type MockAssetRepository struct {
	Assets      []models.Asset
	ErrFindAll  error
	ErrFindByID error
}

var _ ports.AssetRepository = (*MockAssetRepository)(nil)

func NewMockAssetRepository(assets []models.Asset) *MockAssetRepository {
	return &MockAssetRepository{
		Assets: assets,
	}
}

func (m *MockAssetRepository) FindAll() ([]models.Asset, error) {
	if m.ErrFindAll != nil {
		return nil, m.ErrFindAll
	}
	return m.Assets, nil
}

func (m *MockAssetRepository) FindByID(id uint) (*models.Asset, error) {
	if m.ErrFindByID != nil {
		return nil, m.ErrFindByID
	}
	for _, a := range m.Assets {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, nil
}
