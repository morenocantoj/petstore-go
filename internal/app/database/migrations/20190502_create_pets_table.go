package migrations

import (
	"github.com/morenocantoj/petstore-go/internal/app/database"
	"github.com/morenocantoj/petstore-go/internal/app/types/classes"
)

func Run() {
	db := database.Connector{Connection: database.ConnectToDatabase()}
	defer db.Connection.Close()
	db.Connection.AutoMigrate(&classes.Pet{})
}
