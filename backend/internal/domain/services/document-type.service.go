package services

import (
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/database"
	"gorm.io/gorm"
)

type DocumentTypeService struct {
	DB *gorm.DB
}

func (s *DocumentTypeService) GetAllCreditStatuses() ([]models.DocumentType, error) {
	var documentTypes []models.DocumentType
	err := database.DB.Find(&documentTypes).Error
	if err != nil {
		return nil, err
	}
	return documentTypes, nil
}
