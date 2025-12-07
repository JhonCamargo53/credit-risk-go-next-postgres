package customerAsset

import (
	"fmt"

	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/ports"
)

type CustomerAssetService struct {
	customerAssetRepo ports.CustomerAssetRepository
	customerRepo      ports.CustomerRepository
	assetRepo         ports.AssetRepository
	creditRequestRepo ports.CreditRequestRepository
	riskEvaluator     ports.RiskEvaluator
}

func NewCustomerAssetService(customerAssetRepo ports.CustomerAssetRepository, customerRepo ports.CustomerRepository,
	assetRepo ports.AssetRepository, creditRequestRepo ports.CreditRequestRepository,
	riskEvaluator ports.RiskEvaluator) *CustomerAssetService {
	return &CustomerAssetService{
		customerAssetRepo: customerAssetRepo,
		customerRepo:      customerRepo,
		assetRepo:         assetRepo,
		creditRequestRepo: creditRequestRepo,
		riskEvaluator:     riskEvaluator,
	}
}

func (s *CustomerAssetService) GetAllCustomerAssets(creditRequestId *uint) ([]models.CustomerAsset, error) {
	if creditRequestId != nil {
		creditRequest, err := s.creditRequestRepo.FindByID(*creditRequestId)
		if err != nil {
			return nil, err
		}
		if creditRequest == nil {
			return nil, fmt.Errorf("no existe solicitud de crédito %d", *creditRequestId)
		}
	}
	return s.customerAssetRepo.FindAll(creditRequestId)
}

func (s *CustomerAssetService) GetCustomerAssetByID(id uint) (*models.CustomerAsset, error) {
	ca, err := s.customerAssetRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if ca == nil {
		return nil, fmt.Errorf("no existe bien del cliente de id %d", id)
	}
	return ca, nil
}

func (s *CustomerAssetService) CreateCustomerAsset(customerAsset *models.CustomerAsset) (*models.CustomerAsset, error) {
	// Validar cliente
	customer, err := s.customerRepo.FindByID(customerAsset.CustomerID)
	if err != nil {
		return nil, err
	}
	if customer == nil {
		return nil, fmt.Errorf("no existe cliente con id %d", customerAsset.CustomerID)
	}

	// Validar bien (asset)
	asset, err := s.assetRepo.FindByID(customerAsset.AssetID)
	if err != nil {
		return nil, err
	}
	if asset == nil {
		return nil, fmt.Errorf("no existe bien con id %d", customerAsset.AssetID)
	}

	// Validar solicitud de crédito asociada
	creditRequest, err := s.creditRequestRepo.FindByID(customerAsset.CreditRequestID)
	if err != nil {
		return nil, err
	}
	if creditRequest == nil {
		return nil, fmt.Errorf("no se pudo obtener la solicitud de crédito con id %d", customerAsset.CreditRequestID)
	}

	// Crear activo
	if err := s.customerAssetRepo.Create(customerAsset); err != nil {
		return nil, err
	}

	customerData, creditRequest, otherCredits, customerAssets, err := s.creditRequestRepo.FindDataToEvaluateRisk(creditRequest.ID)

	if err != nil {
		return nil, err
	}

	// Recalcular riesgo
	score, category, explanation, err := s.riskEvaluator.Evaluate(customerData, *creditRequest, otherCredits, customerAssets)

	if err != nil {
		return nil, err
	}

	//Actualizar riesgo
	_, err = s.creditRequestRepo.UpdateCreditRiskEvaluation(creditRequest.ID, score, category, explanation)

	if err != nil {
		return nil, err
	}

	return customerAsset, nil
}

func (s *CustomerAssetService) UpdateCustomerAsset(id uint, customerAssetData *models.CustomerAsset) (*models.CustomerAsset, error) {
	// Verificar que el activo exista
	existing, err := s.GetCustomerAssetByID(id)
	if err != nil {
		return nil, err
	}

	// Validar cliente
	customer, err := s.customerRepo.FindByID(customerAssetData.CustomerID)
	if err != nil {
		return nil, err
	}
	if customer == nil {
		return nil, fmt.Errorf("no existe cliente con id %d", customerAssetData.CustomerID)
	}

	// Validar bien (asset)
	asset, err := s.assetRepo.FindByID(customerAssetData.AssetID)
	if err != nil {
		return nil, err
	}
	if asset == nil {
		return nil, fmt.Errorf("no existe bien con id %d", customerAssetData.AssetID)
	}

	// Determinar la solicitud de crédito asociada (puede venir en data o en el registro existente)
	creditRequestID := existing.CreditRequestID
	if customerAssetData.CreditRequestID != 0 {
		creditRequestID = customerAssetData.CreditRequestID
	}

	creditRequest, err := s.creditRequestRepo.FindByID(creditRequestID)
	if err != nil {
		return nil, err
	}
	if creditRequest == nil {
		return nil, fmt.Errorf("no se pudo obtener la solicitud de crédito asociada con id %d", creditRequestID)
	}

	// Actualizar activo
	updated, err := s.customerAssetRepo.Update(id, customerAssetData)
	if err != nil {
		return nil, err
	}

	customerData, creditRequest, otherCredits, customerAssets, err := s.creditRequestRepo.FindDataToEvaluateRisk(creditRequest.ID)

	if err != nil {
		return nil, err
	}

	// Recalcular riesgo
	score, category, explanation, err := s.riskEvaluator.Evaluate(customerData, *creditRequest, otherCredits, customerAssets)

	if err != nil {
		return nil, err
	}

	//Actualizar riesgo
	_, err = s.creditRequestRepo.UpdateCreditRiskEvaluation(creditRequest.ID, score, category, explanation)

	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *CustomerAssetService) DeleteCustomerAsset(id uint) error {
	// Traer el activo
	ca, err := s.GetCustomerAssetByID(id)
	if err != nil {
		return err
	}

	// Traer la solicitud de crédito asociada
	creditRequest, err := s.creditRequestRepo.FindByID(ca.CreditRequestID)
	if err != nil {
		return err
	}
	if creditRequest == nil {
		return fmt.Errorf("no se pudo obtener la solicitud de crédito asociada")
	}

	// Eliminar activo
	if err := s.customerAssetRepo.Delete(id); err != nil {
		return err
	}

	customer, creditRequest, otherCredits, customerAssets, err := s.creditRequestRepo.FindDataToEvaluateRisk(creditRequest.ID)

	if err != nil {
		return err
	}

	// Recalcular riesgo
	score, category, explanation, err := s.riskEvaluator.Evaluate(customer, *creditRequest, otherCredits, customerAssets)

	if err != nil {
		return err
	}

	//Actualizar riesgo
	_, err = s.creditRequestRepo.UpdateCreditRiskEvaluation(creditRequest.ID, score, category, explanation)

	if err != nil {
		return err
	}

	return nil
}
