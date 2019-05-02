package main

import (
	"github.com/morenocantoj/petstore-go/internal/app"
	"github.com/morenocantoj/petstore-go/internal/app/database/schema"
)

func main() {
	// Run migrations
	schemaBuilder := schema.SchemaBuilder{}
	schemaBuilder.Build()
	app.Server()
}
