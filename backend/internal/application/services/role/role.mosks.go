package role

import (
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/ports"
)

type MockRoleRepository struct {
	Roles       map[uint]*models.Role
	ErrFindAll  error
	ErrFindByID error
}

var _ ports.RoleRepository = (*MockRoleRepository)(nil)

func NewMockRoleRepository(initial []*models.Role) *MockRoleRepository {
	m := &MockRoleRepository{
		Roles: make(map[uint]*models.Role),
	}

	for _, r := range initial {
		m.Roles[r.ID] = r
	}

	return m
}

func (m *MockRoleRepository) FindAll() ([]models.Role, error) {
	if m.ErrFindAll != nil {
		return nil, m.ErrFindAll
	}

	var list []models.Role
	for _, r := range m.Roles {
		list = append(list, *r)
	}

	return list, nil
}

func (m *MockRoleRepository) FindByID(id uint) (*models.Role, error) {
	if m.ErrFindByID != nil {
		return nil, m.ErrFindByID
	}

	if r, ok := m.Roles[id]; ok {
		return r, nil
	}

	return nil, nil
}
