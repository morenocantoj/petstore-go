package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/morenocantoj/petstore-go/internal/app/database"
	"github.com/morenocantoj/petstore-go/internal/app/types/classes"
	"github.com/morenocantoj/petstore-go/internal/pkg/utils/errors"
)

type PetsController struct{}

func (c *PetsController) Index(writter http.ResponseWriter, req *http.Request) {
	db := database.Connector{Connection: database.ConnectToDatabase()}
	defer db.Connection.Close()

	pets := db.Connection.Find(&classes.Pets{}).Value

	responseJson, err := json.Marshal(&pets)
	errors.Check(err)
	writter.Write(responseJson)
}
