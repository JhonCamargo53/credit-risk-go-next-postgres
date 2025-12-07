package handlers

import (
	"encoding/json"
	"net/http"

	creditStatus "github.com/JhonCamargo53/prueba-tecnica/internal/application/services/credit-status"
)

var creditStatusService *creditStatus.CreditStatusService

func InitCreditStatusHandler(service *creditStatus.CreditStatusService) {
	creditStatusService = service
}

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
	w.Header().Set("Content-Type", "application/json")

	if creditStatusService == nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "creditStatusService no inicializado",
		})
		return
	}

	statuses, err := creditStatusService.GetAllCreditStatuses()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "No se pudieron obtener los estados de crédito",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(statuses)
}
