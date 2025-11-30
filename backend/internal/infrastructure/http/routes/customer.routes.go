package routes

import (
	"github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/http/handlers"
	"github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/http/middlewares"
	"github.com/gorilla/mux"
)

func RegisterCustomerRoutes(router *mux.Router) {
	customerRouter := router.PathPrefix("/customers").Subrouter()
	customerRouter.Use(middlewares.AuthMiddleware)
	customerRouter.HandleFunc("", handlers.GetCustomersHandle).Methods("GET")
	customerRouter.HandleFunc("/{id}", handlers.GetCustomerHandle).Methods("GET")
	customerRouter.HandleFunc("", handlers.PostCustomerHandle).Methods("POST")
	customerRouter.HandleFunc("/{id}", handlers.UpdateCustomerHandle).Methods("PUT")
	customerRouter.HandleFunc("/{id}", handlers.DeleteCustomerHandle).Methods("DELETE")
}
