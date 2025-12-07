package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	creditRequest "github.com/JhonCamargo53/prueba-tecnica/internal/application/services/credit-request"
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/gorilla/mux"
)

var creditRequestService *creditRequest.CreditRequestService

// InitCreditRequestHandler inyecta el servicio de dominio desde main.go
func InitCreditRequestHandler(s *creditRequest.CreditRequestService) {
	creditRequestService = s
}

// CreateCreditRequestRequest representa el cuerpo de la solicitud para crear una solicitud de crédito
// @Description Datos para crear una nueva solicitud de crédito
type CreateCreditRequestRequest struct {
	Amount         float64 `json:"amount" example:"10000000"`
	TermMonths     int     `json:"termMonths" example:"24"`
	CustomerID     uint    `json:"customerId" example:"1"`
	ProductType    string  `json:"productType" example:"Préstamo Personal"`
	CreditStatusID uint    `json:"creditStatusId" example:"1"`
}

// UpdateCreditRequestRequest representa el cuerpo de la solicitud para actualizar una solicitud de crédito
// @Description Datos para actualizar una solicitud de crédito existente
type UpdateCreditRequestRequest struct {
	Amount         float64 `json:"amount" example:"15000000"`
	TermMonths     int     `json:"termMonths" example:"36"`
	CustomerID     uint    `json:"customerId" example:"1"`
	ProductType    string  `json:"productType" example:"Hipoteca"`
	CreditStatusID uint    `json:"creditStatusId" example:"2"`
}

// GetCreditRequestsHandle godoc
// @Summary      Obtener todas las solicitudes de crédito
// @Description  Retorna una lista de todas las solicitudes de crédito, opcionalmente filtradas por cliente
// @Tags         Credit Requests
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        customerId query int false "ID del cliente para filtrar"
// @Success      200 {array} models.CreditRequest "Lista de solicitudes de crédito"
// @Failure      400 {string} string "customerId inválido"
// @Failure      500 {string} string "Error interno del servidor"
// @Router       /credit-requests [get]
func GetCreditRequestsHandle(w http.ResponseWriter, r *http.Request) {

	customerIdStr := r.URL.Query().Get("customerId")
	var customerId *uint = nil

	if customerIdStr != "" {
		parsedId, err := strconv.ParseUint(customerIdStr, 10, 32)
		if err != nil {
			http.Error(w, "customerId inválido", http.StatusBadRequest)
			return
		}
		temp := uint(parsedId)
		customerId = &temp
	}

	creditRequests, err := creditRequestService.GetAllCreditRequests(customerId)
	if err != nil {
		http.Error(w, "No se pudieron obtener las solicitudes de crédito", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(creditRequests)
}

// GetCreditRequestHandle godoc
// @Summary      Obtener una solicitud de crédito por ID
// @Description  Retorna los detalles de una solicitud de crédito específica
// @Tags         Credit Requests
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID de la solicitud de crédito"
// @Success      200 {object} models.CreditRequest "Solicitud de crédito encontrada"
// @Failure      400 {string} string "ID inválido"
// @Failure      404 {string} string "Solicitud no encontrada"
// @Failure      500 {string} string "Error interno del servidor"
// @Router       /credit-requests/{id} [get]
func GetCreditRequestHandle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	creditRequestId := params["id"]
	id, err := strconv.Atoi(creditRequestId)

	if err != nil || id <= 0 {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	creditRequest, err := creditRequestService.GetCreditRequestByID(uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "no existe") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Error al modificar usuario: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(creditRequest)
}

// PostCreditRequestHandle godoc
// @Summary      Crear una nueva solicitud de crédito
// @Description  Crea una nueva solicitud de crédito en el sistema
// @Tags         Credit Requests
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body CreateCreditRequestRequest true "Datos de la solicitud de crédito"
// @Success      200 {object} models.CreditRequest "Solicitud de crédito creada exitosamente"
// @Failure      400 {string} string "Solicitud inválida"
// @Failure      404 {string} string "Cliente o estado de crédito no encontrado"
// @Failure      500 {string} string "Error interno del servidor"
// @Router       /credit-requests [post]
func PostCreditRequestHandle(w http.ResponseWriter, r *http.Request) {

	var creditRequestData struct {
		Amount         float64 `json:"amount"`
		TermMonths     int     ` json:"termMonths"`
		CustomerID     uint    `json:"customerId"`
		ProductType    string  `json:"productType"`
		CreditStatusID uint    `json:"creditStatusId"`
	}

	err := json.NewDecoder(r.Body).Decode(&creditRequestData)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if creditRequestData.Amount == 0 {
		http.Error(w, "El campo 'Cantidad' es obligatorio.", http.StatusBadRequest)
		return
	}

	if creditRequestData.TermMonths == 0 {
		http.Error(w, "El campo 'Número de meses' es obligatorio.", http.StatusBadRequest)
		return
	}

	if creditRequestData.CustomerID == 0 {
		http.Error(w, "El campo 'Id del solicitante' es obligatorio.", http.StatusBadRequest)
		return
	}

	if creditRequestData.CreditStatusID == 0 {
		http.Error(w, "El campo 'Estado del crédito' es obligatorio.", http.StatusBadRequest)
		return
	}

	creditRequest := models.CreditRequest{
		Amount:         creditRequestData.Amount,
		TermMonths:     creditRequestData.TermMonths,
		CustomerID:     creditRequestData.CustomerID,
		ProductType:    creditRequestData.ProductType,
		CreditStatusID: creditRequestData.CreditStatusID,
	}

	createdCreditRequest, err := creditRequestService.CreateCreditRequest(&creditRequest)

	if err != nil {
		if strings.Contains(err.Error(), "no existe") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Error al crear solicitud de credito: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdCreditRequest)

}

// UpdateCreditRequestHandle godoc
// @Summary      Actualizar una solicitud de crédito
// @Description  Actualiza los datos de una solicitud de crédito existente
// @Tags         Credit Requests
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID de la solicitud de crédito"
// @Param        request body UpdateCreditRequestRequest true "Datos actualizados de la solicitud"
// @Success      200 {object} models.CreditRequest "Solicitud actualizada exitosamente"
// @Failure      400 {string} string "Solicitud inválida"
// @Failure      404 {string} string "Solicitud no encontrada"
// @Failure      500 {string} string "Error interno del servidor"
// @Router       /credit-requests/{id} [put]
func UpdateCreditRequestHandle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	creditRequestId := params["id"]
	id, err := strconv.Atoi(creditRequestId)
	if err != nil || id <= 0 {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var creditRequestData struct {
		Amount         float64 `json:"amount"`
		TermMonths     int     `json:"termMonths"`
		CustomerID     uint    `json:"customerId"`
		ProductType    string  `json:"productType"`
		CreditStatusID uint    `json:"creditStatusId"`
	}

	err = json.NewDecoder(r.Body).Decode(&creditRequestData)

	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if creditRequestData.Amount == 0 {
		http.Error(w, "El campo 'Cantidad' es obligatorio.", http.StatusBadRequest)
		return
	}

	if creditRequestData.TermMonths == 0 {
		http.Error(w, "El campo 'Número de meses' es obligatorio.", http.StatusBadRequest)
		return
	}

	if creditRequestData.CustomerID == 0 {
		http.Error(w, "El campo 'Id del solicitante' es obligatorio.", http.StatusBadRequest)
		return
	}

	if creditRequestData.CreditStatusID == 0 {
		http.Error(w, "El campo 'Estado del crédito' es obligatorio.", http.StatusBadRequest)
		return
	}

	creditRequest := models.CreditRequest{
		Amount:         creditRequestData.Amount,
		TermMonths:     creditRequestData.TermMonths,
		CustomerID:     creditRequestData.CustomerID,
		ProductType:    creditRequestData.ProductType,
		CreditStatusID: creditRequestData.CreditStatusID,
	}

	updated, err := creditRequestService.UpdateCreditRequest(uint(id), &creditRequest)
	if err != nil {
		if strings.Contains(err.Error(), "no existe") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Error al actualizar solicitud de crédito: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updated)
}

// DeleteCreditRequestHandle godoc
// @Summary      Eliminar una solicitud de crédito
// @Description  Elimina una solicitud de crédito del sistema
// @Tags         Credit Requests
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID de la solicitud de crédito"
// @Success      204 "Solicitud eliminada exitosamente"
// @Failure      400 {string} string "ID inválido"
// @Failure      404 {string} string "Solicitud no encontrada"
// @Failure      500 {string} string "Error interno del servidor"
// @Router       /credit-requests/{id} [delete]
func DeleteCreditRequestHandle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = creditRequestService.DeleteCreditRequest(uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "no existe") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Error al eliminar solicitud de credito: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
