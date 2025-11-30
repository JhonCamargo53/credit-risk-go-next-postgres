package services

import (
	"fmt"
	"time"

	appconfig "github.com/JhonCamargo53/prueba-tecnica/internal/config"
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/database"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var config = appconfig.Load()

var jwtSecret = []byte(config.JWTSecretKey)

func LoginService(email string, password string) (string, error) {
	var user models.User

	err := database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return "", fmt.Errorf("usuario o contraseña incorrectos")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("usuario o contraseña incorrectos")
	}

	claims := jwt.MapClaims{
		"id":     user.ID,
		"email":  user.Email,
		"name":   user.Name,
		"roleId": user.RoleId,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("error al generar el token")
	}

	return tokenString, nil
}
