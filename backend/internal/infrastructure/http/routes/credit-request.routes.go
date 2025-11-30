package routes

import (
	"github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/http/handlers"
	"github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/http/middlewares"
	"github.com/gorilla/mux"
)

func RegisterCreditRequestRoutes(router *mux.Router) {
	creditRequestRouter := router.PathPrefix("/credit-requests").Subrouter()
	creditRequestRouter.Use(middlewares.AuthMiddleware)
	creditRequestRouter.HandleFunc("", handlers.GetCreditRequestsHandle).Methods("GET")
	creditRequestRouter.HandleFunc("/{id}", handlers.GetCreditRequestHandle).Methods("GET")
	creditRequestRouter.HandleFunc("", handlers.PostCreditRequestHandle).Methods("POST")
	creditRequestRouter.HandleFunc("/{id}", handlers.UpdateCreditRequestHandle).Methods("PUT")
	creditRequestRouter.HandleFunc("/{id}", handlers.DeleteCreditRequestHandle).Methods("DELETE")
}
