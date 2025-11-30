package routes

import (
	"github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/http/handlers"
	"github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/http/middlewares"
	"github.com/gorilla/mux"
)

func RegisterDocumentTypeRoutes(router *mux.Router) {
	userRouter := router.PathPrefix("/document-types").Subrouter()
	userRouter.Use(middlewares.AuthMiddleware)
	userRouter.HandleFunc("", handlers.GetDocumentTypesHandle).Methods("GET")
}
