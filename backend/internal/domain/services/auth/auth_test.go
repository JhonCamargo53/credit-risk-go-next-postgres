package auth

import (
	"testing"
	"time"

	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func TestLogin_UsuarioNoExiste(t *testing.T) {
	// Repo sin usuarios
	userRepo := NewMockUserRepository(nil)
	secret := []byte("test-secret")

	service := NewAuthService(userRepo, secret)

	token, err := service.Login("noexiste@example.com", "pass")

	if err == nil {
		t.Fatalf("se esperaba error por usuario inexistente")
	}
	if token != "" {
		t.Fatalf("no se esperaba token cuando el usuario no existe")
	}
}

func TestLogin_PasswordIncorrecto(t *testing.T) {

	// Usuario con contraseña correcta
	hashed, err := bcrypt.GenerateFromPassword([]byte("correct-password"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("error generando hash de password: %v", err)
	}

	user := &models.User{
		ID:       1,
		Name:     "Juan",
		Email:    "juan@example.com",
		Password: string(hashed),
		RoleId:   1,
	}

	userRepo := NewMockUserRepository([]*models.User{user})
	secret := []byte("test-secret")

	service := NewAuthService(userRepo, secret)

	token, err := service.Login("juan@example.com", "incorrect-password")

	if err == nil {
		t.Fatalf("se esperaba error por password incorrecto")
	}
	if token != "" {
		t.Fatalf("no se esperaba token cuando el password es incorrecto")
	}
}

func TestLogin_Exitoso_GeneraTokenValido(t *testing.T) {

	// Usuario con contraseña correcta
	hashed, err := bcrypt.GenerateFromPassword([]byte("my-password"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("error generando hash de password: %v", err)
	}

	user := &models.User{
		ID:       42,
		Name:     "Juan Test",
		Email:    "juan@example.com",
		Password: string(hashed),
		RoleId:   1,
	}

	userRepo := NewMockUserRepository([]*models.User{user})
	secret := []byte("test-secret")

	service := NewAuthService(userRepo, secret)

	tokenString, err := service.Login("juan@example.com", "my-password")
	if err != nil {
		t.Fatalf("no se esperaba error en login exitoso: %v", err)
	}
	if tokenString == "" {
		t.Fatalf("se esperaba un token JWT no vacío")
	}

	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			t.Fatalf("método de firma inesperado: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		t.Fatalf("error al parsear token: %v", err)
	}

	if !parsedToken.Valid {
		t.Fatalf("se esperaba un token válido")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		t.Fatalf("no se pudieron leer las claims como MapClaims")
	}

	if claims["email"] != "juan@example.com" {
		t.Fatalf("email en claims inesperado, se obtuvo=%v", claims["email"])
	}
	if claims["name"] != "Juan Test" {
		t.Fatalf("name en claims inesperado, se obtuvo=%v", claims["name"])
	}
	if claims["roleId"] != float64(1) { // jwt numérico → float64
		t.Fatalf("roleId en claims inesperado, se obtuvo=%v", claims["roleId"])
	}
	if claims["id"] != float64(42) {
		t.Fatalf("id en claims inesperado, se obtuvo=%v", claims["id"])
	}

	expFloat, ok := claims["exp"].(float64)
	if !ok {
		t.Fatalf("exp no es numérico, se obtuvo=%T", claims["exp"])
	}
	expTime := time.Unix(int64(expFloat), 0)
	if time.Now().After(expTime) {
		t.Fatalf("exp ya está vencido, exp=%v, now=%v", expTime, time.Now())
	}
}
