package middleware

import (
	"context"
	"gateway/config"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		// Parse and validate token
		token, err := jwt.ParseWithClaims(tokenString, &config.JWTCustomClass{}, func(token *jwt.Token) (interface{}, error) {
			return config.GetJWTSerectKey(), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Add claims to context
		claims := token.Claims.(*config.JWTCustomClass)
		ctx := context.WithValue(r.Context(), "username", claims.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}
