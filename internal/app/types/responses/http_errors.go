package responses

type HttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type BadRequest struct {
	HttpError
}

type NotFound struct {
	HttpError
}

type ServerError struct {
	HttpError
}

type Forbidden struct {
	HttpError
}
