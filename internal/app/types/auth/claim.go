package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = "pineapplejuice"

type Claims struct {
	Email string
	jwt.StandardClaims
}

func CreateJWT(email string) (string, error) {
	claims := Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, &claims).SignedString([]byte(jwtKey))
}
