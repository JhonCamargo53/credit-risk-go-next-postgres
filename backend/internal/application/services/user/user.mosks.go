package user

import (
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/ports"
)

/* Mock de UserRepository */

type MockUserRepository struct {
	UsersByID    map[uint]*models.User
	UsersByEmail map[string]*models.User

	ErrFindAllOrdered error
	ErrFindByID       error
	ErrFindByEmail    error
	ErrCreate         error
	ErrSave           error
	ErrDelete         error
}

var _ ports.UserRepository = (*MockUserRepository)(nil)

func NewMockUserRepository(initial []*models.User) *MockUserRepository {
	m := &MockUserRepository{
		UsersByID:    make(map[uint]*models.User),
		UsersByEmail: make(map[string]*models.User),
	}

	for _, u := range initial {
		m.UsersByID[u.ID] = u
		m.UsersByEmail[u.Email] = u
	}

	return m
}

func (m *MockUserRepository) FindAllOrderedByCreatedDesc() ([]models.User, error) {
	if m.ErrFindAllOrdered != nil {
		return nil, m.ErrFindAllOrdered
	}

	var res []models.User
	for _, u := range m.UsersByID {
		res = append(res, *u)
	}
	return res, nil
}

func (m *MockUserRepository) FindByID(id uint) (*models.User, error) {
	if m.ErrFindByID != nil {
		return nil, m.ErrFindByID
	}

	if u, ok := m.UsersByID[id]; ok {
		return u, nil
	}
	return nil, nil
}

func (m *MockUserRepository) FindByEmail(email string) (*models.User, error) {
	if m.ErrFindByEmail != nil {
		return nil, m.ErrFindByEmail
	}

	if u, ok := m.UsersByEmail[email]; ok {
		return u, nil
	}
	return nil, nil
}

func (m *MockUserRepository) Create(user *models.User) error {
	if m.ErrCreate != nil {
		return m.ErrCreate
	}

	m.UsersByID[user.ID] = user
	m.UsersByEmail[user.Email] = user
	return nil
}

func (m *MockUserRepository) Save(user *models.User) error {
	if m.ErrSave != nil {
		return m.ErrSave
	}

	m.UsersByID[user.ID] = user
	m.UsersByEmail[user.Email] = user
	return nil
}

func (m *MockUserRepository) Delete(id uint) error {
	if m.ErrDelete != nil {
		return m.ErrDelete
	}

	if u, ok := m.UsersByID[id]; ok {
		delete(m.UsersByEmail, u.Email)
	}
	delete(m.UsersByID, id)
	return nil
}

/* Mock de RoleRepository */
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
	var res []models.Role
	for _, r := range m.Roles {
		res = append(res, *r)
	}
	return res, nil
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
