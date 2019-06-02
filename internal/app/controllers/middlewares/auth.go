package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/morenocantoj/petstore-go/internal/app/controllers"
	"github.com/morenocantoj/petstore-go/internal/app/types/auth"
	"github.com/morenocantoj/petstore-go/internal/app/types/responses"
)

// AuthMiddleware controlles that subsequent requests are authorized by token provided
type AuthMiddleware struct {
	controllers.BaseController
}

// ValidateJWT validates the json-web token passed in the request
func (a *AuthMiddleware) ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(writter http.ResponseWriter, req *http.Request) {
		tokenString, err := getTokenFromHeader(req)
		if err != nil {
			response := responses.Forbidden{HttpError: responses.HttpError{Code: 403, Message: fmt.Sprintf("%s", err)}}
			a.WriteResponse(response, writter, http.StatusInternalServerError)
			return
		}
		jwtToken, err := auth.VerifyTokenString(tokenString)
		if !jwtToken.Valid || err == jwt.ErrSignatureInvalid {
			response := responses.Unauthorized{
				HttpError: responses.HttpError{Code: 401, Message: "Token not valid"},
				LoginURL:  req.Host + "/auth/",
			}
			a.WriteResponse(response, writter, http.StatusUnauthorized)
			return
		}
		if err != nil {
			response := responses.ServerError{HttpError: responses.HttpError{Code: 500, Message: "Error ocurred decoding the token"}}
			a.WriteResponse(response, writter, http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(writter, req)
	})
}

func getTokenFromHeader(req *http.Request) (string, error) {
	authHeader := req.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("Authorization header needed for this request")
	}
	// Authorization header format: Bearer {token}
	token := strings.Split(authHeader, " ")
	if len(token) != 2 {
		return "", fmt.Errorf("Authorization header is not correctly formatted")
	}
	return token[1], nil
}
