package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/services"
)

var creditStatusService *services.CreditStatusService

// GetCreditStatusesHandle godoc
// @Summary      Obtener todos los estados de crédito
// @Description  Retorna una lista de todos los estados de crédito disponibles
// @Tags         Credit Statuses
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {array} models.CreditStatus "Lista de estados de crédito"
// @Failure      500 {string} string "Error interno del servidor"
// @Router       /credit-statuses [get]
func GetCreditStatusesHandle(w http.ResponseWriter, r *http.Request) {
	creditStatuses, err := creditStatusService.GetAllCreditStatuses()
	if err != nil {
		http.Error(w, "No se pudieron obtener los estados de creditos", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(creditStatuses)
}
