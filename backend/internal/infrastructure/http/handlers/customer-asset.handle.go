package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	customerAsset "github.com/JhonCamargo53/prueba-tecnica/internal/domain/services/customer-asset"
	"github.com/gorilla/mux"
)

var customerAssetService *customerAsset.CustomerAssetService

func InitCustomerAssetHandler(s *customerAsset.CustomerAssetService) {
	customerAssetService = s
}

// CreateCustomerAssetRequest representa el cuerpo de la solicitud para crear un bien del cliente
// @Description Datos para crear un nuevo bien del cliente
type CreateCustomerAssetRequest struct {
	CustomerID      uint    `json:"customerId" example:"1"`
	AssetID         uint    `json:"assetId" example:"1"`
	CreditRequestID uint    `json:"creditRequestId" example:"1"`
	MarketValue     float64 `json:"marketValue" example:"50000000"`
	Description     string  `json:"description" example:"Vehículo Toyota Corolla 2020"`
}

// UpdateCustomerAssetRequest representa el cuerpo de la solicitud para actualizar un bien del cliente
// @Description Datos para actualizar un bien del cliente existente
type UpdateCustomerAssetRequest struct {
	CustomerID      uint    `json:"customerId" example:"1"`
	AssetID         uint    `json:"assetId" example:"2"`
	CreditRequestID uint    `json:"creditRequestId" example:"1"`
	MarketValue     float64 `json:"marketValue" example:"55000000"`
	Description     string  `json:"description" example:"Vehículo Toyota Corolla 2021"`
}

// GetCustomerAssetsHandle godoc
// @Summary      Obtener todos los bienes de clientes
// @Description  Retorna una lista de todos los bienes de clientes, opcionalmente filtrados por solicitud de crédito
// @Tags         Customer Assets
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        creditRequestId query int false "ID de la solicitud de crédito para filtrar"
// @Success      200 {array} models.CustomerAsset "Lista de bienes de clientes"
// @Failure      400 {string} string "creditRequestId inválido"
// @Failure      404 {string} string "Solicitud de crédito no encontrada"
// @Failure      500 {string} string "Error interno del servidor"
// @Router       /customer-assets [get]
func GetCustomerAssetsHandle(w http.ResponseWriter, r *http.Request) {

	creditRequestIDStr := r.URL.Query().Get("creditRequestId")
	var creditRequestId *uint = nil

	if creditRequestIDStr != "" {
		parsedID, err := strconv.ParseUint(creditRequestIDStr, 10, 32)
		if err != nil {
			http.Error(w, "creditRequestId inválido", http.StatusBadRequest)
			return
		}
		temp := uint(parsedID)
		creditRequestId = &temp
	}

	customerAssets, err := customerAssetService.GetAllCustomerAssets(creditRequestId)

	if err != nil {
		if strings.Contains(err.Error(), "no existe") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Error al obtener bienes del cliente: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customerAssets)
}

// PostCustomerAssetHandle godoc
// @Summary      Crear un nuevo bien del cliente
// @Description  Crea un nuevo bien asociado a un cliente y solicitud de crédito
// @Tags         Customer Assets
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body CreateCustomerAssetRequest true "Datos del bien del cliente"
// @Success      201 {object} models.CustomerAsset "Bien creado exitosamente"
// @Failure      400 {string} string "Solicitud inválida"
// @Failure      404 {string} string "Cliente, activo o solicitud no encontrada"
// @Failure      500 {string} string "Error interno del servidor"
// @Router       /customer-assets [post]
func PostCustomerAssetHandle(w http.ResponseWriter, r *http.Request) {

	var customerAssetData struct {
		CustomerID      uint    `json:"customerId"`
		AssetID         uint    `json:"assetId"`
		CreditRequestID uint    `json:"creditRequestId"`
		MarketValue     float64 `json:"marketValue"`
		Description     string  `json:"description"`
	}

	err := json.NewDecoder(r.Body).Decode(&customerAssetData)
	fmt.Println(&customerAssetData)

	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if customerAssetData.CustomerID == 0 {
		http.Error(w, "El campo CustomerID es obligatorio", http.StatusBadRequest)
		return
	}

	if customerAssetData.AssetID == 0 {
		http.Error(w, "El campo AssetID es obligatorio", http.StatusBadRequest)
		return
	}

	if customerAssetData.MarketValue == 0 {
		http.Error(w, "El campo MarketValue es obligatorio", http.StatusBadRequest)
		return
	}

	if customerAssetData.Description == "" {
		http.Error(w, "El campo Description es obligatorio", http.StatusBadRequest)
		return
	}

	if customerAssetData.CreditRequestID == 0 {
		http.Error(w, "El campo CreditRequest es obligatorio", http.StatusBadRequest)
		return
	}

	customerAsset := models.CustomerAsset{
		CustomerID:      customerAssetData.CustomerID,
		AssetID:         customerAssetData.AssetID,
		CreditRequestID: customerAssetData.CreditRequestID,
		MarketValue:     customerAssetData.MarketValue,
		Description:     customerAssetData.Description,
	}

	createdCustomerAsset, err := customerAssetService.CreateCustomerAsset(&customerAsset)

	if err != nil {
		if strings.Contains(err.Error(), "no existe") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Error al crear bien del cliente: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdCustomerAsset)
}

// UpdateCustomerAssetHandle godoc
// @Summary      Actualizar un bien del cliente
// @Description  Actualiza los datos de un bien del cliente existente
// @Tags         Customer Assets
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID del bien del cliente"
// @Param        request body UpdateCustomerAssetRequest true "Datos actualizados del bien"
// @Success      201 {object} models.CustomerAsset "Bien actualizado exitosamente"
// @Failure      400 {string} string "Solicitud inválida"
// @Failure      404 {string} string "Bien no encontrado"
// @Failure      500 {string} string "Error interno del servidor"
// @Router       /customer-assets/{id} [put]
func UpdateCustomerAssetHandle(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	customerAssetId := params["id"]
	id, err := strconv.Atoi(customerAssetId)
	if err != nil || id <= 0 {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var customerAssetData struct {
		CustomerID    uint    `json:"customerId"`
		AssetID       uint    `json:"assetId"`
		CreditRequest uint    `json:"creditRequestId"`
		MarketValue   float64 `json:"marketValue"`
		Description   string  `json:"description"`
	}

	err = json.NewDecoder(r.Body).Decode(&customerAssetData)

	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if customerAssetData.CustomerID == 0 {
		http.Error(w, "El campo CustomerID es obligatorio", http.StatusBadRequest)
		return
	}

	if customerAssetData.AssetID == 0 {
		http.Error(w, "El campo AssetID es obligatorio", http.StatusBadRequest)
		return
	}

	if customerAssetData.MarketValue == 0 {
		http.Error(w, "El campo MarketValue es obligatorio", http.StatusBadRequest)
		return
	}

	if customerAssetData.Description == "" {
		http.Error(w, "El campo Description es obligatorio", http.StatusBadRequest)
		return
	}

	if customerAssetData.CreditRequest == 0 {
		http.Error(w, "El campo CreditRequest es obligatorio", http.StatusBadRequest)
		return
	}

	customerAsset := models.CustomerAsset{
		CustomerID:  customerAssetData.CustomerID,
		AssetID:     customerAssetData.AssetID,
		MarketValue: customerAssetData.MarketValue,
		Description: customerAssetData.Description,
	}

	updatedCustomerAsset, err := customerAssetService.UpdateCustomerAsset(uint(id), &customerAsset)

	if err != nil {
		if strings.Contains(err.Error(), "no existe") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Error al actualizar bien del cliente: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(updatedCustomerAsset)
}

// DeleteCustomerAssetHandle godoc
// @Summary      Eliminar un bien del cliente
// @Description  Elimina un bien del cliente del sistema
// @Tags         Customer Assets
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID del bien del cliente"
// @Success      204 "Bien eliminado exitosamente"
// @Failure      400 {string} string "ID inválido"
// @Failure      404 {string} string "Bien no encontrado"
// @Failure      500 {string} string "Error interno del servidor"
// @Router       /customer-assets/{id} [delete]
func DeleteCustomerAssetHandle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = customerAssetService.DeleteCustomerAsset(uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "no existe") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Error al eliminar usuario: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
