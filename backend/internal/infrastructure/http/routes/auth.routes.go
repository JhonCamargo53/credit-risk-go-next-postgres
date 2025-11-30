package routes

import (
	"github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/http/handlers"
	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(router *mux.Router) {
	router.HandleFunc("/login", handlers.LoginHandle).Methods("POST")
}
