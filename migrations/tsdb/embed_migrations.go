package tsdb

import (
	"embed"

	"boilerplate/internal/infrastructure/repository/tsdb"
)

//go:embed *.sql
var EmbedMigrations embed.FS

var CurrentRepo *tsdb.Repository
