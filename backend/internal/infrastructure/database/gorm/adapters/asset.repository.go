package adapters

import (
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/ports"
	"gorm.io/gorm"
)

type AssetGormRepository struct {
	db *gorm.DB
}

func NewAssetGormRepository(db *gorm.DB) ports.AssetRepository {
	return &AssetGormRepository{
		db: db,
	}
}

func (r *AssetGormRepository) FindAll() ([]models.Asset, error) {
	var assets []models.Asset
	if err := r.db.Find(&assets).Error; err != nil {
		return nil, err
	}
	return assets, nil
}

func (r *AssetGormRepository) FindByID(id uint) (*models.Asset, error) {
	var asset models.Asset
	if err := r.db.First(&asset, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &asset, nil
}
