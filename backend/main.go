package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/JhonCamargo53/prueba-tecnica/docs"
	"github.com/JhonCamargo53/prueba-tecnica/internal/config"
	"github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/bootstrap"
	databaseGorm "github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/database/gorm"
	"github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/database/gorm/migrations"
	"github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/http/middlewares"
	"github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/http/routes"
	"github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/logger"
	seed "github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/seeders"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title           Credit Risk Assessment API
// @version         1.0
// @description     API para gestión de solicitudes de crédito y evaluación de riesgo crediticio.

// @contact.name   API Support
// @contact.email  support@example.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:4000
// @BasePath  /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Ingresa el token JWT con el prefijo Bearer. Ejemplo: "Bearer {token}"

func main() {

	logger.InitLogger()

	config := config.Load()

	databaseGorm.DatabaseConnection()

	if err := migrations.AutoMigrateAll(databaseGorm.DB); err != nil {
		log.Fatal("Error en migración de modelos: ", err)
	}

	if err := seed.SeedAll(databaseGorm.DB); err != nil {
		log.Fatal("Error insertando datos iniciales: ", err)
	}

	bootstrap.InitializeDependencies(databaseGorm.DB, config)

	router := mux.NewRouter()

	routes.RegisterAllRoutes(router)
	router.Use(middlewares.RequestLogger)

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	fmt.Println(`
░█▀▄▀█ ─█▀▀█ ░█▄─░█ ─█▀▀█ ░█▀▀█ ░█▀▀▀ ░█▀▄▀█ ░█▀▀▀   ░█▄─░█ ▀▀█▀▀ ░█─── ▀█▀ ░█──░█ ░█▀▀▀ 
░█░█░█ ░█▄▄█ ░█░█░█ ░█▄▄█ ░█─▄▄ ░█▀▀▀ ░█░█░█ ░█▀▀▀   ░█░█░█ ─░█── ░█─── ░█─ ─░█░█─ ░█▀▀▀ 
░█──░█ ░█─░█ ░█──▀█ ░█─░█ ░█▄▄█ ░█▄▄▄ ░█──░█ ░█▄▄▄   ░█──▀█ ─░█── ░█▄▄█ ▄█▄ ──▀▄▀─ ░█▄▄▄`)

	log.Printf("Servidor corriendo en http://localhost:%s\n", config.Port)

	err := http.ListenAndServe(":"+config.Port, handler)
	if err != nil {
		log.Fatal("Error iniciando servidor: ", err)
	}
}
