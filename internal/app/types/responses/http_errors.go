package responses

type BadRequest struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type NotFound struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
