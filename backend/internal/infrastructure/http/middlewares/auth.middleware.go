package middlewares

import (
	"context"
	"net/http"
	"strings"

	appconfig "github.com/JhonCamargo53/prueba-tecnica/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

var config = appconfig.Load()

var jwtSecret = []byte(config.JWTSecretKey)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "No autorizado. Faltante header Authorization", http.StatusUnauthorized)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			http.Error(w, "Formato Authorization inválido", http.StatusUnauthorized)
			return
		}

		tokenStr := parts[1]
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Token inválido o expirado", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["id"] == nil {
			http.Error(w, "Token sin identificador de usuario", http.StatusUnauthorized)
			return
		}

		requesterId := uint(claims["id"].(float64))

		ctx := context.WithValue(r.Context(), "requesterId", requesterId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
