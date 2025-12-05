package documentType

import (
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/ports"
)

type DocumentTypeService struct {
	repo ports.DocumentTypeRepository
}

func NewDocumentTypeService(repo ports.DocumentTypeRepository) *DocumentTypeService {
	return &DocumentTypeService{
		repo: repo,
	}
}

func (s *DocumentTypeService) GetAllDocumentTypes() ([]models.DocumentType, error) {
	return s.repo.FindAll()
}
