package authHandler

import (
	"awesomeProject/pkg/models"
	"fmt"
	"net/http"
)

func (h authHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("refreshToken")

	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tknStr := c.Value

	if len(tknStr) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oldRefreshToken := models.RefreshToken{}
	oldRefreshToken.Find(tknStr, h.DB)

	if len(oldRefreshToken.Token) == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tokenString, refreshString, err := createToken(oldRefreshToken.Login, 15)

	authLine := fmt.Sprintf("Bearer %s", tokenString)

	w.Header().Set("Authorization", authLine)

	refreshToken := &models.RefreshToken{
		Login: oldRefreshToken.Login,
		Token: refreshString,
	}

	oldRefreshToken.Delete(h.DB)
	refreshToken.Add(h.DB)

	w.WriteHeader(http.StatusOK)
}
