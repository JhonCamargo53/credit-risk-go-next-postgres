package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/services"
)

var documentTypeService *services.DocumentTypeService

// GetDocumentTypesHandle godoc
// @Summary      Obtener todos los tipos de documento
// @Description  Retorna una lista de todos los tipos de documento de identidad disponibles
// @Tags         Document Types
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {array} models.DocumentType "Lista de tipos de documento"
// @Failure      500 {string} string "Error interno del servidor"
// @Router       /document-types [get]
func GetDocumentTypesHandle(w http.ResponseWriter, r *http.Request) {
	documentTypes, err := documentTypeService.GetAllCreditStatuses()
	if err != nil {
		http.Error(w, "No se pudieron obtener los tipos de documento de identidad", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(documentTypes)
}
