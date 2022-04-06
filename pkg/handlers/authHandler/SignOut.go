package authHandler

import (
	"awesomeProject/pkg/models"
	"encoding/json"
	"net/http"
)

func (h authHandler) SignOut(w http.ResponseWriter, r *http.Request) {

	credentials := &Credentials{}

	err := json.NewDecoder(r.Body).Decode(&credentials)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Del("Authorization")
	refreshToken := &models.RefreshToken{}
	refreshToken.DeleteByLogin(credentials.Login, h.DB)
	w.WriteHeader(http.StatusOK)
}
