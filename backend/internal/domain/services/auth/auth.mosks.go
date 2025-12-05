package auth

import (
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/ports"
)

type MockUserRepository struct {
	UsersByEmail   map[string]*models.User
	ErrFindByEmail error
}

var _ ports.UserRepository = (*MockUserRepository)(nil)

func NewMockUserRepository(users []*models.User) *MockUserRepository {
	m := &MockUserRepository{
		UsersByEmail: make(map[string]*models.User),
	}
	for _, u := range users {
		m.UsersByEmail[u.Email] = u
	}
	return m
}

func (m *MockUserRepository) FindAllOrderedByCreatedDesc() ([]models.User, error) {
	var res []models.User
	for _, u := range m.UsersByEmail {
		res = append(res, *u)
	}
	return res, nil
}

func (m *MockUserRepository) FindByID(id uint) (*models.User, error) {
	for _, u := range m.UsersByEmail {
		if u.ID == id {
			return u, nil
		}
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
	m.UsersByEmail[user.Email] = user
	return nil
}

func (m *MockUserRepository) Save(user *models.User) error {
	m.UsersByEmail[user.Email] = user
	return nil
}

func (m *MockUserRepository) Delete(id uint) error {
	for email, u := range m.UsersByEmail {
		if u.ID == id {
			delete(m.UsersByEmail, email)
			break
		}
	}
	return nil
}
