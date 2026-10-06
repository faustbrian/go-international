//nolint:staticcheck // This compatibility test intentionally exercises deprecated facades.
package international_test

//lint:file-ignore SA1019 This compatibility test intentionally exercises deprecated facades.

import (
	"errors"
	"testing"

	"github.com/faustbrian/go-international/v4/internationalpgx"
	"github.com/faustbrian/go-international/v4/internationalvalidation"
	"github.com/faustbrian/go-international/v4/internationalwire"
	validation "github.com/faustbrian/go-validation/v2"
	"github.com/faustbrian/go-wire/v3"
	"github.com/jackc/pgx/v5/pgtype"
)

func TestLegacyAndSuccessorAdaptersCoexist(t *testing.T) {
	t.Parallel()

	internationalpgx.Register(pgtype.NewMap())
	ctx, err := validation.NewContext(validation.DefaultLimits())
	if err != nil {
		t.Fatal(err)
	}
	if report := internationalvalidation.Country().Validate(ctx, "FI"); report.Err() != nil {
		t.Fatalf("legacy country report = %v", report)
	}
	if _, err := internationalwire.Encode(wire.FormatSOAP, struct{}{}); !errors.Is(err, internationalwire.ErrUnsupportedFormat) {
		t.Fatalf("legacy wire error = %v", err)
	}
}
