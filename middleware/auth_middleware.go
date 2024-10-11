package middleware

import (
	"login-user/utils"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "token ausente", http.StatusUnauthorized)
			return
		}

		tokenstring := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ValidateJWT(tokenstring)

		if err != nil {
			http.Error(w, "Token invalido", http.StatusUnauthorized)
			return
		}

		r.Header.Set("UserID", string(claims.UserID))
		r.Header.Set("Email", claims.Email)

		next.ServeHTTP(w, r)
	})

}
