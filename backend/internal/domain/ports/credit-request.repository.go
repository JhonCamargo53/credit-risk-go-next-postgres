package ports

import "github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"

type CreditRequestRepository interface {
	FindAll(customerID *uint) ([]models.CreditRequest, error)
	FindByID(id uint) (*models.CreditRequest, error)
	HasRequestsByCustomerID(customerID uint) (bool, error)
	Create(creditRequest *models.CreditRequest) (*models.CreditRequest, error)
	Update(id uint, creditRequest *models.CreditRequest) (*models.CreditRequest, error)
	Delete(id uint) error
	UpdateCreditRiskEvaluation(id uint, score float64, category string, explanation string) (*models.CreditRequest, error)
	FindDataToEvaluateRisk(id uint) (models.Customer, *models.CreditRequest, []models.CreditRequest, []models.CustomerAsset, error)
}
