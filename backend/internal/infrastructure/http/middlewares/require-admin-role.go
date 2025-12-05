package middlewares

import (
	"context"
	"net/http"

	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	database "github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/database/gorm"
)

var Roles = map[string]int{
	"ADMIN": 1,
	"USER":  2,
}

func RequireAdminRole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		requesterIDCtx := r.Context().Value("requesterId")
		if requesterIDCtx == nil {
			http.Error(w, "No autorizado. Usuario no encontrado en contexto", http.StatusUnauthorized)
			return
		}

		requesterID, ok := requesterIDCtx.(uint)
		if !ok {
			http.Error(w, "Error procesando identificador del usuario", http.StatusUnauthorized)
			return
		}

		var user models.User
		if err := database.DB.First(&user, requesterID).Error; err != nil {
			http.Error(w, "Usuario no encontrado", http.StatusUnauthorized)
			return
		}

		if user.RoleId != uint(Roles["ADMIN"]) {
			http.Error(w, "No tiene acceso a este recurso", http.StatusForbidden)
			return
		}
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
