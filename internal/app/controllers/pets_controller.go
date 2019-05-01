package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/morenocantoj/petstore-go/internal/app/types/classes"
	"github.com/morenocantoj/petstore-go/internal/pkg/utils/errors"
)

type PetsController struct{}

func (c *PetsController) Index(w http.ResponseWriter, r *http.Request) {
	mockPet := classes.Pet{
		Id:       1,
		Category: "dog",
		Name:     "Rudolf",
		Status:   classes.Available,
	}
	mockPets := classes.Pets{mockPet}
	responseJson, err := json.Marshal(&mockPets)
	errors.Check(err)

	w.Write(responseJson)
}
