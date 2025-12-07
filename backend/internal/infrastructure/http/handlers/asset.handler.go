package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/JhonCamargo53/prueba-tecnica/internal/application/services/asset"
)

var assetService *asset.AssetService

func InitAssetHandler(service *asset.AssetService) {
	assetService = service
}

// GetAssetsHandle godoc
// @Summary      Obtener todos los tipos de bienes
// @Description  Retorna una lista de todos los tipos de bienes disponibles
// @Tags         Assets
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {array} models.Asset "Lista de tipos de bienes"
// @Failure      500 {string} string "Error interno del servidor"
// @Router       /assets [get]
func GetAssetsHandle(w http.ResponseWriter, r *http.Request) {

	assets, err := assetService.GetAllAssets()

	if err != nil {
		http.Error(w, "No se pudieron obtener todos los tipos de bienes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(assets)
}
