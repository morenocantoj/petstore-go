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

type UserUpdatedOK struct {
	Code     int32  `json:"code"`
	Message  string `json:"message"`
	UsersURL string `json:"users_url"`
}
