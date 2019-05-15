package responses

import "github.com/morenocantoj/petstore-go/internal/app/types/classes"

// PetCreatedOK returns a 200 OK with the new pet created
type PetCreatedOK struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Pet     classes.Pet `json:"pet"`
	PetURL  string      `json:"pet_url"`
}

// PetOK returns a 200 OK with the pet found
type PetOK struct {
	Code int32       `json:"code"`
	Pet  classes.Pet `json:"pet"`
}

// PetsOK returns a 201 OK with a list of pets
type PetsOK struct {
	Code         int32        `json:"code"`
	Pets         classes.Pets `json:"pets"`
	CreatePetURL string       `json:"create_pet_url"`
}
