package homepage

import "embed"

//go:embed "assets"
var Files embed.FS

//go:embed "sql/schema.sql"
var SQLFiles embed.FS

//go:embed "dev.db"
var SeedDB embed.FS
