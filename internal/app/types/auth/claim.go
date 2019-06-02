package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JwtKey = "pineapplejuice"

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

	return jwt.NewWithClaims(jwt.SigningMethodHS256, &claims).SignedString([]byte(JwtKey))
}

// VerifyTokenString verifies token string from request parsing it with jwt
func VerifyTokenString(token string) (*jwt.Token, error) {
	claims := Claims{}
	jwtToken, err := jwt.ParseWithClaims(token, &claims, func(tkn *jwt.Token) (interface{}, error) {
		return []byte(JwtKey), nil
	})
	return jwtToken, err
}
