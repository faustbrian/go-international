// Package internationalpgx registers international value types with pgx
// without adding a pgx dependency to the core domain packages.
package internationalpgx

import (
	"github.com/faustbrian/go-international/country"
	"github.com/faustbrian/go-international/currency"
	"github.com/faustbrian/go-international/language"
	"github.com/faustbrian/go-international/locale"
	"github.com/faustbrian/go-international/phone"
	"github.com/faustbrian/go-international/postal"
	"github.com/faustbrian/go-international/subdivision"
	"github.com/jackc/pgx/v5/pgtype"
)

// Register maps every scalar value to PostgreSQL text on a caller-owned map.
// Call it before the map is used concurrently.
func Register(typeMap *pgtype.Map) {
	if typeMap == nil {
		return
	}
	for _, value := range []any{
		country.Code{}, country.Alpha3{}, country.Numeric{},
		currency.Code{}, currency.Numeric{}, language.Code{}, locale.Tag{},
		phone.Number{}, phone.CallingCode{}, postal.Code{}, subdivision.Code{},
	} {
		typeMap.RegisterDefaultPgType(value, "text")
	}
}
