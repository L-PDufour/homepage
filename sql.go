// sql.go
package homepage

import "embed"

//go:embed "sql/schema.sql" "sql/seed.sql"
var SQLFiles embed.FS
