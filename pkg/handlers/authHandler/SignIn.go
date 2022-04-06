package authHandler

import (
	"awesomeProject/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h authHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	credentials := &Credentials{}

	err := json.NewDecoder(r.Body).Decode(&credentials)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the expected password from our in memory map
	user := &models.User{}
	user.FindUserByLogin(credentials.Login, h.DB)

	if len(user.Password) == 0 || user.Password != credentials.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tokenString, refreshString, err := createToken(credentials.Login, 15)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	authLine := fmt.Sprintf("Bearer %s", tokenString)

	w.Header().Set("Authorization", authLine)

	refreshToken := &models.RefreshToken{
		Login: credentials.Login,
		Token: refreshString,
	}

	refreshToken.Add(h.DB)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&Token{Token: authLine, Login: credentials.Login})
}
