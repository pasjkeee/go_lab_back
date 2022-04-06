package authHandler

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Credentials struct {
	Password string `json:"password"`
	Login    string `json:"login"`
}

type Claims struct {
	Login string `json:"login"`
	jwt.StandardClaims
}

func (claims *Claims) Create(time time.Time, login string) {
	claims.Login = login
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Unix(),
	}
}

type Token struct {
	Token string `json:"token"`
	Login string `json:"login"`
}
