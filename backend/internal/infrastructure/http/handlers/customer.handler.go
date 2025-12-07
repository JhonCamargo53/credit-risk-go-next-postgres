package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/JhonCamargo53/prueba-tecnica/internal/application/services/customer"
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/gorilla/mux"
)

var customerService *customer.CustomerService

func InitCustomerHandler(s *customer.CustomerService) {
	customerService = s
}

// CreateCustomerRequest representa el cuerpo de la solicitud para crear un cliente
// @Description Datos para crear un nuevo cliente
type CreateCustomerRequest struct {
	Name           string  `json:"name" example:"María García"`
	Email          string  `json:"email" example:"maria.garcia@example.com"`
	PhoneNumber    string  `json:"phoneNumber" example:"+57 300 123 4567"`
	DocumentNumber string  `json:"documentNumber" example:"1234567890"`
	DocumentTypeId uint    `json:"documentTypeId" example:"1"`
	MonthlyIncome  float64 `json:"monthlyIncome" example:"5000000"`
}

// UpdateCustomerRequest representa el cuerpo de la solicitud para actualizar un cliente
// @Description Datos para actualizar un cliente existente
type UpdateCustomerRequest struct {
	Name           string  `json:"name" example:"María García Actualizada"`
	Email          string  `json:"email" example:"maria.garcia.updated@example.com"`
	PhoneNumber    string  `json:"phoneNumber" example:"+57 300 987 6543"`
	DocumentNumber string  `json:"documentNumber" example:"0987654321"`
	DocumentTypeId uint    `json:"documentTypeId" example:"2"`
	MonthlyIncome  float64 `json:"monthlyIncome" example:"6000000"`
}

// GetCustomersHandle godoc
// @Summary      Obtener todos los clientes
// @Description  Retorna una lista de todos los clientes del sistema
// @Tags         Customers
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {array} models.Customer "Lista de clientes"
// @Failure      500 {string} string "Error interno del servidor"
// @Router       /customers [get]
func GetCustomersHandle(w http.ResponseWriter, r *http.Request) {
	users, err := customerService.GetAllCustomers()
	if err != nil {
		http.Error(w, "No se pudieron obtener los clientes", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetCustomerHandle godoc
// @Summary      Obtener un cliente por ID
// @Description  Retorna los detalles de un cliente específico
// @Tags         Customers
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID del cliente"
// @Success      200 {object} models.Customer "Cliente encontrado"
// @Failure      400 {string} string "ID inválido"
// @Failure      404 {string} string "Cliente no encontrado"
// @Failure      500 {string} string "Error interno del servidor"
// @Router       /customers/{id} [get]
func GetCustomerHandle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	customerId := params["id"]
	id, err := strconv.Atoi(customerId)
	if err != nil || id <= 0 {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	customer, err := customerService.GetCustomerByID(uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "no existe") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Error al obtener cliente: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

// PostCustomerHandle godoc
// @Summary      Crear un nuevo cliente
// @Description  Crea un nuevo cliente en el sistema
// @Tags         Customers
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body CreateCustomerRequest true "Datos del cliente a crear"
// @Success      201 {object} models.Customer "Cliente creado exitosamente"
// @Failure      400 {string} string "Solicitud inválida"
// @Failure      409 {string} string "El cliente ya existe"
// @Failure      500 {string} string "Error interno del servidor"
// @Router       /customers [post]
func PostCustomerHandle(w http.ResponseWriter, r *http.Request) {

	requesterId := r.Context().Value("requesterId").(uint)

	var customerData struct {
		Name           string  `json:"name"`
		Email          string  `json:"email"`
		PhoneNumber    string  `json:"phoneNumber"`
		DocumentNumber string  `json:"documentNumber"`
		DocumentTypeId uint    `json:"documentTypeId"`
		CreatedByID    uint    `json:"createdById"`
		MonthlyIncome  float64 `json:"monthlyIncome"`
	}

	err := json.NewDecoder(r.Body).Decode(&customerData)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if customerData.Name == "" || customerData.Email == "" || customerData.PhoneNumber == "" ||
		customerData.DocumentNumber == "" || customerData.DocumentTypeId == 0 {
		http.Error(w, "Por favor, completa todos los campos obligatorios: Nombre, Email, Teléfono, Número de documento , Tipo de documentos", http.StatusBadRequest)
		return
	}

	customer := models.Customer{
		Name:           customerData.Name,
		Email:          customerData.Email,
		PhoneNumber:    customerData.PhoneNumber,
		DocumentNumber: customerData.DocumentNumber,
		MonthlyIncome:  customerData.MonthlyIncome,
		DocumentTypeId: customerData.DocumentTypeId,
		CreatedByID:    requesterId,
	}

	createdCustomer, err := customerService.CreateCustomer(&customer)

	if err != nil {
		if strings.Contains(err.Error(), "ya existe") {
			http.Error(w, err.Error(), http.StatusConflict)
		} else {
			http.Error(w, "Error al crear cliente: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdCustomer)
}

// UpdateCustomerHandle godoc
// @Summary      Actualizar un cliente
// @Description  Actualiza los datos de un cliente existente
// @Tags         Customers
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID del cliente"
// @Param        request body UpdateCustomerRequest true "Datos actualizados del cliente"
// @Success      200 {object} models.Customer "Cliente actualizado exitosamente"
// @Failure      400 {string} string "Solicitud inválida"
// @Failure      404 {string} string "Cliente no encontrado"
// @Failure      409 {string} string "El email ya existe"
// @Failure      500 {string} string "Error interno del servidor"
// @Router       /customers/{id} [put]
func UpdateCustomerHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var customerData struct {
		Name           string  `json:"name"`
		Email          string  `json:"email"`
		PhoneNumber    string  `json:"phoneNumber"`
		DocumentNumber string  `json:"documentNumber"`
		DocumentTypeId uint    `json:"documentTypeId"`
		MonthlyIncome  float64 `json:"monthlyIncome"`
	}

	if err := json.NewDecoder(r.Body).Decode(&customerData); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, "Cliente no encontrado", http.StatusNotFound)
		return
	}

	customer := &models.Customer{
		Name:           customerData.Name,
		Email:          customerData.Email,
		PhoneNumber:    customerData.PhoneNumber,
		DocumentNumber: customerData.DocumentNumber,
		DocumentTypeId: customerData.DocumentTypeId,
		MonthlyIncome:  customerData.MonthlyIncome,
	}

	updatedCustomer, err := customerService.UpdateCustomer(uint(id), customer)

	if err != nil {
		if strings.Contains(err.Error(), "ya existe") {
			http.Error(w, err.Error(), http.StatusConflict)
		} else if strings.Contains(err.Error(), "no existe") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Error al crear usuario: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedCustomer)
}

// DeleteCustomerHandle godoc
// @Summary      Eliminar un cliente
// @Description  Elimina un cliente del sistema
// @Tags         Customers
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID del cliente"
// @Success      204 "Cliente eliminado exitosamente"
// @Failure      400 {string} string "ID inválido"
// @Failure      500 {string} string "Error interno del servidor"
// @Router       /customers/{id} [delete]
func DeleteCustomerHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = customerService.DeleteCustomer(uint(id))
	if err != nil {
		http.Error(w, "No se pudo eliminar el cliente: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
