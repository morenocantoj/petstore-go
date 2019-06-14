package classes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/morenocantoj/petstore-go/internal/pkg/utils/errors"
)

type Pet struct {
	ID       int64     `gorm:"AUTO_INCREMENT" json:"id"`
	Category string    `json:"category"`
	Name     string    `json:"name"`
	Status   StatusPet `json:"status"`
}

type Pets []Pet

type StatusPet int

const (
	Available StatusPet = iota
	Pending
	Sold
)

var statusDict = map[string]StatusPet{
	"available": Available,
	"pending":   Pending,
	"sold":      Sold,
}

// NewPetFromBody creates a new Pet from a body request
func NewPetFromBody(req *http.Request) Pet {
	// Read body
	body, err := ioutil.ReadAll(req.Body)
	errors.Check(err)
	defer req.Body.Close()

	newPet := Pet{}
	err = json.Unmarshal(body, &newPet)
	errors.Check(err)

	return newPet
}

func NewPetFromFile(name string, category string, status string) Pet {
	return Pet{
		Name:     name,
		Category: category,
		Status:   statusDict[status],
	}
}
