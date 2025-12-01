package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadEnvFile() {

	inRailway := os.Getenv("RAILWAY_ENVIRONMENT") != "" ||
		os.Getenv("RAILWAY_PROJECT_ID") != ""

	if inRailway {
		log.Println("Modo Railway detectado: NO se cargan archivos .env")
		return
	}

	env := os.Getenv("ENV")
	filename := ".env.development"

	if env == "production" {
		filename = ".env.production"
	}

	if err := godotenv.Load(filename); err != nil {
		log.Printf("No se pudo cargar %s: %v", filename, err)
	} else {
		log.Printf("Archivo de entorno cargado: %s", filename)
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

type Config struct {
	ENV          string
	Port         string
	DatabaseURL  string
	JWTSecretKey string
}

func Load() *Config {

	loadEnvFile()

	return &Config{
		ENV:  getEnv("ENV", "development"),
		Port: getEnv("PORT", "4000"),

		DatabaseURL: getEnv(
			"DATABASE_URL",
			"postgres://admin:20Acc3ss25@localhost:5435/credit?sslmode=disable",
		),

		JWTSecretKey: getEnv("JWT_SECRET_KEY", "default-secret-key"),
	}
}
