package role

import (
	"fmt"

	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/ports"
)

type RoleService struct {
	roleRepo ports.RoleRepository
}

func NewRoleService(roleRepo ports.RoleRepository) *RoleService {
	return &RoleService{
		roleRepo: roleRepo,
	}
}

func (s *RoleService) GetAllRoles() ([]models.Role, error) {
	return s.roleRepo.FindAll()
}

func (s *RoleService) GetRoleByID(id uint) (*models.Role, error) {
	role, err := s.roleRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if role == nil {
		return nil, fmt.Errorf("no existe rol con id %d", id)
	}
	return role, nil
}
