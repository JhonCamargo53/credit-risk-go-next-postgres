package routes

import (
	"github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/http/handlers"
	"github.com/gorilla/mux"
)

func RegisterHealthRoutes(router *mux.Router) {
	router.HandleFunc("/health", handlers.HealthController).Methods("GET")
}
