package authHandler

import (
	"awesomeProject/pkg/logging"
	"awesomeProject/pkg/models"
	"encoding/json"
	"net/http"
)

func (h authHandler) SignOut(w http.ResponseWriter, r *http.Request) {

	logger := logging.GetLogger()
	credentials := &Credentials{}

	err := json.NewDecoder(r.Body).Decode(&credentials)

	if err != nil {
		logger.Error("Can not decode credentials from request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Del("Authorization")
	refreshToken := &models.RefreshToken{}

	refreshToken.DeleteByLogin(credentials.Login, h.DB)
	logger.Infof("Successfull logout for %s", credentials.Login)
	w.WriteHeader(http.StatusOK)
}
