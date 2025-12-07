package creditStatus

import (
	"errors"
	"testing"

	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
)

func TestGetAllCreditStatuses_Exitoso(t *testing.T) {

	mockRepo := NewMockCreditStatusRepository([]models.CreditStatus{
		{ID: 1, Name: "PENDIENTE", Status: true},
		{ID: 2, Name: "APROBADO", Status: true},
		{ID: 3, Name: "RECHAZADO", Status: true},
	})

	service := NewCreditStatusService(mockRepo)

	statuses, err := service.GetAllCreditStatuses()

	if err != nil {
		t.Fatalf("no se esperaba error: %v", err)
	}

	if len(statuses) != 3 {
		t.Fatalf("se esperaban 3 estados de crédito, set obtuvo=%d", len(statuses))
	}

	if statuses[1].Name != "APROBADO" {
		t.Fatalf("se esperaba estado APROBADO, set obtuvo=%s", statuses[1].Name)
	}
}

func TestGetAllCreditStatuses_ErrorRepositorio(t *testing.T) {

	mockRepo := NewMockCreditStatusRepository(nil)
	mockRepo.ErrFindAll = errors.New("falló la BD")

	service := NewCreditStatusService(mockRepo)

	statuses, err := service.GetAllCreditStatuses()

	if err == nil {
		t.Fatalf("se esperaba error del repositorio")
	}

	if statuses != nil {
		t.Fatalf("la lista debería ser nil cuando hay error")
	}
}
