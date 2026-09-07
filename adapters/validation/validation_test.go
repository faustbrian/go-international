//nolint:staticcheck // This compatibility test intentionally exercises the deprecated facade.
package internationalvalidation_test

//lint:file-ignore SA1019 This compatibility test intentionally exercises the deprecated facade.

import (
	"errors"
	"strings"
	"testing"

	international "github.com/faustbrian/go-international"
	successor "github.com/faustbrian/go-international/adapters/validation"
	"github.com/faustbrian/go-international/country"
	legacy "github.com/faustbrian/go-international/internationalvalidation"
	"github.com/faustbrian/go-international/phone"
	validation "github.com/faustbrian/go-validation"
)

func TestFactoriesPreserveLegacyReports(t *testing.T) {
	t.Parallel()

	finland, _ := country.Parse("FI")
	ctx, err := validation.NewContext(validation.DefaultLimits())
	if err != nil {
		t.Fatal(err)
	}
	ctx = ctx.WithPath(validation.Field("value"))
	tests := []struct {
		name, valid, invalid string
		legacy, successor    validation.Validator[string]
		cause                error
	}{
		{"country", "FI", "fin", legacy.Country(), successor.Country(), international.ErrInvalid},
		{"country alpha-3", "FIN", "fin", legacy.CountryAlpha3(), successor.CountryAlpha3(), international.ErrInvalid},
		{"country numeric", "246", "24", legacy.CountryNumeric(), successor.CountryNumeric(), international.ErrInvalid},
		{"subdivision", "FI-18", "FI", legacy.Subdivision(), successor.Subdivision(), international.ErrInvalid},
		{"language", "fi", "fin", legacy.Language(), successor.Language(), international.ErrInvalid},
		{"locale", "fi-FI", "fi_ FI", legacy.Locale(), successor.Locale(), international.ErrInvalid},
		{"currency", "EUR", "eur", legacy.Currency(), successor.Currency(), international.ErrInvalid},
		{"currency numeric", "978", "97", legacy.CurrencyNumeric(), successor.CurrencyNumeric(), international.ErrInvalid},
		{"calling code", "+358", "358", legacy.CallingCode(), successor.CallingCode(), international.ErrInvalid},
		{"phone", "040 123 4567", "not a number", legacy.Phone(phone.ParseOptions{RegionHint: finland}), successor.Phone(phone.ParseOptions{RegionHint: finland}), international.ErrInvalid},
		{"valid phone", "+16502530000", "+12001230101", legacy.ValidPhone(phone.ParseOptions{}), successor.ValidPhone(phone.ParseOptions{}), international.ErrInvalid},
		{"postal", "00100", "", legacy.Postal(finland), successor.Postal(finland), international.ErrInvalid},
		{"resource limit", "fi-FI", strings.Repeat("X", 600), legacy.Locale(), successor.Locale(), international.ErrResourceLimit},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if report := test.successor.Validate(ctx, test.valid); report.Err() != nil {
				t.Fatalf("valid report = %v", report)
			}
			legacyReport := test.legacy.Validate(ctx, test.invalid)
			successorReport := test.successor.Validate(ctx, test.invalid)
			legacyViolations := legacyReport.Violations()
			successorViolations := successorReport.Violations()
			if len(legacyViolations) != 1 || len(successorViolations) != 1 {
				t.Fatalf("legacy=%#v successor=%#v", legacyReport, successorReport)
			}
			want, got := legacyViolations[0], successorViolations[0]
			if got.Code() != want.Code() || got.Path().String() != want.Path().String() ||
				got.Severity() != want.Severity() || successorReport.Truncated() != legacyReport.Truncated() ||
				!errors.Is(got.Cause(), test.cause) || !errors.Is(want.Cause(), test.cause) {
				t.Fatalf("legacy=%#v successor=%#v", want, got)
			}
		})
	}
}
