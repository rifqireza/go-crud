package config

import (
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrInvalidKey
			}
			return []byte("secret_key"), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		r.Header.Set("Authorization", tokenString)
		next.ServeHTTP(w, r)
	})
}
