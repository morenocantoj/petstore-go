package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/morenocantoj/petstore-go/internal/pkg/utils/errors"
)

type Connector struct {
	Connection *gorm.DB
}

func ConnectToDatabase() *gorm.DB {
	db, err := gorm.Open("postgres",
		"host=127.0.0.1 port=5432 user=petstore-go dbname=petstore-go password=petstore-go sslmode=disable")
	errors.Check(err)
	return db
}
