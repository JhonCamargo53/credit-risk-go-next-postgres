package creditStatus

import (
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/ports"
)

type MockCreditStatusRepository struct {
	Statuses    []models.CreditStatus
	ErrFindAll  error
	ErrFindByID error
}

var _ ports.CreditStatusRepository = (*MockCreditStatusRepository)(nil)

func NewMockCreditStatusRepository(initial []models.CreditStatus) *MockCreditStatusRepository {
	return &MockCreditStatusRepository{
		Statuses: initial,
	}
}

func (m *MockCreditStatusRepository) FindAll() ([]models.CreditStatus, error) {
	if m.ErrFindAll != nil {
		return nil, m.ErrFindAll
	}
	return m.Statuses, nil
}

func (m *MockCreditStatusRepository) FindByID(id uint) (*models.CreditStatus, error) {
	if m.ErrFindByID != nil {
		return nil, m.ErrFindByID
	}

	for _, cs := range m.Statuses {
		if cs.ID == id {
			return &cs, nil
		}
	}
	return nil, nil
}
