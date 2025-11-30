package routes

import (
	"github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/http/handlers"
	"github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/http/middlewares"
	"github.com/gorilla/mux"
)

func RegisterCreditStatusRoutes(router *mux.Router) {
	userRouter := router.PathPrefix("/credit-statuses").Subrouter()
	userRouter.Use(middlewares.AuthMiddleware)
	userRouter.HandleFunc("", handlers.GetCreditStatusesHandle).Methods("GET")
}
