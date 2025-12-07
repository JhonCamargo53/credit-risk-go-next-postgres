package asset

import (
	"errors"
	"testing"

	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
)

func TestGetAllAssets_Exitoso(t *testing.T) {

	mockRepo := NewMockAssetRepository([]models.Asset{
		{ID: 1, Name: "Casa"},
		{ID: 2, Name: "Vehículo"},
	})

	service := NewAssetService(mockRepo)

	assets, err := service.GetAllAssets()

	if err != nil {
		t.Fatalf("no se esperaba error: %v", err)
	}

	if len(assets) != 2 {
		t.Fatalf("se esperaban 2 assets, se obtuvo=%d", len(assets))
	}

	if assets[0].Name != "Casa" {
		t.Fatalf("primer asset incorrecto, se obtuvo=%s", assets[0].Name)
	}
}

func TestGetAllAssets_ErrorRepositorio(t *testing.T) {
	mockRepo := NewMockAssetRepository(nil)
	mockRepo.ErrFindAll = errors.New("error en BD")

	service := NewAssetService(mockRepo)

	assets, err := service.GetAllAssets()

	if err == nil {
		t.Fatalf("se esperaba error del repositorio, se obtuvo nil")
	}

	if assets != nil {
		t.Fatalf("assets debería ser nil cuando ocurre un error")
	}
}
