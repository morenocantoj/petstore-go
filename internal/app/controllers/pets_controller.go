package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/morenocantoj/petstore-go/internal/app/database"
	"github.com/morenocantoj/petstore-go/internal/app/types/classes"
	"github.com/morenocantoj/petstore-go/internal/app/types/responses"
	"github.com/morenocantoj/petstore-go/internal/pkg/utils/errors"
)

// PetsController Defines a main base controller
type PetsController struct {
	BaseController
}

// Index Get all pets stored in database
func (c *PetsController) Index(writter http.ResponseWriter, req *http.Request) {
	fmt.Println("GET /pets")
	db := database.Connector{Connection: database.ConnectToDatabase()}
	defer db.Connection.Close()

	pets := classes.Pets{}
	db.Connection.Find(&pets)
	petsResponse := responses.PetsOK{
		Code:         200,
		Pets:         pets,
		CreatePetURL: req.Host + "/pets",
	}

	responseJSON, err := json.Marshal(&petsResponse)
	errors.Check(err)
	c.writeResponse(responseJSON, writter, http.StatusOK)
}

// Show returns a specific pet stored in database
func (c *PetsController) Show(writter http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	petID := params["id"]
	fmt.Printf("GET /pets/%s\n", petID)

	if petID != "" {
		pet := classes.Pet{}
		db := database.Connector{Connection: database.ConnectToDatabase()}
		defer db.Connection.Close()

		if db.Connection.Where("id = ?", petID).First(&pet).RecordNotFound() {
			response := responses.NotFound{Code: 404, Message: "La mascota que buscas no se encuentra"}
			responseJSON, err := json.Marshal(&response)
			errors.Check(err)
			c.writeResponse(responseJSON, writter, http.StatusNotFound)

		} else {
			response := responses.PetOK{Code: 200, Pet: pet}
			responseJSON, err := json.Marshal(&response)
			errors.Check(err)
			c.writeResponse(responseJSON, writter, http.StatusOK)
		}

	} else {
		response := responses.BadRequest{Code: 400, Message: "El id introducido no es válido"}
		responseJSON, err := json.Marshal(&response)
		errors.Check(err)
		c.writeResponse(responseJSON, writter, http.StatusBadRequest)
	}
}

// Create Creates a new pet and stores it
func (c *PetsController) Create(writter http.ResponseWriter, req *http.Request) {
	fmt.Println("POST /pets")
	newPet := classes.NewPetFromBody(req)

	db := database.Connector{Connection: database.ConnectToDatabase()}
	defer db.Connection.Close()

	err := db.Connection.Create(&newPet).Error
	errors.Check(err)

	petCreatedResponse := responses.PetCreatedOK{
		Code:    200,
		Message: "Pet created successfully",
		Pet:     newPet,
		PetURL:  req.Host + "/pets/" + strconv.FormatInt(newPet.ID, 10),
	}

	responseJSON, err := json.Marshal(&petCreatedResponse)
	errors.Check(err)
	c.writeResponse(responseJSON, writter, http.StatusOK)
}
