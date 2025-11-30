package routes

import (
	"github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/http/handlers"
	"github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/http/middlewares"
	"github.com/gorilla/mux"
)

func RegisterAssetRoutes(router *mux.Router) {
	userRouter := router.PathPrefix("/assets").Subrouter()
	userRouter.Use(middlewares.AuthMiddleware)
	userRouter.HandleFunc("", handlers.GetAssetsHandle).Methods("GET")
}
