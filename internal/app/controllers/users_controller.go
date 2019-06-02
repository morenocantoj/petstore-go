package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/morenocantoj/petstore-go/internal/app/database"
	"github.com/morenocantoj/petstore-go/internal/app/types/classes"
	"github.com/morenocantoj/petstore-go/internal/app/types/responses"
)

// PetsController Defines a user routes controller
type UsersController struct {
	BaseController
}

// Create a new user and encode its password
func (u *UsersController) Create(writter http.ResponseWriter, req *http.Request) {
	fmt.Println("POST /signup")
	newUser := classes.NewUserFromBody(req)

	db := database.Connector{Connection: database.ConnectToDatabase()}
	defer db.Connection.Close()

	err := db.Connection.Create(&newUser).Error
	if err != nil {
		response := responses.BadRequest{HttpError: responses.HttpError{Code: 400, Message: "Invalid email or password"}}
		u.WriteResponse(response, writter, http.StatusBadRequest)

	} else {
		userCreatedResponse := responses.UserCreatedOK{
			Code:    201,
			Message: "User created successfully",
			User:    newUser.SanitizeForJSON(),
			UserURL: req.Host + "/users/" + strconv.FormatInt(newUser.ID, 10),
		}
		u.WriteResponse(userCreatedResponse, writter, http.StatusCreated)
	}
}
