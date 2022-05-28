package authHandler

import (
	"awesomeProject/pkg/models"
	"awesomeProject/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
)

func (h authHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	credentials := &Credentials{}

	err := json.NewDecoder(r.Body).Decode(&credentials)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the expected password from our in memory map
	user := &models.User{}
	userNotFound := user.IsUserByLogin(credentials.Login, h.DB)

	if userNotFound == nil && errors.Is(userNotFound, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	hashedPass, err := utils.HashPassword(credentials.Password)

	if err != nil {
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
