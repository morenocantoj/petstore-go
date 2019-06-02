package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/morenocantoj/petstore-go/internal/app/database"
	"github.com/morenocantoj/petstore-go/internal/app/types/auth"
	"github.com/morenocantoj/petstore-go/internal/app/types/classes"
	"github.com/morenocantoj/petstore-go/internal/app/types/responses"
	"github.com/morenocantoj/petstore-go/internal/pkg/utils/errors"
)

// AuthController defines authentication request logic
type AuthController struct {
	BaseController
}

// Create a JSON-Web-Token if user
func (a *AuthController) Create(writter http.ResponseWriter, req *http.Request) {
	fmt.Println("POST /auth")
	db := database.Connector{Connection: database.ConnectToDatabase()}
	defer db.Connection.Close()

	// Read body
	body := make(map[string]string)
	err := json.NewDecoder(req.Body).Decode(&body)
	errors.Check(err)
	email := body["email"]
	existingUser, err := existsUser(email, db.Connection)
	if err != nil {
		response := responses.NotFound{HttpError: responses.HttpError{Code: 404, Message: "User does not exist in our database"}}
		a.WriteResponse(response, writter, http.StatusNotFound)

	} else {
		if existingUser.CheckPassword(body["password"]) {
			token, err := auth.CreateJWT(email)
			if err != nil {
				response := responses.ServerError{HttpError: responses.HttpError{Code: 500, Message: "Error creating auth token"}}
				a.WriteResponse(response, writter, http.StatusInternalServerError)

			} else {
				response := responses.AuthResponseOK{Code: 200, Message: fmt.Sprintf("User %s logged in successfully", email), Token: token}
				a.WriteResponse(response, writter, http.StatusOK)
			}

		} else {
			response := responses.Forbidden{HttpError: responses.HttpError{Code: 403, Message: "Password not valid"}}
			a.WriteResponse(response, writter, http.StatusNotFound)
		}
	}
}

func existsUser(email string, db *gorm.DB) (user classes.User, err error) {
	user = classes.User{}
	if db.Where("email = ?", email).First(&user).RecordNotFound() {
		err = fmt.Errorf("User not found")
	}
	return
}
