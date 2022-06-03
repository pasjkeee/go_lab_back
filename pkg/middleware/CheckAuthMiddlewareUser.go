package middleware

import (
	"awesomeProject/pkg/handlers/authHandler"
	"awesomeProject/pkg/logging"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

var jwtKey = []byte("my_secret_key")

func CheckAuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := logging.GetLogger()
		logger.Infof("New request on %s", r.URL.String())

		if !(r.URL.String() == "/login" || r.URL.String() == "/signup" || strings.Contains(r.URL.String(), "/ws/")) {

			authHeader := r.Header.Get("Authorization")
			if len(strings.Split(authHeader, " ")) < 2 {
				logger.Warn("User didnt have token")
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
					logger.Warn("User have not valid signature")
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				logger.Warn("User have bad request")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if !tkn.Valid {
				logger.Warn("User have not valid token")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
