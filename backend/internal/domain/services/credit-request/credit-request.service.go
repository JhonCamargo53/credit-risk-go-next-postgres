package creditRequest

import (
	"fmt"

	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/ports"
)

type CreditRequestService struct {
	creditRequestRepo ports.CreditRequestRepository
	customerRepo      ports.CustomerRepository
	creditStatusRepo  ports.CreditStatusRepository
	customerAssetRepo ports.CustomerAssetRepository
	riskEvaluator     ports.RiskEvaluator
}

func NewCreditRequestService(creditRequestRepo ports.CreditRequestRepository, customerRepo ports.CustomerRepository,
	creditStatusRepo ports.CreditStatusRepository, customerAssetRepo ports.CustomerAssetRepository, riskEvaluator ports.RiskEvaluator) *CreditRequestService {
	return &CreditRequestService{
		creditRequestRepo: creditRequestRepo,
		customerRepo:      customerRepo,
		creditStatusRepo:  creditStatusRepo,
		customerAssetRepo: customerAssetRepo,
		riskEvaluator:     riskEvaluator,
	}
}

func (s *CreditRequestService) GetAllCreditRequests(customerId *uint) ([]models.CreditRequest, error) {
	if customerId != nil {
		customer, err := s.customerRepo.FindByID(*customerId)
		if err != nil {
			return nil, err
		}
		if customer == nil {
			return nil, fmt.Errorf("no existe cliente con id %d", *customerId)
		}
	}

	return s.creditRequestRepo.FindAll(customerId)
}

func (s *CreditRequestService) GetCreditRequestByID(id uint) (*models.CreditRequest, error) {
	cr, err := s.creditRequestRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if cr == nil {
		return nil, fmt.Errorf("no existe solicitud de crédito con id %d", id)
	}
	return cr, nil
}

func (s *CreditRequestService) CreateCreditRequest(creditRequest *models.CreditRequest) (*models.CreditRequest, error) {
	// Validar cliente
	customer, err := s.customerRepo.FindByID(creditRequest.CustomerID)
	if err != nil {
		return nil, err
	}
	if customer == nil {
		return nil, fmt.Errorf("no existe cliente con id %d", creditRequest.CustomerID)
	}

	// Validar estado de crédito
	status, err := s.creditStatusRepo.FindByID(creditRequest.CreditStatusID)
	if err != nil {
		return nil, err
	}
	if status == nil {
		return nil, fmt.Errorf("no existe el estado de solicitud con id %d", creditRequest.CreditStatusID)
	}
	_, err = s.creditRequestRepo.Create(creditRequest)
	// Crear solicitud
	if err != nil {
		return nil, err
	}

	// Evaluar riesgo (IA/MOCK)
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
	updatedCreditRequest, err := s.creditRequestRepo.UpdateCreditRiskEvaluation(creditRequest.ID, score, category, explanation)

	if err != nil {
		return nil, err
	}

	return updatedCreditRequest, nil
}

func (s *CreditRequestService) UpdateCreditRequest(id uint, crData *models.CreditRequest) (*models.CreditRequest, error) {
	// Verificar que la solicitud exista
	existing, err := s.GetCreditRequestByID(id)
	if err != nil {
		return nil, err
	}

	// Validar cliente
	customer, err := s.customerRepo.FindByID(crData.CustomerID)
	if err != nil {
		return nil, err
	}
	if customer == nil {
		return nil, fmt.Errorf("no existe cliente con id %d", crData.CustomerID)
	}

	// Validar estado de crédito
	status, err := s.creditStatusRepo.FindByID(crData.CreditStatusID)
	if err != nil {
		return nil, err
	}
	if status == nil {
		return nil, fmt.Errorf("no existe el estado de solicitud con id %d", crData.CreditStatusID)
	}

	// Actualizar
	updated, err := s.creditRequestRepo.Update(id, crData)
	if err != nil {
		return nil, err
	}

	customerData, creditRequest, otherCredits, customerAssets, err := s.creditRequestRepo.FindDataToEvaluateRisk(id)

	if err != nil {
		return nil, err
	}

	// Recalcular riesgo
	score, category, explanation, err := s.riskEvaluator.Evaluate(customerData, *creditRequest, otherCredits, customerAssets)

	if err != nil {
		return nil, err
	}

	//Actualizar riesgo
	updatedCreditRequest, err := s.creditRequestRepo.UpdateCreditRiskEvaluation(creditRequest.ID, score, category, explanation)

	if err != nil {
		return nil, err
	}

	if updated == nil {
		updatedCreditRequest = existing
	}

	return updatedCreditRequest, nil
}

func (s *CreditRequestService) DeleteCreditRequest(id uint) error {

	// Verificar que exista
	if _, err := s.GetCreditRequestByID(id); err != nil {
		return err
	}

	// Verificar activos asociados
	count, err := s.customerAssetRepo.CountByCreditRequestID(id)
	if err != nil {
		return err
	}

	if count > 0 {
		return fmt.Errorf("no se puede eliminar la solicitud de crédito porque tiene activos asociados")
	}

	if err := s.creditRequestRepo.Delete(id); err != nil {
		return err
	}

	return nil
}
