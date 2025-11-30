package routes

import (
	"github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/http/handlers"
	"github.com/gorilla/mux"
)

func RegisterMetricRoutes(router *mux.Router) {
	metricRouter := router.PathPrefix("/metrics").Subrouter()
	metricRouter.HandleFunc("", handlers.Metrics).Methods("GET")
}
