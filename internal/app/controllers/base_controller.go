package controllers

import "net/http"

// BaseController is the base struct from subcontrollers make use
type BaseController struct{}

func (b *BaseController) writeResponse(jsonData []byte, w http.ResponseWriter, h int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(h)
	w.Write(jsonData)
}
