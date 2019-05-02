package schema

import "github.com/morenocantoj/petstore-go/internal/app/database/migrations"

type SchemaBuilder struct{}

func (s *SchemaBuilder) Build() {
	migrations.Run()
}
