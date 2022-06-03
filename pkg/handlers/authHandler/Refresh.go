package authHandler

import (
	"awesomeProject/pkg/logging"
	"awesomeProject/pkg/models"
	"fmt"
	"net/http"
)

func (h authHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	logger := logging.GetLogger()
	c, err := r.Cookie("refreshToken")

	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			logger.Error("User unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		logger.Error("User has a bad request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tknStr := c.Value

	if len(tknStr) == 0 {
		logger.Error("User didnt have token")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	logger.Debug("Try to find refresh token in database")
	oldRefreshToken := models.RefreshToken{}
	oldRefreshToken.Find(tknStr, h.DB)

	if len(oldRefreshToken.Token) == 0 {
		logger.Error("User didnt have old refresh token")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tokenString, refreshString, err := createToken(oldRefreshToken.Login, 15)

	if err != nil {
		logger.Errorf("Can't create token for %s", oldRefreshToken.Login)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	authLine := fmt.Sprintf("Bearer %s", tokenString)

	w.Header().Set("Authorization", authLine)

	refreshToken := &models.RefreshToken{
		Login: oldRefreshToken.Login,
		Token: refreshString,
	}

	oldRefreshToken.Delete(h.DB)
	refreshToken.Add(h.DB)

	logger.Infof("Successfully created token for %s", oldRefreshToken.Login)

	w.WriteHeader(http.StatusOK)
}
