package authHandler

import (
	"awesomeProject/pkg/logging"
	"awesomeProject/pkg/models"
	"awesomeProject/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h authHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	logger := logging.GetLogger()

	credentials := &Credentials{}

	err := json.NewDecoder(r.Body).Decode(&credentials)

	if err != nil {
		logger.Error("Can not decode credentials from request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	logger.Debug("Try to find user by login in database")

	// Get the expected password from our in memory map
	user := &models.User{}
	user.FindUserByLogin(credentials.Login, h.DB)

	if len(user.Password) == 0 || utils.CheckPasswordHash(user.Password, credentials.Password) {
		logger.Errorf("Can not authorize user %s", user.Login)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tokenString, refreshString, err := createToken(credentials.Login, 15)
	if err != nil {
		logger.Errorf("Can't create token for %s", credentials.Login)
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

	logger.Infof("Successfull authorization for %s", credentials.Login)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&Token{Token: authLine, Login: credentials.Login})
}
