package responses

import (
	"github.com/morenocantoj/petstore-go/internal/app/types/classes"
)

type UserCreatedOK struct {
	Code    int32         `json:"code"`
	Message string        `json:"message"`
	User    *classes.User `json:"user"`
	UserURL string        `json:"user_url"`
}
