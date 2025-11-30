package routes

import (
	"github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/http/handlers"
	"github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/http/middlewares"
	"github.com/gorilla/mux"
)

func RegisterCustomerAssetRoutes(router *mux.Router) {
	customerAssetRouter := router.PathPrefix("/customer-assets").Subrouter()
	customerAssetRouter.Use(middlewares.AuthMiddleware)
	customerAssetRouter.HandleFunc("", handlers.GetCustomerAssetsHandle).Methods("GET")
	customerAssetRouter.HandleFunc("", handlers.PostCustomerAssetHandle).Methods("POST")
	customerAssetRouter.HandleFunc("/{id}", handlers.UpdateCustomerAssetHandle).Methods("PUT")
	customerAssetRouter.HandleFunc("/{id}", handlers.DeleteCustomerAssetHandle).Methods("DELETE")

}
