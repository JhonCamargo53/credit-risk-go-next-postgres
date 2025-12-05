package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/services/auth"
)

var authService *auth.AuthService

func InitAuthHandler(service *auth.AuthService) {
	authService = service
}

type LoginRequest struct {
	Email    string `json:"email" example:"admin@example.com"`
	Password string `json:"password" example:"password123"`
}
type LoginResponse struct {
	Token string `json:"token,omitempty" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	Error string `json:"error,omitempty" example:""`
}

// LoginHandle godoc
// @Summary      Iniciar sesión
// @Description  Autentica a un usuario y retorna un token JWT
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body LoginRequest true "Credenciales de inicio de sesión"
// @Success      200 {object} LoginResponse "Token de autenticación"
// @Failure      400 {object} LoginResponse "Solicitud inválida"
// @Failure      401 {object} LoginResponse "Credenciales inválidas"
// @Failure      403 {object} LoginResponse "Usuario no activo"
// @Router       /login [post]
func LoginHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if authService == nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(LoginResponse{Error: "authService no inicializado"})
		return
	}

	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(LoginResponse{Error: "Bad request"})
		return
	}

	if req.Email == "" || req.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(LoginResponse{Error: "El correo y la contraseña son obligatorios"})
		return
	}

	token, err := authService.Login(req.Email, req.Password)
	if err != nil {
		if err.Error() == "El usuario no está activo" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(LoginResponse{Error: err.Error()})
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(LoginResponse{Error: err.Error()})
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(LoginResponse{Token: token})
}
