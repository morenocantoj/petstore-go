package classes

type pet struct {
	id       int64
	category string
	name     string
	status   status
}

type status int

const (
	available status = iota
	pending
	sold
)
