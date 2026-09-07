package internationalpgx_test

import (
	"reflect"
	"testing"

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

func TestRegisterSupportsPgxTextEncodeAndScan(t *testing.T) {
	t.Parallel()
	typeMap := pgtype.NewMap()
	internationalpgx.Register(typeMap)
	internationalpgx.Register(nil)

	finland, err := country.Parse("FI")
	if err != nil {
		t.Fatal(err)
	}
	dataType, ok := typeMap.TypeForValue(finland)
	if !ok || dataType == nil || dataType.Name != "text" {
		t.Fatalf("type = %#v, %v", dataType, ok)
	}
	encoded, err := typeMap.Encode(pgtype.TextOID, pgtype.TextFormatCode, finland, nil)
	if err != nil || string(encoded) != "FI" {
		t.Fatalf("Encode() = %q, %v", encoded, err)
	}
	var decoded country.Code
	if err := typeMap.Scan(pgtype.TextOID, pgtype.TextFormatCode, encoded, &decoded); err != nil || decoded != finland {
		t.Fatalf("Scan() = %q, %v", decoded, err)
	}

	number, err := phone.ParseE164("+16502530000")
	if err != nil {
		t.Fatal(err)
	}
	encoded, err = typeMap.Encode(pgtype.TextOID, pgtype.TextFormatCode, number, nil)
	if err != nil || string(encoded) != number.E164() {
		t.Fatalf("phone Encode() = %q, %v", encoded, err)
	}
}

func TestRegisterMapsEveryReleasedScalarTypeToText(t *testing.T) {
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

	typeMap := pgtype.NewMap()
	internationalpgx.Register(typeMap)
	for _, value := range []any{
		finland, alpha3, countryNumeric, euro, currencyNumeric, finnish,
		localeTag, number, callingCode, postalCode, subdivisionCode,
	} {
		dataType, ok := typeMap.TypeForValue(value)
		if !ok || dataType == nil || dataType.Name != "text" {
			t.Fatalf("%s maps to %#v, %v", reflect.TypeOf(value), dataType, ok)
		}
	}
}
