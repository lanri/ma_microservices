package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/lanri/diktum-service/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized: No token", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := utils.ValidateJWT(tokenStr)
		if err != nil {
			fmt.Println("JWT validation error:", err)
			http.Error(w, "Forbidden: Invalid token", http.StatusForbidden)
			return
		}

		fmt.Println("✅ Token valid for sub:", claims.Subject)
		// Lanjut ke handler
		next.ServeHTTP(w, r)
	})
}
