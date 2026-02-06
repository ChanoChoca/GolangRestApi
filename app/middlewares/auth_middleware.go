package middlewares

import (
	"context"
	"flashpage/app/utils"
	"net/http"

	"github.com/golang-jwt/jwt"
)

type contextKey string

const UserIDKey contextKey = "userID"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth_token")
		if err != nil {
			http.Error(w, "No autenticado", http.StatusUnauthorized)
			return
		}

		token, err := utils.ParseJWT(cookie.Value)
		if err != nil || !token.Valid {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Claims inválidos", http.StatusUnauthorized)
			return
		}
		
		userID := int(claims["user_id"].(float64))

		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}