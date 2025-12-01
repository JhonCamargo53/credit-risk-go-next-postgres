package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterAboutTypeRoutes(router *mux.Router) {

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		response := map[string]interface{}{
			"name":        "Credit Risk Management System",
			"authors":     "Jhon Camargo - GitHub: JhonCamargo53",
			"version":     "1.0.0",
			"language":    "Go + Next.js",
			"description": "Sistema gestor de solicitudes de crédito con evaluación de riesgo, generación de reportes financieros y motor IA mock.",
			"features": []string{
				"Evaluación de riesgo crediticio",
				"Gestión de clientes, activos y solicitudes",
				"Motor IA mock explicable",
				"Generación automática de reportes",
				"Integración con base de datos PostgreSQL",
			},
		}

		json.NewEncoder(w).Encode(response)
	}).Methods("GET")
}
