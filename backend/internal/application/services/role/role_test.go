package role

import (
	"errors"
	"testing"

	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
)

func TestGetAllRoles_Exitoso(t *testing.T) {

	mockRepo := NewMockRoleRepository([]*models.Role{
		{ID: 1, Name: "ADMIN"},
		{ID: 2, Name: "USER"},
	})

	service := NewRoleService(mockRepo)

	roles, err := service.GetAllRoles()

	if err != nil {
		t.Fatalf("no se esperaba error: %v", err)
	}

	if len(roles) != 2 {
		t.Fatalf("se esperaban 2 roles, se obtuvo=%d", len(roles))
	}

	if roles[0].Name != "ADMIN" {
		t.Fatalf("primer rol incorrecto, se obtuvo=%s", roles[0].Name)
	}
}

func TestGetAllRoles_ErrorRepositorio(t *testing.T) {

	mockRepo := NewMockRoleRepository(nil)
	mockRepo.ErrFindAll = errors.New("falló la BD")

	service := NewRoleService(mockRepo)

	roles, err := service.GetAllRoles()

	if err == nil {
		t.Fatalf("se esperaba error del repositorio, se obtuvo nil")
	}

	if roles != nil {
		t.Fatalf("la lista debería ser nil cuando ocurre un error")
	}
}

func TestGetRoleByID_Exitoso(t *testing.T) {

	mockRepo := NewMockRoleRepository([]*models.Role{
		{ID: 10, Name: "SUPERVISOR"},
	})

	service := NewRoleService(mockRepo)

	role, err := service.GetRoleByID(10)

	if err != nil {
		t.Fatalf("no se esperaba error: %v", err)
	}

	if role == nil || role.ID != 10 {
		t.Fatalf("se esperaba el rol con ID 10")
	}
}

func TestGetRoleByID_NoExiste(t *testing.T) {

	mockRepo := NewMockRoleRepository(nil)

	service := NewRoleService(mockRepo)

	role, err := service.GetRoleByID(99)

	if err == nil {
		t.Fatalf("se esperaba error porque el rol no existe")
	}

	if role != nil {
		t.Fatalf("no se esperaba rol cuando no existe")
	}
}

func TestGetRoleByID_ErrorRepositorio(t *testing.T) {

	mockRepo := NewMockRoleRepository(nil)
	mockRepo.ErrFindByID = errors.New("falló el repositorio")

	service := NewRoleService(mockRepo)

	role, err := service.GetRoleByID(1)

	if err == nil {
		t.Fatalf("se esperaba error propagado del repositorio")
	}

	if role != nil {
		t.Fatalf("role debería ser nil cuando ocurre un error")
	}
}
