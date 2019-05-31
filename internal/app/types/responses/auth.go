package responses

type AuthResponseOK struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Token   string `json:"token"`
}
