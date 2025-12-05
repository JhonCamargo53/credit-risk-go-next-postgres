package adapters

import (
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/ports"
	"gorm.io/gorm"
)

type CustomerAssetGormRepository struct {
	db *gorm.DB
}

func NewCustomerAssetGormRepository(db *gorm.DB) ports.CustomerAssetRepository {
	return &CustomerAssetGormRepository{
		db: db,
	}
}

func (r *CustomerAssetGormRepository) FindAll(creditRequestID *uint) ([]models.CustomerAsset, error) {
	query := r.db.Model(&models.CustomerAsset{})

	if creditRequestID != nil {
		query = query.Where("credit_request_id = ?", *creditRequestID)
	}

	var customerAssets []models.CustomerAsset
	if err := query.Order("created_at desc").Find(&customerAssets).Error; err != nil {
		return nil, err
	}

	return customerAssets, nil
}

func (r *CustomerAssetGormRepository) FindByID(id uint) (*models.CustomerAsset, error) {
	var ca models.CustomerAsset
	if err := r.db.First(&ca, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &ca, nil
}

func (r *CustomerAssetGormRepository) CountByCreditRequestID(creditRequestID uint) (int64, error) {
	var count int64
	if err := r.db.Model(&models.CustomerAsset{}).Where("credit_request_id = ?", creditRequestID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *CustomerAssetGormRepository) Create(ca *models.CustomerAsset) error {
	return r.db.Create(ca).Error
}

func (r *CustomerAssetGormRepository) Update(id uint, data *models.CustomerAsset) (*models.CustomerAsset, error) {
	var ca models.CustomerAsset
	if err := r.db.First(&ca, id).Error; err != nil {
		return nil, err
	}

	if err := r.db.Model(&ca).Updates(data).Error; err != nil {
		return nil, err
	}

	return &ca, nil
}

func (r *CustomerAssetGormRepository) Delete(id uint) error {
	return r.db.Delete(&models.CustomerAsset{}, id).Error
}
