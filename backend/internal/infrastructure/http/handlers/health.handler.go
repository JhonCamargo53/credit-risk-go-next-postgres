package handlers

import "net/http"

// HealthController godoc
// @Summary      Health check
// @Description  Verifica el estado del servidor
// @Tags         Health
// @Accept       json
// @Produce      plain
// @Success      200 {string} string "Hello, I'm working fine!"
// @Router       /health [get]
func HealthController(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, I'm working fine!"))
}
