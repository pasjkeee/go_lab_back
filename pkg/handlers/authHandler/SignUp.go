package authHandler

import (
	"awesomeProject/pkg/logging"
	"awesomeProject/pkg/models"
	"awesomeProject/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h authHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	logger := logging.GetLogger()
	credentials := &Credentials{}

	err := json.NewDecoder(r.Body).Decode(&credentials)

	if err != nil {
		logger.Error("Can not decode credentials from request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the expected password from our in memory map
	logger.Debugf("Try to find user %s in database", credentials.Login)
	user := &models.User{}
	userNotFound := user.IsUserByLogin(credentials.Login, h.DB)

	if userNotFound == nil {
		logger.Warnf("User %s is already exists", credentials.Login)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	hashedPass, err := utils.HashPassword(credentials.Password)

	if err != nil {
		logger.Infof("Can hash password for user %s", credentials.Login)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var newUser = models.User{
		Login:    credentials.Login,
		Password: hashedPass,
	}

	h.DB.Create(&newUser)

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

	logger.Infof("Successfull signup and authorization for %s", credentials.Login)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&Token{Token: authLine, Login: credentials.Login})
}
