package handlers

import (
	"encoding/json"
	"net/http"

	documentType "github.com/JhonCamargo53/prueba-tecnica/internal/domain/services/document-type"
)

var documentTypeService *documentType.DocumentTypeService

func InitDocumentTypeHandler(s *documentType.DocumentTypeService) {
	documentTypeService = s
}

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

	if documentTypeService == nil {
		http.Error(w, "documentTypeService no inicializado", http.StatusInternalServerError)
		return
	}

	documentTypes, err := documentTypeService.GetAllDocumentTypes()
	if err != nil {
		http.Error(w, "Error obteniendo tipos de documento", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(documentTypes)
}
