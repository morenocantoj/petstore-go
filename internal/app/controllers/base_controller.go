package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/morenocantoj/petstore-go/internal/pkg/utils/errors"
)

// BaseController is the base struct from subcontrollers make use
type BaseController struct{}

func (b *BaseController) WriteResponse(data interface{}, w http.ResponseWriter, h int) {
	responseJSON, err := json.Marshal(&data)
	errors.Check(err)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(h)
	w.Write(responseJSON)
}
