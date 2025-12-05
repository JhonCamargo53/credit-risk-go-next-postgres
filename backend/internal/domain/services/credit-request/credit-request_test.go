package creditRequest

import (
	"testing"

	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
)

/* GetAllCreditRequests */

func TestGetAllCreditRequests_ClienteNoExiste(t *testing.T) {

	creditRequestRepo := NewMockCreditRequestRepository(nil)
	customerRepo := NewMockCustomerRepository(nil)
	statusRepo := NewMockCreditStatusRepository(nil)
	customerAssetRepo := NewMockCustomerAssetRepository(nil)
	riskEvaluator := &MockRiskEvaluator{}

	service := NewCreditRequestService(creditRequestRepo, customerRepo, statusRepo, customerAssetRepo, riskEvaluator)

	customerID := uint(1)
	creditRequest, err := service.GetAllCreditRequests(&customerID)

	if err == nil {
		t.Fatalf("se esperaba error porque el cliente no existe")
	}
	if creditRequest != nil {
		t.Fatalf("no se esperaba lista de solicitudes cuando el cliente no existe")
	}
}

func TestGetAllCreditRequests_Exitoso(t *testing.T) {
	creditRequestRepo := NewMockCreditRequestRepository([]*models.CreditRequest{
		{ID: 1, CustomerID: 10},
		{ID: 2, CustomerID: 10},
	})
	customerRepo := NewMockCustomerRepository([]*models.Customer{
		{ID: 10, Name: "Juan Test"},
	})
	statusRepo := NewMockCreditStatusRepository(nil)
	customerAssetRepo := NewMockCustomerAssetRepository(nil)
	riskEvaluator := &MockRiskEvaluator{}
	service := NewCreditRequestService(creditRequestRepo, customerRepo, statusRepo, customerAssetRepo, riskEvaluator)

	customerID := uint(10)
	creditRequest, err := service.GetAllCreditRequests(&customerID)

	if err != nil {
		t.Fatalf("no se esperaba error: %v", err)
	}
	if len(creditRequest) != 2 {
		t.Fatalf("se esperaban 2 solicitudes, se obtuvo=%d", len(creditRequest))
	}
}

/* GetCreditRequestByID */

func TestGetCreditRequestByID_NoExiste(t *testing.T) {
	creditRequestRepo := NewMockCreditRequestRepository(nil)
	customerRepo := NewMockCustomerRepository(nil)
	statusRepo := NewMockCreditStatusRepository(nil)
	customerAssetRepo := NewMockCustomerAssetRepository(nil)
	riskEvaluator := &MockRiskEvaluator{}

	service := NewCreditRequestService(creditRequestRepo, customerRepo, statusRepo, customerAssetRepo, riskEvaluator)

	cr, err := service.GetCreditRequestByID(99)

	if err == nil {
		t.Fatalf("se esperaba error porque la solicitud no existe")
	}
	if cr != nil {
		t.Fatalf("no se esperaba solicitud de crédito")
	}
}

func TestGetCreditRequestByID_Exitoso(t *testing.T) {
	creditRequestRepo := NewMockCreditRequestRepository([]*models.CreditRequest{
		{ID: 5, CustomerID: 1},
	})
	customerRepo := NewMockCustomerRepository(nil)
	statusRepo := NewMockCreditStatusRepository(nil)
	customerAssetRepo := NewMockCustomerAssetRepository(nil)
	riskEvaluator := &MockRiskEvaluator{}

	service := NewCreditRequestService(creditRequestRepo, customerRepo, statusRepo, customerAssetRepo, riskEvaluator)

	cr, err := service.GetCreditRequestByID(5)

	if err != nil {
		t.Fatalf("no se esperaba error: %v", err)
	}
	if cr == nil || cr.ID != 5 {
		t.Fatalf("se esperaba solicitud con ID=5")
	}
}

/* CreateCreditRequest */

func TestCreateCreditRequest_ClienteNoExiste(t *testing.T) {
	creditRequestRepo := NewMockCreditRequestRepository(nil)
	customerRepo := NewMockCustomerRepository(nil) // ningún cliente
	statusRepo := NewMockCreditStatusRepository([]*models.CreditStatus{
		{ID: 1, Name: "PENDIENTE"},
	})
	customerAssetRepo := NewMockCustomerAssetRepository(nil)
	riskEvaluator := &MockRiskEvaluator{
		Score:       70,
		Category:    "LOW",
		Explanation: "OK",
	}

	service := NewCreditRequestService(creditRequestRepo, customerRepo, statusRepo, customerAssetRepo, riskEvaluator)

	cr := &models.CreditRequest{
		CustomerID:     99, // no existe
		CreditStatusID: 1,
		Amount:         10_000_000,
	}

	created, err := service.CreateCreditRequest(cr)

	if err == nil {
		t.Fatalf("se esperaba error porque el cliente no existe")
	}
	if created != nil {
		t.Fatalf("no se debería crear la solicitud")
	}
}

func TestCreateCreditRequest_EstadoNoExiste(t *testing.T) {
	creditRequestRepo := NewMockCreditRequestRepository(nil)
	customerRepo := NewMockCustomerRepository([]*models.Customer{
		{ID: 1, Name: "Cliente Test"},
	})
	statusRepo := NewMockCreditStatusRepository(nil) // ningún estado
	customerAssetRepo := NewMockCustomerAssetRepository(nil)
	riskEvaluator := &MockRiskEvaluator{
		Score:       70,
		Category:    "LOW",
		Explanation: "OK",
	}

	service := NewCreditRequestService(creditRequestRepo, customerRepo, statusRepo, customerAssetRepo, riskEvaluator)

	cr := &models.CreditRequest{
		CustomerID:     1,
		CreditStatusID: 99, // no existe
		Amount:         10_000_000,
	}

	created, err := service.CreateCreditRequest(cr)

	if err == nil {
		t.Fatalf("se esperaba error porque el estado de crédito no existe")
	}
	if created != nil {
		t.Fatalf("no se debería crear la solicitud")
	}
}

func TestCreateCreditRequest_Exitoso_RecalculaRiesgo(t *testing.T) {
	// Cliente y estado existen
	customerRepo := NewMockCustomerRepository([]*models.Customer{
		{ID: 1, Name: "Cliente Test", MonthlyIncome: 5_000_000},
	})
	statusRepo := NewMockCreditStatusRepository([]*models.CreditStatus{
		{ID: 1, Name: "PENDIENTE"},
	})

	creditRequestRepo := NewMockCreditRequestRepository(nil)
	customerAssetRepo := NewMockCustomerAssetRepository(nil)

	riskEvaluator := &MockRiskEvaluator{
		Score:       75, // antes 750
		Category:    "LOW",
		Explanation: "Buen perfil",
	}

	service := NewCreditRequestService(creditRequestRepo, customerRepo, statusRepo, customerAssetRepo, riskEvaluator)

	cr := &models.CreditRequest{
		CustomerID:     1,
		CreditStatusID: 1,
		Amount:         20_000_000,
	}

	created, err := service.CreateCreditRequest(cr)
	if err != nil {
		t.Fatalf("no se esperaba error al crear solicitud: %v", err)
	}
	if created == nil {
		t.Fatalf("se esperaba una solicitud creada")
	}
	if created.ID == 0 {
		t.Fatalf("se esperaba que la solicitud tuviera ID asignado por el repo mock")
	}

	// Verificar riesgo
	if !riskEvaluator.Called {
		t.Fatalf("se esperaba que el evaluador de riesgo fuera llamado")
	}
	if !creditRequestRepo.UpdateRiskCalled {
		t.Fatalf("se esperaba que se actualizara la evaluación de riesgo")
	}
	if creditRequestRepo.LastScore != 75 {
		t.Fatalf("score inesperado, se obtuvo=%.2f", creditRequestRepo.LastScore)
	}
	if creditRequestRepo.LastCategory != "LOW" {
		t.Fatalf("categoría inesperada, se obtuvo=%s", creditRequestRepo.LastCategory)
	}
}

/* UpdateCreditRequest */

func TestUpdateCreditRequest_ClienteNoExiste(t *testing.T) {
	// Solicitud existente
	creditRequestRepo := NewMockCreditRequestRepository([]*models.CreditRequest{
		{ID: 10, CustomerID: 1, CreditStatusID: 1},
	})

	// Cliente NO existe en el repo
	customerRepo := NewMockCustomerRepository(nil)

	statusRepo := NewMockCreditStatusRepository([]*models.CreditStatus{
		{ID: 1, Name: "PENDIENTE"},
	})
	customerAssetRepo := NewMockCustomerAssetRepository(nil)

	riskEvaluator := &MockRiskEvaluator{
		Score:       70,
		Category:    "LOW",
		Explanation: "OK",
	}

	service := NewCreditRequestService(creditRequestRepo, customerRepo, statusRepo, customerAssetRepo, riskEvaluator)

	updateData := &models.CreditRequest{
		CustomerID:     99, // no existe
		CreditStatusID: 1,
		Amount:         30_000_000,
	}

	updated, err := service.UpdateCreditRequest(10, updateData)

	if err == nil {
		t.Fatalf("se esperaba error porque el cliente no existe")
	}
	if updated != nil {
		t.Fatalf("no se debería actualizar la solicitud")
	}
}

func TestUpdateCreditRequest_EstadoNoExiste(t *testing.T) {
	creditRequestRepo := NewMockCreditRequestRepository([]*models.CreditRequest{
		{ID: 10, CustomerID: 1, CreditStatusID: 1},
	})
	customerRepo := NewMockCustomerRepository([]*models.Customer{
		{ID: 1, Name: "Cliente Test"},
	})
	statusRepo := NewMockCreditStatusRepository(nil) // ningún estado
	customerAssetRepo := NewMockCustomerAssetRepository(nil)
	riskEvaluator := &MockRiskEvaluator{
		Score:       72,
		Category:    "MEDIUM",
		Explanation: "Cambio de estado",
	}

	service := NewCreditRequestService(creditRequestRepo, customerRepo, statusRepo, customerAssetRepo, riskEvaluator)

	updateData := &models.CreditRequest{
		CustomerID:     1,
		CreditStatusID: 99, // no existe
		Amount:         30_000_000,
	}

	updated, err := service.UpdateCreditRequest(10, updateData)

	if err == nil {
		t.Fatalf("se esperaba error porque el estado no existe")
	}
	if updated != nil {
		t.Fatalf("no se debería actualizar la solicitud")
	}
}

func TestUpdateCreditRequest_Exitoso_RecalculaRiesgo(t *testing.T) {
	creditRequestRepo := NewMockCreditRequestRepository([]*models.CreditRequest{
		{ID: 10, CustomerID: 1, CreditStatusID: 1, Amount: 20_000_000},
	})
	customerRepo := NewMockCustomerRepository([]*models.Customer{
		{ID: 1, Name: "Cliente Test", MonthlyIncome: 5_000_000},
	})
	statusRepo := NewMockCreditStatusRepository([]*models.CreditStatus{
		{ID: 1, Name: "PENDIENTE"},
		{ID: 2, Name: "EN ESTUDIO"},
	})
	customerAssetRepo := NewMockCustomerAssetRepository(nil)

	riskEvaluator := &MockRiskEvaluator{
		Score:       80,
		Category:    "MEDIUM",
		Explanation: "Actualización de condiciones",
	}

	service := NewCreditRequestService(creditRequestRepo, customerRepo, statusRepo, customerAssetRepo, riskEvaluator)

	updateData := &models.CreditRequest{
		CustomerID:     1,
		CreditStatusID: 2,
		Amount:         25_000_000,
	}

	updated, err := service.UpdateCreditRequest(10, updateData)

	if err != nil {
		t.Fatalf("no se esperaba error al actualizar: %v", err)
	}
	if updated == nil {
		t.Fatalf("se esperaba solicitud actualizada")
	}

	if !riskEvaluator.Called {
		t.Fatalf("se esperaba que el evaluador de riesgo fuera llamado")
	}
	if !creditRequestRepo.UpdateRiskCalled {
		t.Fatalf("se esperaba que se actualizara la evaluación de riesgo")
	}
}

/* DeleteCreditRequest */

func TestDeleteCreditRequest_ConActivosAsociados(t *testing.T) {
	creditRequestRepo := NewMockCreditRequestRepository([]*models.CreditRequest{
		{ID: 10, CustomerID: 1, CreditStatusID: 1},
	})
	customerRepo := NewMockCustomerRepository(nil)
	statusRepo := NewMockCreditStatusRepository(nil)

	// Un activo asociado a la solicitud 10
	customerAssetRepo := NewMockCustomerAssetRepository([]*models.CustomerAsset{
		{ID: 1, CreditRequestID: 10},
	})

	riskEvaluator := &MockRiskEvaluator{}

	service := NewCreditRequestService(creditRequestRepo, customerRepo, statusRepo, customerAssetRepo, riskEvaluator)

	err := service.DeleteCreditRequest(10)
	if err == nil {
		t.Fatalf("se esperaba error porque hay activos asociados")
	}
}

func TestDeleteCreditRequest_Exitoso(t *testing.T) {
	creditRequestRepo := NewMockCreditRequestRepository([]*models.CreditRequest{
		{ID: 10, CustomerID: 1, CreditStatusID: 1},
	})
	customerRepo := NewMockCustomerRepository(nil)
	statusRepo := NewMockCreditStatusRepository(nil)

	// Ningún activo asociado
	customerAssetRepo := NewMockCustomerAssetRepository(nil)
	riskEvaluator := &MockRiskEvaluator{}

	service := NewCreditRequestService(creditRequestRepo, customerRepo, statusRepo, customerAssetRepo, riskEvaluator)

	err := service.DeleteCreditRequest(10)
	if err != nil {
		t.Fatalf("no se esperaba error al eliminar solicitud: %v", err)
	}

	if _, ok := creditRequestRepo.Requests[10]; ok {
		t.Fatalf("la solicitud de crédito debería haberse eliminado del repositorio")
	}
}
