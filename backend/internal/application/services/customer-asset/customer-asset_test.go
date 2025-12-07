package customerAsset

import (
	"testing"

	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
)

func TestCreateCustomerAsset_ClienteNoExiste(t *testing.T) {
	assetRepo := NewMockAssetRepository(nil)

	customerRepo := NewMockCustomerRepository(nil) // ningún cliente
	creditRequestRepo := NewMockCreditRequestRepository(nil)
	customerAssetRepo := NewMockCustomerAssetRepository(nil)
	riskEvaluator := &MockRiskEvaluator{
		Score:       700,
		Category:    "LOW",
		Explanation: "ok",
	}

	service := NewCustomerAssetService(customerAssetRepo, customerRepo, assetRepo, creditRequestRepo, riskEvaluator)

	newAsset := &models.CustomerAsset{
		CustomerID:      1, // no existe
		AssetID:         1,
		CreditRequestID: 1,
		MarketValue:     10_000_000,
		Description:     "Casa",
	}

	_, err := service.CreateCustomerAsset(newAsset)
	if err == nil {
		t.Fatalf("se esperaba error porque el cliente no existe")
	}

}

func TestCreateCustomerAsset_AssetNoExiste(t *testing.T) {
	// Cliente sí existe
	customerRepo := NewMockCustomerRepository([]*models.Customer{
		{ID: 1, Name: "Cliente Test"},
	})

	assetRepo := NewMockAssetRepository(nil) // ningún asset

	// CreditRequest sí existe
	creditRequestRepo := NewMockCreditRequestRepository([]*models.CreditRequest{
		{ID: 1, CustomerID: 1},
	})

	customerAssetRepo := NewMockCustomerAssetRepository(nil)
	riskEvaluator := &MockRiskEvaluator{
		Score:       700,
		Category:    "LOW",
		Explanation: "ok",
	}

	service := NewCustomerAssetService(customerAssetRepo, customerRepo, assetRepo, creditRequestRepo, riskEvaluator)

	newAsset := &models.CustomerAsset{
		CustomerID:      1,
		AssetID:         999, // no existe
		CreditRequestID: 1,
		MarketValue:     10_000_000,
		Description:     "Casa",
	}

	created, err := service.CreateCustomerAsset(newAsset)
	if err == nil {
		t.Fatalf("se esperaba error porque el asset no existe")
	}
	if created != nil {
		t.Fatalf("no se debería crear el asset cuando el bien no existe")
	}
}

func TestCreateCustomerAsset_Exitoso_RecalculaRiesgo(t *testing.T) {
	// Cliente existe
	customerRepo := NewMockCustomerRepository([]*models.Customer{
		{ID: 1, Name: "Cliente Test", MonthlyIncome: 5_000_000},
	})

	// Asset existe
	assetRepo := NewMockAssetRepository([]*models.Asset{
		{ID: 1, Name: "Casa"},
	})

	// CreditRequest existe
	creditRequestRepo := NewMockCreditRequestRepository([]*models.CreditRequest{
		{ID: 10, CustomerID: 1, Amount: 20_000_000},
	})

	customerAssetRepo := NewMockCustomerAssetRepository(nil)

	riskEvaluator := &MockRiskEvaluator{
		Score:       98,
		Category:    "LOW",
		Explanation: "Cliente con buen ingreso y buenos activos",
	}

	service := NewCustomerAssetService(customerAssetRepo, customerRepo, assetRepo, creditRequestRepo, riskEvaluator)

	newAsset := &models.CustomerAsset{
		CustomerID:      1,
		AssetID:         1,
		CreditRequestID: 10,
		MarketValue:     50_000_000,
		Description:     "Casa principal",
	}

	created, err := service.CreateCustomerAsset(newAsset)
	if err != nil {
		t.Fatalf("no se esperaba error al crear CustomerAsset válido: %v", err)
	}
	if created == nil {
		t.Fatalf("se esperaba un CustomerAsset creado")
	}
	if created.ID == 0 {
		t.Fatalf("se esperaba que el CustomerAsset tuviera un ID asignado")
	}

	// Verificar que se haya llamado el evaluador de riesgo
	if !riskEvaluator.Called {
		t.Fatalf("se esperaba que el evaluador de riesgo fuera llamado")
	}
	if riskEvaluator.CalledWithCR != 10 {
		t.Fatalf("se esperaba que el riesgo se evaluara para la CreditRequest 10, se obtuvo= %d", riskEvaluator.CalledWithCR)
	}

	// Verificar que se haya actualizado la evaluación de riesgo en el repositorio
	if !creditRequestRepo.UpdateRiskCalled {
		t.Fatalf("se esperaba que se actualizara la evaluación de riesgo en el CreditRequest")
	}
	if creditRequestRepo.LastScore != 98 {
		t.Fatalf("score inesperado, se obtuvo=%.2f", creditRequestRepo.LastScore)
	}
	if creditRequestRepo.LastCategory != "LOW" {
		t.Fatalf("categoría inesperada, se obtuvo=%s", creditRequestRepo.LastCategory)
	}
}

func TestGetAllCustomerAssets_CreditRequestNoExiste(t *testing.T) {
	customerAssetRepo := NewMockCustomerAssetRepository(nil)
	customerRepo := NewMockCustomerRepository(nil)
	assetRepo := NewMockAssetRepository(nil)
	creditRequestRepo := NewMockCreditRequestRepository(nil) // ninguna credit request
	riskEvaluator := &MockRiskEvaluator{}

	service := NewCustomerAssetService(customerAssetRepo, customerRepo, assetRepo, creditRequestRepo, riskEvaluator)

	creditRequestID := uint(99)

	assets, err := service.GetAllCustomerAssets(&creditRequestID)
	if err == nil {
		t.Fatalf("se esperaba error porque la solicitud de crédito no existe")
	}
	if assets != nil {
		t.Fatalf("no se esperaba lista de assets cuando la credit request no existe")
	}
}

func TestDeleteCustomerAsset_Exitoso_RecalculaRiesgo(t *testing.T) {
	// Asset existente asociado a CreditRequest 10
	existingAsset := &models.CustomerAsset{
		ID:              1,
		CustomerID:      1,
		AssetID:         1,
		CreditRequestID: 10,
		MarketValue:     10_000_000,
		Description:     "Casa",
	}

	customerAssetRepo := NewMockCustomerAssetRepository([]*models.CustomerAsset{existingAsset})

	// CreditRequest 10 en el sistema
	creditRequestRepo := NewMockCreditRequestRepository([]*models.CreditRequest{
		{ID: 10, CustomerID: 1, Amount: 20_000_000},
	})

	customerRepo := NewMockCustomerRepository([]*models.Customer{
		{ID: 1, Name: "Juan", MonthlyIncome: 5_000_000},
	})

	assetRepo := NewMockAssetRepository([]*models.Asset{
		{ID: 1, Name: "Casa"},
	})

	riskEvaluator := &MockRiskEvaluator{
		Score:       680,
		Category:    "MEDIUM",
		Explanation: "Menos respaldo en activos tras eliminación",
	}

	service := NewCustomerAssetService(customerAssetRepo, customerRepo, assetRepo, creditRequestRepo, riskEvaluator)

	err := service.DeleteCustomerAsset(1)
	if err != nil {
		t.Fatalf("no se esperaba error al eliminar CustomerAsset: %v", err)
	}

	// Asset debe estar eliminado del repo
	if _, ok := customerAssetRepo.Assets[1]; ok {
		t.Fatalf("el CustomerAsset debería haberse eliminado del repositorio")
	}

	// Riesgo debe recalcularse
	if !riskEvaluator.Called {
		t.Fatalf("se esperaba que el evaluador de riesgo fuera llamado al eliminar un asset")
	}
	if !creditRequestRepo.UpdateRiskCalled {
		t.Fatalf("se esperaba que se actualizara la evaluación de riesgo tras eliminar el asset")
	}
}
