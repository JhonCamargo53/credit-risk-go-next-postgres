package routes

import (
	"github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/http/handlers"
	"github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/http/middlewares"
	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router) {
	userRouter := router.PathPrefix("/users").Subrouter()
	userRouter.Use(middlewares.AuthMiddleware)
	userRouter.Use(middlewares.RequireAdminRole)
	userRouter.HandleFunc("", handlers.GetUsersHandle).Methods("GET")
	userRouter.HandleFunc("/{id}", handlers.GetUserHandle).Methods("GET")
	userRouter.HandleFunc("", handlers.PostUserHandle).Methods("POST")
	userRouter.HandleFunc("/{id}", handlers.UpdateUserHandle).Methods("PUT")
	userRouter.HandleFunc("/{id}", handlers.DeleteUserHandle).Methods("DELETE")
}
