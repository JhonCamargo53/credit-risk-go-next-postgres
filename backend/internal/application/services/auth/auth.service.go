package auth

import (
	"fmt"
	"time"

	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/ports"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo  ports.UserRepository
	jwtSecret []byte
}

func NewAuthService(userRepo ports.UserRepository, jwtSecret []byte) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (s *AuthService) Login(email string, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)

	if err != nil || user == nil {
		return "", fmt.Errorf("usuario o contraseña incorrectos")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
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
	tokenString, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", fmt.Errorf("error al generar el token")
	}

	return tokenString, nil
}
