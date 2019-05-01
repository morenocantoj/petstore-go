package classes

type Pet struct {
	Id       int64
	Category string
	Name     string
	Status   StatusPet
}

type Pets []Pet

type StatusPet int

const (
	Available StatusPet = iota
	Pending
	Sold
)
