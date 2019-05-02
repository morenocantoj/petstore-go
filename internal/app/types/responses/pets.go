package responses

import "github.com/morenocantoj/petstore-go/internal/app/types/classes"

type PetCreatedOK struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Pet     classes.Pet `json:"pet"`
	PetURL  string      `json:"pet_url"`
}
