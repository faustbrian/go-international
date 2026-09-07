package international_test

import (
	"errors"
	"fmt"
	"testing"

	internationalpostgres "github.com/faustbrian/go-international/adapters/postgres"
	internationalvalidation "github.com/faustbrian/go-international/adapters/validation"
	internationalwire "github.com/faustbrian/go-international/adapters/wire"
	"github.com/faustbrian/go-international/country"
	validation "github.com/faustbrian/go-validation"
	"github.com/faustbrian/go-wire"
	"github.com/jackc/pgx/v5/pgtype"
)

func TestSuccessorAdaptersExposeFrozenEntryPoints(t *testing.T) {
	t.Parallel()

	typeMap := pgtype.NewMap()
	internationalpostgres.Register(typeMap)
	finland, _ := country.Parse("FI")
	if dataType, ok := typeMap.TypeForValue(finland); !ok || dataType == nil || dataType.Name != "text" {
		t.Fatalf("country type = %#v, %v", dataType, ok)
	}

	ctx, err := validation.NewContext(validation.DefaultLimits())
	if err != nil {
		t.Fatal(err)
	}
	if report := internationalvalidation.Country().Validate(ctx, "FI"); report.Err() != nil {
		t.Fatalf("country report = %v", report)
	}

	if _, err := internationalwire.Encode(wire.FormatSOAP, struct{}{}); !errors.Is(err, internationalwire.ErrUnsupportedFormat) {
		t.Fatalf("wire error = %v", err)
	}
}

func Example_successorAdapters() {
	finland, _ := country.Parse("FI")
	typeMap := pgtype.NewMap()
	internationalpostgres.Register(typeMap)
	dataType, _ := typeMap.TypeForValue(finland)

	ctx, _ := validation.NewContext(validation.DefaultLimits())
	report := internationalvalidation.Country().Validate(ctx, "FI")

	_, unsupported := internationalwire.Encode(wire.FormatSOAP, finland)
	fmt.Println(dataType.Name, report.Err(), errors.Is(unsupported, internationalwire.ErrUnsupportedFormat))
	// Output: text <nil> true
}
