package authHandler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

func createExpirationTime(timeInMinute time.Duration) time.Time {
	return time.Now().Add(timeInMinute * time.Minute)
}

func createToken(login string, timeInMinute time.Duration) (string, string, error) {
	claims := &Claims{}

	expirationTime := createExpirationTime(timeInMinute)
	claims.Create(expirationTime, login)

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	refreshToken := uuid.New()
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", "", err
	}

	return tokenString, refreshToken.String(), nil
}
