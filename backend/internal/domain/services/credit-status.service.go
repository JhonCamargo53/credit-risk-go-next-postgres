package services

import (
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/database"
	"gorm.io/gorm"
)

type CreditStatusService struct {
	DB *gorm.DB
}

func (s *CreditStatusService) GetAllCreditStatuses() ([]models.CreditStatus, error) {
	var creditStatuses []models.CreditStatus
	err := database.DB.Find(&creditStatuses).Error
	if err != nil {
		return nil, err
	}
	return creditStatuses, nil
}
