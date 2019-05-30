package classes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/morenocantoj/petstore-go/internal/pkg/utils/errors"
)

type User struct {
	ID         int64      `gorm:"AUTO_INCREMENT" json:"id"`
	Username   string     `json:"username"`
	FirstName  string     `json:"first_name"`
	LastName   string     `json:"last_name"`
	Email      string     `gorm:"UNIQUE; NOT NULL" json:"email"`
	Password   string     `gorm:"NOT NULL" json:"password,omitempty"`
	Phone      string     `json:"phone"`
	UserStatus StatusUser `json:"status"`
}

type StatusUser int

const (
	Active StatusUser = iota
	Disabled
)

func NewUserFromBody(req *http.Request) User {
	// Read body
	body, err := ioutil.ReadAll(req.Body)
	errors.Check(err)
	defer req.Body.Close()

	newUser := User{}
	err = json.Unmarshal(body, &newUser)
	errors.Check(err)

	return newUser
}

// sanitizeForJSON prevents password to be serialized to client
func (u *User) SanitizeForJSON() *User {
	u.Password = ""
	return u
}
