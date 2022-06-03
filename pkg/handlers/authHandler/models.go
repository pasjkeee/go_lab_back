package authHandler

import (
	"awesomeProject/pkg/logging"
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
	logger := logging.GetLogger()
	claims.Login = login
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Unix(),
	}
	logger.Infof("Token created for %s", login)
}

type Token struct {
	Token string `json:"token"`
	Login string `json:"login"`
}
