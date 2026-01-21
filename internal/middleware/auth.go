package middleware

import (
	"context"
	"net/http"
	"strings"

	"example.com/internal/utils"
	"github.com/golang-jwt/jwt/v5"
)

func JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}
		tokenStr := strings.TrimPrefix(header, "Bearer")
		token, err := utils.ValidateToken(tokenStr)
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}
		email, ok := claims["email"].(string)
		if !ok {
			http.Error(w, "Invalid token email claim", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "email", email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
