package engines

import (
	"strings"
	"testing"

	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
)

// Escenario 1: Cliente con riesgo bajo
func TestEvaluateCreditRisk_LowRiskScenario(t *testing.T) {

	customer := models.Customer{
		MonthlyIncome: 8_000_000,
	}

	current := models.CreditRequest{
		Amount:      5_000_000,
		TermMonths:  24,
		ProductType: "Crédito de vivienda",
	}

	assets := []models.CustomerAsset{
		{
			MarketValue: 10_000_000,
			Description: "Vivienda principal",
		},
	}

	otherCredits := []models.CreditRequest{
		{CreditStatusID: 2}, // APROBADO
		{CreditStatusID: 2}, // APROBADO
	}

	score, category, explanation, err := EvaluateCreditRisk(
		customer,
		current,
		otherCredits,
		assets,
	)

	if err != nil {
		t.Fatalf("no se esperaba error, pero se obtuvo: %v", err)
	}

	if score < 0 || score > 100 {
		t.Errorf("el score debe estar entre 0 y 100, obtenido: %.2f", score)
	}

	if category != "LOW" {
		t.Errorf("se esperaba categoría LOW para un cliente sólido, obtenido: %s", category)
	}

	if explanation == "" {
		t.Errorf("se esperaba explicación no vacía")
	}

	if !strings.Contains(explanation, "Puntaje de riesgo") {
		t.Errorf("la explicación debería contener el texto 'Puntaje de riesgo', obtenido: %s", explanation)
	}
}

// Escenario 2: cliente con alto riesgo con categoría HIGH
func TestEvaluateCreditRisk_HighRiskScenario(t *testing.T) {
	customer := models.Customer{
		MonthlyIncome: 2_000_000, // ingreso bajo
	}

	current := models.CreditRequest{
		Amount:      20_000_000, // monto alto
		TermMonths:  12,         // plazo corto con cuota alta
		ProductType: "Crédito de libre inversión",
	}

	// Sin activos como respaldo
	var assets []models.CustomerAsset

	// Varios créditos rechazados en el historial
	otherCredits := []models.CreditRequest{
		{CreditStatusID: 3}, // RECHAZADO
		{CreditStatusID: 3},
		{CreditStatusID: 3},
		{CreditStatusID: 3},
		{CreditStatusID: 3},
	}

	score, category, explanation, err := EvaluateCreditRisk(
		customer,
		current,
		otherCredits,
		assets,
	)

	if err != nil {
		t.Fatalf("no se esperaba error, pero se obtuvo: %v", err)
	}

	if score < 0 || score > 100 {
		t.Errorf("el score debe estar entre 0 y 100, obtenido: %.2f", score)
	}

	if category != "HIGH" {
		t.Errorf("se esperaba categoría HIGH para un cliente riesgoso, obtenido: %s", category)
	}

	if explanation == "" {
		t.Errorf("se esperaba explicación no vacía")
	}

	// En este caso debería mencionar que no hay activos o que el riesgo es alto
	if !strings.Contains(explanation, "No se registran activos") &&
		!strings.Contains(strings.ToLower(explanation), "riesgo") {
		t.Errorf("la explicación debería mencionar la falta de activos o el riesgo, obtenido: %s", explanation)
	}
}

// Escenario 3: validación de que la categoría corresponda al rango de score
func TestEvaluateCreditRisk_CategoryMatchesScore(t *testing.T) {
	customer := models.Customer{
		MonthlyIncome: 4_000_000,
	}

	current := models.CreditRequest{
		Amount:      6_000_000,
		TermMonths:  24,
		ProductType: "Consumo",
	}

	assets := []models.CustomerAsset{
		{
			MarketValue: 2_000_000,
			Description: "Moto",
		},
	}

	otherCredits := []models.CreditRequest{
		{CreditStatusID: 2}, // aprobado
		{CreditStatusID: 3}, // rechazado
	}

	score, category, explanation, err := EvaluateCreditRisk(
		customer,
		current,
		otherCredits,
		assets,
	)

	if err != nil {
		t.Fatalf("no se esperaba error, pero se obtuvo: %v", err)
	}

	if score < 0 || score > 100 {
		t.Errorf("el score debe estar entre 0 y 100, obtenido: %.2f", score)
	}

	switch {
	case score >= 80 && category != "LOW":
		t.Errorf("para score %.2f (>=80) se esperaba categoría LOW, obtenido: %s", score, category)
	case score >= 55 && score < 80 && category != "MEDIUM":
		t.Errorf("para score %.2f (55-79) se esperaba categoría MEDIUM, obtenido: %s", score, category)
	case score < 55 && category != "HIGH":
		t.Errorf("para score %.2f (<55) se esperaba categoría HIGH, obtenido: %s", score, category)
	}

	if explanation == "" {
		t.Errorf("se esperaba explicación no vacía")
	}
}
