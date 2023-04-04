package middlewares

import (
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verificar que el token de autorización esté presente en la solicitud
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Authorization token missing", http.StatusUnauthorized)
			return
		}
		tokenString = tokenString[len("Bearer "):]
		// Verificar y decodificar el token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			secretKey := os.Getenv("SECRET_KEY")
			return []byte(secretKey), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}
		// Si el token es válido, permitir que la solicitud continúe
		next.ServeHTTP(w, r)
	})
}
