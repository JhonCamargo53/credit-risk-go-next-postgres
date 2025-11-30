package database

import (
	"log"

	appconfig "github.com/JhonCamargo53/prueba-tecnica/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var config = appconfig.Load()
var DSN = config.DatabaseURL
var DB *gorm.DB

func DatabaseConnection() {
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect: ", err)
	} else {
		log.Println("Database connected successfully")
	}
}
