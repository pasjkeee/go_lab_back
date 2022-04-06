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

		if r.URL.String() != "/login" {
			//c, err := r.Cookie("token")
			//fmt.Println(c)

			//if err != nil {
			//	if err == http.ErrNoCookie {
			//		// If the cookie is not set, return an unauthorized status
			//		w.WriteHeader(http.StatusUnauthorized)
			//		return
			//	}
			//	// For any other type of error, return a bad request status
			//	w.WriteHeader(http.StatusBadRequest)
			//	return
			//}

			//tknStr := c.Value

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
