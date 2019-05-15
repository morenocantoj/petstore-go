package responses

import "github.com/morenocantoj/petstore-go/internal/app/types/classes"

// PetCreatedOK returns a 201 OK with the new pet created
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

// PetDestroyedOK returns a 200 OK if the pet is deleted successfully
type PetDestroyedOK struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	PetsURL string `json:"pets_url"`
}

// PetsOK returns a 200 OK with a list of pets
type PetsOK struct {
	Code         int32        `json:"code"`
	Pets         classes.Pets `json:"pets"`
	CreatePetURL string       `json:"create_pet_url"`
}
