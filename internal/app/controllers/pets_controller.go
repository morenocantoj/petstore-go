package controllers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/morenocantoj/petstore-go/internal/app/core/manager"
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
	c.WriteResponse(petsResponse, writter, http.StatusOK)
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
			response := responses.NotFound{HttpError: responses.HttpError{Code: 404, Message: "Pet not found"}}
			c.WriteResponse(response, writter, http.StatusNotFound)

		} else {
			response := responses.PetOK{Code: 200, Pet: pet}
			c.WriteResponse(response, writter, http.StatusOK)
		}

	} else {
		response := responses.BadRequest{HttpError: responses.HttpError{Code: 400, Message: "Invalid ID"}}
		c.WriteResponse(response, writter, http.StatusBadRequest)
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
	c.WriteResponse(petCreatedResponse, writter, http.StatusOK)
}

// Destroy an existing pet
func (c *PetsController) Destroy(writter http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	petID := params["id"]
	fmt.Printf("DELETE /pets/%s\n", petID)

	if petID != "" {
		db := database.Connector{Connection: database.ConnectToDatabase()}
		defer db.Connection.Close()

		petID, _ := strconv.ParseInt(petID, 10, 64)
		pet := classes.Pet{ID: petID}

		db.Connection.Delete(&pet)

		petDestroyedResponse := responses.PetDestroyedOK{
			Code:    200,
			Message: "Pet deleted succesfully",
			PetsURL: req.Host + "/pets",
		}
		c.WriteResponse(petDestroyedResponse, writter, http.StatusOK)

	} else {
		response := responses.BadRequest{HttpError: responses.HttpError{Code: 400, Message: "Invalid ID"}}
		c.WriteResponse(response, writter, http.StatusBadRequest)
	}
}

// Update a existing pet
func (c *PetsController) Update(writter http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	petID := params["id"]
	fmt.Printf("PATCH /pets/%s\n", petID)

	if petID != "" {
		pet := classes.NewPetFromBody(req)
		pet.ID, _ = strconv.ParseInt(petID, 10, 64)
		db := database.Connector{Connection: database.ConnectToDatabase()}
		defer db.Connection.Close()

		if db.Connection.Model(&pet).Where("id = ?", pet.ID).Update(&pet).RowsAffected > 0 {
			response := responses.PetUpdatedOK{
				Code:    200,
				Message: "Pet updated successfully",
				PetsURL: req.Host + "/pets",
			}
			c.WriteResponse(response, writter, http.StatusOK)

		} else {
			response := responses.NotFound{HttpError: responses.HttpError{Code: 404, Message: "Pet not found"}}
			c.WriteResponse(response, writter, http.StatusNotFound)
		}

	} else {
		response := responses.BadRequest{HttpError: responses.HttpError{Code: 400, Message: "Invalid ID"}}
		c.WriteResponse(response, writter, http.StatusBadRequest)
	}
}

// Upload a CSV file with pets and classificates them
func (c *PetsController) Upload(writter http.ResponseWriter, req *http.Request) {
	fmt.Println("POST /pets/upload")

	// File has to have 10 MB as much
	req.ParseMultipartForm(10 << 20)
	file, _, err := req.FormFile("pets_file")
	if err != nil {
		response := responses.ServerError{HttpError: responses.HttpError{Code: 500, Message: "Error loading file"}}
		c.WriteResponse(response, writter, http.StatusInternalServerError)
		return
	}

	defer file.Close()

	csvManager := manager.CSVPetsFile{}
	pets, err := csvManager.StorePets(file)
	if err != nil && err != io.EOF {
		response := responses.ServerError{HttpError: responses.HttpError{Code: 500, Message: "Error reading file"}}
		c.WriteResponse(response, writter, http.StatusInternalServerError)
		return
	}

	db := database.Connector{Connection: database.ConnectToDatabase()}
	defer db.Connection.Close()

	for i := range pets {
		err = db.Connection.Create(&pets[i]).Error
		if err != nil {
			response := responses.ServerError{HttpError: responses.HttpError{Code: 500, Message: "Error creating pets"}}
			c.WriteResponse(response, writter, http.StatusInternalServerError)
			return
		}
	}

	petsResponse := responses.PetsOK{
		Code:         200,
		Pets:         pets,
		CreatePetURL: req.Host + "/pets",
	}
	c.WriteResponse(petsResponse, writter, http.StatusOK)
}
