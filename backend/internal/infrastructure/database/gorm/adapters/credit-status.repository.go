package adapters

import (
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/ports"
	"gorm.io/gorm"
)

type CreditStatusGormRepository struct {
	db *gorm.DB
}

func NewCreditStatusGormRepository(db *gorm.DB) ports.CreditStatusRepository {
	return &CreditStatusGormRepository{
		db: db,
	}
}

func (r *CreditStatusGormRepository) FindAll() ([]models.CreditStatus, error) {
	var creditStatuses []models.CreditStatus
	if err := r.db.Find(&creditStatuses).Error; err != nil {
		return nil, err
	}
	return creditStatuses, nil
}

func (r *CreditStatusGormRepository) FindByID(id uint) (*models.CreditStatus, error) {
	var creditStatus models.CreditStatus
	if err := r.db.First(&creditStatus, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
	}
	return &creditStatus, nil
}
