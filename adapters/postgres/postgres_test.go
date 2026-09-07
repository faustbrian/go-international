//nolint:staticcheck // This compatibility test intentionally exercises the deprecated facade.
package internationalpostgres_test

//lint:file-ignore SA1019 This compatibility test intentionally exercises the deprecated facade.

import (
	"reflect"
	"testing"

	internationalpostgres "github.com/faustbrian/go-international/adapters/postgres"
	"github.com/faustbrian/go-international/country"
	"github.com/faustbrian/go-international/currency"
	"github.com/faustbrian/go-international/internationalpgx"
	"github.com/faustbrian/go-international/language"
	"github.com/faustbrian/go-international/locale"
	"github.com/faustbrian/go-international/phone"
	"github.com/faustbrian/go-international/postal"
	"github.com/faustbrian/go-international/subdivision"
	"github.com/jackc/pgx/v5/pgtype"
)

func TestRegisterPreservesLegacyTextMappings(t *testing.T) {
	t.Parallel()

	finland, _ := country.Parse("FI")
	alpha3, _ := country.ParseAlpha3("FIN")
	countryNumeric, _ := country.ParseNumeric("246")
	euro, _ := currency.Parse("EUR")
	currencyNumeric, _ := currency.ParseNumeric("978")
	finnish, _ := language.Parse("fi")
	localeTag, _ := locale.Parse("fi-FI")
	number, _ := phone.ParseE164("+16502530000")
	callingCode, _ := phone.ParseCallingCode("+358")
	postalCode, _ := postal.Parse("00100", finland)
	subdivisionCode, _ := subdivision.Parse("FI-18")
	values := []any{finland, alpha3, countryNumeric, euro, currencyNumeric,
		finnish, localeTag, number, callingCode, postalCode, subdivisionCode}

	legacyMap := pgtype.NewMap()
	successorMap := pgtype.NewMap()
	internationalpgx.Register(legacyMap)
	internationalpostgres.Register(successorMap)
	internationalpostgres.Register(nil)
	for _, value := range values {
		legacyType, legacyOK := legacyMap.TypeForValue(value)
		successorType, successorOK := successorMap.TypeForValue(value)
		if !legacyOK || !successorOK || legacyType == nil || successorType == nil ||
			legacyType.Name != "text" || successorType.Name != legacyType.Name {
			t.Fatalf("%s legacy=%#v/%v successor=%#v/%v", reflect.TypeOf(value), legacyType, legacyOK, successorType, successorOK)
		}
		legacyBytes, legacyErr := legacyMap.Encode(pgtype.TextOID, pgtype.TextFormatCode, value, nil)
		successorBytes, successorErr := successorMap.Encode(pgtype.TextOID, pgtype.TextFormatCode, value, nil)
		if legacyErr != nil || successorErr != nil || string(successorBytes) != string(legacyBytes) {
			t.Fatalf("%s legacy=%q/%v successor=%q/%v", reflect.TypeOf(value), legacyBytes, legacyErr, successorBytes, successorErr)
		}
		destination := reflect.New(reflect.TypeOf(value)).Interface()
		if err := successorMap.Scan(pgtype.TextOID, pgtype.TextFormatCode, successorBytes, destination); err != nil || !reflect.DeepEqual(reflect.ValueOf(destination).Elem().Interface(), value) {
			t.Fatalf("%s scan=%#v/%v", reflect.TypeOf(value), destination, err)
		}
	}
}
