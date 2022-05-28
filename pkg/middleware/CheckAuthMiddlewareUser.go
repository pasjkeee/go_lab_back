package middleware

import (
	"awesomeProject/pkg/handlers/authHandler"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

var jwtKey = []byte("my_secret_key")

func CheckAuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if !(r.URL.String() == "/login" || r.URL.String() == "/signup" || strings.Contains(r.URL.String(), "/ws/")) {

			authHeader := r.Header.Get("Authorization")
			if len(strings.Split(authHeader, " ")) < 2 {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			tknString := strings.Split(authHeader, " ")[1]

			claims := &authHandler.Claims{}

			tkn, err := jwt.ParseWithClaims(tknString, claims, func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})

			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if !tkn.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
