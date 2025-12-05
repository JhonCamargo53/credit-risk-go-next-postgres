package ports

import "github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"

type RiskEvaluator interface {
	Evaluate(customer models.Customer, currentCreditRequest models.CreditRequest,
		otherCredits []models.CreditRequest, assets []models.CustomerAsset) (float64, string, string, error)
}
