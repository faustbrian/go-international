package internationalvalidation_test

import (
	"errors"
	"strings"
	"testing"

	international "github.com/faustbrian/go-international"
	"github.com/faustbrian/go-international/country"
	"github.com/faustbrian/go-international/internationalvalidation"
	"github.com/faustbrian/go-international/phone"
	validation "github.com/faustbrian/go-validation"
)

func TestRulesDelegateToDistinctStrictParsers(t *testing.T) {
	t.Parallel()
	finland, err := country.Parse("FI")
	if err != nil {
		t.Fatal(err)
	}
	ctx, err := validation.NewContext(validation.DefaultLimits())
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name, valid, invalid string
		rule                 validation.Validator[string]
	}{
		{"country", "FI", "fin", internationalvalidation.Country()},
		{"country alpha-3", "FIN", "fin", internationalvalidation.CountryAlpha3()},
		{"country numeric", "246", "24", internationalvalidation.CountryNumeric()},
		{"subdivision", "FI-18", "FI", internationalvalidation.Subdivision()},
		{"language", "fi", "fin", internationalvalidation.Language()},
		{"locale", "fi-FI", "fi_ FI", internationalvalidation.Locale()},
		{"currency", "EUR", "eur", internationalvalidation.Currency()},
		{"currency numeric", "978", "97", internationalvalidation.CurrencyNumeric()},
		{"calling code", "+358", "358", internationalvalidation.CallingCode()},
		{"phone", "040 123 4567", "not a number", internationalvalidation.Phone(phone.ParseOptions{RegionHint: finland})},
		{"valid phone", "+16502530000", "+12001230101", internationalvalidation.ValidPhone(phone.ParseOptions{})},
		{"postal", "00100", "", internationalvalidation.Postal(finland)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if report := test.rule.Validate(ctx, test.valid); report.Err() != nil {
				t.Fatalf("valid report = %v", report)
			}
			if report := test.rule.Validate(ctx, test.invalid); report.Err() == nil {
				t.Fatal("invalid value passed")
			}
		})
	}
}

func TestRulesPreserveResourceLimitAsSafeCause(t *testing.T) {
	t.Parallel()
	ctx, err := validation.NewContext(validation.DefaultLimits())
	if err != nil {
		t.Fatal(err)
	}
	report := internationalvalidation.Locale().Validate(ctx, strings.Repeat("X", 600))
	violations := report.Violations()
	if len(violations) != 1 || !errors.Is(violations[0].Cause(), international.ErrResourceLimit) {
		t.Fatalf("violations = %#v", violations)
	}
}

func TestRulesPreserveReleasedDiagnosticContract(t *testing.T) {
	t.Parallel()

	finland, err := country.Parse("FI")
	if err != nil {
		t.Fatal(err)
	}
	ctx, err := validation.NewContext(validation.DefaultLimits())
	if err != nil {
		t.Fatal(err)
	}
	ctx = ctx.WithPath(validation.Field("value"))
	tests := []struct {
		name, code, invalid string
		rule                validation.Validator[string]
		cause               error
	}{
		{"country", "country", "fin", internationalvalidation.Country(), international.ErrInvalid},
		{"country alpha-3", "country_alpha3", "fin", internationalvalidation.CountryAlpha3(), international.ErrInvalid},
		{"country numeric", "country_numeric", "24", internationalvalidation.CountryNumeric(), international.ErrInvalid},
		{"subdivision", "subdivision", "FI", internationalvalidation.Subdivision(), international.ErrInvalid},
		{"language", "language", "fin", internationalvalidation.Language(), international.ErrInvalid},
		{"locale", "locale", "fi_ FI", internationalvalidation.Locale(), international.ErrInvalid},
		{"currency", "currency", "eur", internationalvalidation.Currency(), international.ErrInvalid},
		{"currency numeric", "currency_numeric", "97", internationalvalidation.CurrencyNumeric(), international.ErrInvalid},
		{"calling code", "calling_code", "358", internationalvalidation.CallingCode(), international.ErrInvalid},
		{"phone", "phone", "not a number", internationalvalidation.Phone(phone.ParseOptions{RegionHint: finland}), international.ErrInvalid},
		{"valid phone", "valid_phone", "+12001230101", internationalvalidation.ValidPhone(phone.ParseOptions{}), international.ErrInvalid},
		{"postal", "postal", "", internationalvalidation.Postal(finland), international.ErrInvalid},
		{"resource limit", "locale", strings.Repeat("X", 600), internationalvalidation.Locale(), international.ErrResourceLimit},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			report := test.rule.Validate(ctx, test.invalid)
			violations := report.Violations()
			if len(violations) != 1 || report.Truncated() {
				t.Fatalf("report = %#v", report)
			}
			violation := violations[0]
			if violation.Code() != test.code || violation.Path().String() != "value" ||
				violation.Severity() != validation.Error || !errors.Is(violation.Cause(), test.cause) {
				t.Fatalf("violation = code %q, path %q, severity %v, cause %v", violation.Code(), violation.Path(), violation.Severity(), violation.Cause())
			}
		})
	}
}
