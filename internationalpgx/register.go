// Package internationalpgx registers international value types with pgx
// without adding a pgx dependency to the core domain packages.
//
// Deprecated: use github.com/faustbrian/go-international/adapters/postgres.
// This package remains supported for the longer of 180 days after successor
// public availability and two subsequently published stable root-module minor
// releases.
package internationalpgx

import (
	internationalpostgres "github.com/faustbrian/go-international/adapters/postgres"
	"github.com/jackc/pgx/v5/pgtype"
)

// Register maps every scalar value to PostgreSQL text on a caller-owned map.
// Call it before the map is used concurrently.
func Register(typeMap *pgtype.Map) {
	internationalpostgres.Register(typeMap)
}
