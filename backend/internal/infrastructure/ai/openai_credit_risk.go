package risk

import (
	"fmt"
	"strings"

	"github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/helper"

	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"gorm.io/gorm"
)

/*

EvaluateCreditRisk recibe un CreditRequest (con sólo el ID llenado),
carga la info necesaria, calcula el riesgo, actualiza el registro
y devuelve la explicación en lenguaje natural.

*/

func EvaluateCreditRisk(db *gorm.DB, cr *models.CreditRequest) (string, error) {
	// Cargar la info relacionada
	if err := db.
		Preload("Customer").
		Preload("Customer.CreditRequests").
		First(cr, cr.ID).Error; err != nil {
		return "", err
	}

	// Activos asociados SÓLO a este crédito
	var assets []models.CustomerAsset
	if err := db.
		Where("credit_request_id = ? AND status = ?", cr.ID, true).
		Find(&assets).Error; err != nil {
		return "", err
	}

	//Calcular puntaje
	score, reasons, improvements := calculateScore(db, cr, assets, cr.ID)

	//Determinar categoría y recomendación
	categoryEN, categoryES := riskCategory(score)
	recommendation := recommendationFromScore(score)

	//Construir explicación en lenguaje natural
	explanation := buildExplanation(score, categoryES, recommendation, reasons, improvements)

	//Guardar en la solicitud de crédito
	cr.RiskScore = score
	cr.RiskCategory = categoryEN
	cr.RiskExplanation = explanation

	if err := db.Save(cr).Error; err != nil {
		return "", err
	}

	return explanation, nil
}

func calculateScore(db *gorm.DB, cr *models.CreditRequest, assets []models.CustomerAsset, currentCreditRequestId uint) (float64, []string, []string) {
	score := 50.0 // base
	var reasons []string
	var improvements []string

	customer := cr.Customer

	// Relación cuota / ingreso
	amount := cr.Amount
	term := float64(cr.TermMonths)
	income := customer.MonthlyIncome

	if amount <= 0 || term <= 0 || income <= 0 {
		score -= 15
		reasons = append(reasons,
			"No fue posible calcular adecuadamente la relación cuota/ingreso (monto, plazo o ingreso inválidos).")
		improvements = append(improvements,
			"Registrar un ingreso mensual realista y/o ajustar el monto y el plazo del crédito.")
	} else {
		quota := amount / term
		ratio := quota / income // cuota / ingreso

		ratioPct := ratio * 100
		reasons = append(reasons,
			fmt.Sprintf(
				"La cuota mensual estimada es de %s, lo que corresponde al %.1f%% del ingreso mensual del cliente (%s).",
				helper.FormatCOP(quota),
				ratioPct,
				helper.FormatCOP(income),
			),
		)

		switch {
		case ratio <= 0.20:
			score += 25
		case ratio <= 0.30:
			score += 20
		case ratio <= 0.40:
			score += 10
		case ratio <= 0.50:
		default:
			score -= 15
			improvements = append(improvements,
				"Se recomienda reducir el monto solicitado o ampliar el plazo para que la cuota no supere el 30% del ingreso mensual.")
		}
	}

	//Activos asociados a ESTE crédito
	totalAssetsValue := 0.0
	hasViviendaAsset := false

	for _, a := range assets {
		totalAssetsValue += a.MarketValue

		desc := strings.ToLower(a.Description)
		if strings.Contains(desc, "vivienda") ||
			strings.Contains(desc, "casa") ||
			strings.Contains(desc, "apartamento") {
			hasViviendaAsset = true
		}
	}

	if totalAssetsValue <= 0 {
		reasons = append(reasons,
			"No se registran activos con valor asociado específicamente a esta solicitud de crédito.")
		improvements = append(improvements,
			"Incluir activos con valor de mercado (por ejemplo vivienda o vehículo) como respaldo del crédito.")
	} else {

		ratioAssets := totalAssetsValue / amount
		reasons = append(reasons,
			fmt.Sprintf(
				"El valor total de los activos registrados para este crédito es de %s, lo que equivale a %.1f veces el monto solicitado.",
				helper.FormatCOP(totalAssetsValue),
				ratioAssets,
			),
		)

		switch {
		case ratioAssets >= 2.0:
			score += 20
		case ratioAssets >= 1.0:
			score += 15
		case ratioAssets >= 0.5:
			score += 8
		default:
			// pocos activos frente al monto
			improvements = append(improvements,
				"El valor de los activos es bajo frente al monto solicitado; se recomienda aumentar garantías.")
		}

		if hasViviendaAsset {
			score += 8
			reasons = append(reasons,
				"Se registra al menos un activo tipo vivienda como respaldo, lo cual mejora el perfil de riesgo.")
		}
	}

	//Historial de créditos del cliente (Sin añadir el actual)
	var otherCredits []models.CreditRequest

	if err := db.
		Where("customer_id = ? AND id <> ?", customer.ID, currentCreditRequestId).
		Find(&otherCredits).Error; err != nil {
	}

	totalCredits := len(otherCredits)
	approvedCount := 0
	rejectedCount := 0

	for _, other := range otherCredits {
		switch other.CreditStatusID {
		case 2: // APROBADO
			approvedCount++
		case 3: // RECHAZADO
			rejectedCount++
		}
	}

	// Número total de créditos solicitados
	switch {
	case totalCredits == 0:
		score += 5
		reasons = append(reasons, "Es la primera solicitud de crédito registrada para este cliente.")
	case totalCredits <= 3:
		reasons = append(reasons,
			fmt.Sprintf("El cliente ha realizado %d solicitudes de crédito en el sistema.", totalCredits))
	case totalCredits <= 5:
		score -= 5
		reasons = append(reasons,
			fmt.Sprintf("El cliente ha realizado %d solicitudes de crédito; esto incrementa ligeramente el riesgo.", totalCredits))
	default:
		score -= 10
		reasons = append(reasons,
			fmt.Sprintf("El cliente ha realizado %d solicitudes de crédito; un número alto de solicitudes eleva el riesgo.", totalCredits))
		improvements = append(improvements,
			"Reducir la cantidad de solicitudes de crédito simultáneas o recientes.")
	}

	// Créditos aprobados / culminados
	if approvedCount > 0 {
		score += 5
		reasons = append(reasons,
			fmt.Sprintf("Historial positivo: %d crédito(s) aprobado(s) en el sistema.", approvedCount))
		if approvedCount >= 3 {
			score += 5
			reasons = append(reasons,
				"El cliente tiene varios créditos aprobados, lo que indica buen comportamiento histórico.")
		}
	} else {
		improvements = append(improvements,
			"No se encuentran créditos aprobados previos; mantener un buen comportamiento en este crédito ayudará al historial.")
	}

	// Créditos rechazados
	if rejectedCount > 0 {
		if rejectedCount == 1 {
			score -= 8
		} else {
			score -= 12
		}
		reasons = append(reasons,
			fmt.Sprintf("Se encuentran %d crédito(s) rechazado(s) previamente, lo que disminuye el puntaje de riesgo.", rejectedCount))
		improvements = append(improvements,
			"Revisar las causas de rechazo de solicitudes anteriores y corregirlas antes de solicitar nuevos créditos.")
	}

	//Tipo de producto (vivienda / libre inversión)
	productType := strings.ToUpper(strings.TrimSpace(cr.ProductType))

	if strings.Contains(productType, "VIVIENDA") || strings.Contains(productType, "HIPOTEC") {
		score += 10
		reasons = append(reasons,
			"El producto corresponde a crédito de vivienda/hipotecario, que suele estar respaldado en activos reales.")
	}

	if strings.Contains(productType, "LIBRE") || strings.Contains(productType, "CONSUMO") {
		score -= 10
		reasons = append(reasons,
			"El producto es de libre inversión/consumo, usualmente más riesgoso por no estar asociado a un activo específico.")
		improvements = append(improvements,
			"Para montos altos se recomienda preferir créditos respaldados en vivienda u otros activos.")
	}

	if score < 0 {
		score = 0
	}
	if score > 100 {
		score = 100
	}

	return score, reasons, improvements
}

func riskCategory(score float64) (categoryEN, categoryES string) {
	switch {
	case score >= 80:
		return "LOW", "Bajo"
	case score >= 55:
		return "MEDIUM", "Medio"
	default:
		return "HIGH", "Alto"
	}
}

func recommendationFromScore(score float64) string {
	switch {
	case score >= 75:
		return "APROBAR"
	case score >= 55:
		return "DEJAR EN ESTUDIO / APROBAR CON CONDICIONES"
	default:
		return "NO APROBAR"
	}
}

func buildExplanation(
	score float64,
	categoryES string,
	recommendation string,
	reasons []string,
	improvements []string,
) string {
	var b strings.Builder

	fmt.Fprintf(&b, "- {Puntaje de riesgo:} %.1f/100\n", score)
	fmt.Fprintf(&b, "- {Rango de riesgo:} %s\n", categoryES)
	fmt.Fprintf(&b, "- {Recomendación del motor:} %s.\n\n", recommendation)

	for _, r := range reasons {
		b.WriteString("- " + r + "\n")
	}

	if len(improvements) > 0 {
		b.WriteString("\nPosibles mejoras para futuros análisis:\n")
		for _, m := range improvements {
			b.WriteString("- " + m + "\n")
		}
	}

	return b.String()
}
