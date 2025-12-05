package adapters

import (
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/ports"
	engines "github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/ai/credit-risk/engines"
)

type RiskEvaluatorAdapter struct {
}

func NewRiskEvaluatorAdapter() ports.RiskEvaluator {
	return &RiskEvaluatorAdapter{}
}

func (a *RiskEvaluatorAdapter) Evaluate(customer models.Customer, currentCreditRequest models.CreditRequest,
	otherCredits []models.CreditRequest, assets []models.CustomerAsset) (float64, string, string, error) {

	// Realizar el analisis de riesgo
	score, category, explanation, err := engines.EvaluateCreditRisk(
		customer,
		currentCreditRequest,
		otherCredits,
		assets,
	)

	if err != nil {
		return 0, "", "", err
	}

	return score, category, explanation, nil
}
