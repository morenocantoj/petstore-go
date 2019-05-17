package classes

type User struct {
	ID         int64      `gorm:"AUTO_INCREMENT" json:"id"`
	Username   string     `json:"username"`
	FirstName  string     `json:"first_name"`
	LastName   string     `json:"last_name"`
	Email      string     `gorm:"UNIQUE; NOT NULL" json:"email"`
	Password   string     `gorm:"NOT NULL" json:"password"`
	Phone      string     `json:"phone"`
	UserStatus StatusUser `json:"status"`
}

type StatusUser int

const (
	Active StatusUser = iota
	Disabled
)
