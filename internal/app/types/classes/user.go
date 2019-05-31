package classes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/morenocantoj/petstore-go/internal/pkg/utils/errors"
	"golang.org/x/crypto/bcrypt"
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

// BeforeCreate hook to hash user's password before inserting into database
func (u *User) BeforeCreate() (err error) {
	u.Password, err = u.hashPassword()
	return
}

func (u *User) hashPassword() (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	return string(bytes), err
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
