// Package internationalpostgres registers International scalar values with
// caller-owned pgx type maps without connection I/O or retained resources.
package internationalpostgres

import (
	"github.com/faustbrian/go-international/v4/country"
	"github.com/faustbrian/go-international/v4/currency"
	"github.com/faustbrian/go-international/v4/language"
	"github.com/faustbrian/go-international/v4/locale"
	"github.com/faustbrian/go-international/v4/phone"
	"github.com/faustbrian/go-international/v4/postal"
	"github.com/faustbrian/go-international/v4/subdivision"
	"github.com/jackc/pgx/v5/pgtype"
)

// Register registers all eleven International scalar types as PostgreSQL text.
// A nil map is a no-op. Call Register before using the map concurrently.
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
