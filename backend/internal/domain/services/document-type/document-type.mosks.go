package documentType

import (
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/ports"
)

type MockDocumentTypeRepository struct {
	Types         []models.DocumentType
	ErrFindAll    error
	ErrExistsByID error
}

var _ ports.DocumentTypeRepository = (*MockDocumentTypeRepository)(nil)

func NewMockDocumentTypeRepository(types []models.DocumentType) *MockDocumentTypeRepository {
	return &MockDocumentTypeRepository{
		Types: types,
	}
}

func (m *MockDocumentTypeRepository) FindAll() ([]models.DocumentType, error) {
	if m.ErrFindAll != nil {
		return nil, m.ErrFindAll
	}
	return m.Types, nil
}

func (m *MockDocumentTypeRepository) ExistsByID(id uint) (bool, error) {
	if m.ErrExistsByID != nil {
		return false, m.ErrExistsByID
	}
	for _, dt := range m.Types {
		if dt.ID == id {
			return true, nil
		}
	}
	return false, nil
}
