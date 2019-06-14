package manager

import (
	"encoding/csv"
	"mime/multipart"

	"github.com/morenocantoj/petstore-go/internal/app/types/classes"
)

type PetsFile interface {
	StorePets() (classes.Pets, error)
}

type CSVPetsFile struct{}

func (fm *CSVPetsFile) StorePets(file multipart.File) (classes.Pets, error) {
	pets := classes.Pets{}
	existingPets := make(map[classes.Pet]bool)
	reader := csv.NewReader(file)
	// reader.Comma = ';'
	lineCounter := 0
	for {
		line, err := reader.Read()
		if err != nil {
			return pets, err
		}
		// Don't read CSV header
		if lineCounter != 0 {
			pet := classes.NewPetFromFile(line[0], line[1], line[2])
			// Check if we already have the same pet at the same category
			if _, ok := existingPets[pet]; !ok {
				// Insert non duplicate pet
				pets = append(pets, pet)
				existingPets[pet] = true
			}
		}
		lineCounter++
	}
}
