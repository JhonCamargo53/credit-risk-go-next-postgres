package documentType

import (
	"errors"
	"testing"

	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
)

func TestGetAllDocumentTypes_Exitoso(t *testing.T) {

	mockRepo := NewMockDocumentTypeRepository([]models.DocumentType{
		{
			ID:          1,
			Code:        "CC",
			Description: "Cédula de ciudadanía",
			Status:      true,
		},
		{
			ID:          2,
			Code:        "TI",
			Description: "Tarjeta de identidad",
			Status:      true,
		},
	})

	service := NewDocumentTypeService(mockRepo)

	types, err := service.GetAllDocumentTypes()
	if err != nil {
		t.Fatalf("no se esperaba error: %v", err)
	}

	if len(types) != 2 {
		t.Fatalf("se esperaban 2 tipos de documento, se obtuvo=%d", len(types))
	}

	if types[0].Code != "CC" {
		t.Fatalf("primer tipo debería tener code CC, se obtuvo=%s", types[0].Code)
	}

	if types[1].Description != "Tarjeta de identidad" {
		t.Fatalf("descripción incorrecta, se obtuvo=%s", types[1].Description)
	}
}

func TestGetAllDocumentTypes_ErrorRepositorio(t *testing.T) {

	mockRepo := NewMockDocumentTypeRepository(nil)
	mockRepo.ErrFindAll = errors.New("falló la BD")

	service := NewDocumentTypeService(mockRepo)

	types, err := service.GetAllDocumentTypes()
	if err == nil {
		t.Fatalf("se esperaba error del repositorio")
	}

	if types != nil {
		t.Fatalf("lista debería ser nil cuando ocurre un error")
	}
}
