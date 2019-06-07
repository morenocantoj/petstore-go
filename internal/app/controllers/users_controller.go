package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
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

// Update a specific user without updating password
func (u *UsersController) Update(writter http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userID := params["id"]
	fmt.Printf("PATCH /users/%s\n", userID)

	user := classes.NewUserFromBody(req)
	user.ID, _ = strconv.ParseInt(userID, 10, 64)
	user.Password = ""
	db := database.Connector{Connection: database.ConnectToDatabase()}
	defer db.Connection.Close()

	if db.Connection.Model(&user).Where("id = ?", user.ID).Update(&user).RowsAffected <= 0 {
		response := responses.NotFound{HttpError: responses.HttpError{Code: 404, Message: "User not found"}}
		u.WriteResponse(response, writter, http.StatusNotFound)
		return
	}
	response := responses.UserUpdatedOK{
		Code:     200,
		Message:  "User updated successfully",
		UsersURL: req.Host + "/users",
	}
	u.WriteResponse(response, writter, http.StatusOK)
}

// Modify a specific user without updating password
func (u *UsersController) Modify(writter http.ResponseWriter, req *http.Request) {
	email := context.Get(req, "user")
	fmt.Println("PATCH /users/me")
	if email == nil {
		response := responses.Forbidden{HttpError: responses.HttpError{Code: 403, Message: "You must provide a valid token"}}
		u.WriteResponse(response, writter, http.StatusInternalServerError)
		return
	}

	user := classes.NewUserFromBody(req)
	user.Email = email.(string)
	user.Password = ""
	db := database.Connector{Connection: database.ConnectToDatabase()}
	defer db.Connection.Close()

	if db.Connection.Model(&user).Where("email = ?", user.Email).Update(&user).RowsAffected <= 0 {
		response := responses.NotFound{HttpError: responses.HttpError{Code: 404, Message: "User not found"}}
		u.WriteResponse(response, writter, http.StatusNotFound)
		return
	}
	response := responses.UserUpdatedOK{
		Code:     200,
		Message:  "User updated successfully",
		UsersURL: req.Host + "/users",
	}
	u.WriteResponse(response, writter, http.StatusOK)
}
