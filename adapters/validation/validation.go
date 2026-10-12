// Package internationalvalidation exposes pure synchronous Validation rules
// backed by International's strict, bounded parsers. Rules borrow the Validation
// context for limits and paths but do not poll cancellation or call external IO.
// They emit package-owned violation codes and safe invalid/resource-limit causes,
// not personal input or provider diagnostics. Parsing does not apply business policy.
package internationalvalidation

import (
	"errors"

	international "github.com/faustbrian/go-international/v4"
	"github.com/faustbrian/go-international/v4/country"
	"github.com/faustbrian/go-international/v4/currency"
	"github.com/faustbrian/go-international/v4/language"
	"github.com/faustbrian/go-international/v4/locale"
	"github.com/faustbrian/go-international/v4/phone"
	"github.com/faustbrian/go-international/v4/postal"
	"github.com/faustbrian/go-international/v4/subdivision"
	validation "github.com/faustbrian/go-validation/v2"
)

// Country returns a strict current alpha-2 rule.
func Country() validation.Validator[string] { return parserRule("country", country.Parse) }

// CountryAlpha3 returns a strict current alpha-3 rule.
func CountryAlpha3() validation.Validator[string] {
	return parserRule("country_alpha3", country.ParseAlpha3)
}

// CountryNumeric returns a strict current numeric country rule.
func CountryNumeric() validation.Validator[string] {
	return parserRule("country_numeric", country.ParseNumeric)
}

// Subdivision returns a strict current subdivision rule.
func Subdivision() validation.Validator[string] {
	return parserRule("subdivision", subdivision.Parse)
}

// Language returns a canonical language rule.
func Language() validation.Validator[string] { return parserRule("language", language.Parse) }

// Locale returns a bounded provider-backed locale rule.
func Locale() validation.Validator[string] { return parserRule("locale", locale.Parse) }

// Currency returns a strict active currency rule.
func Currency() validation.Validator[string] { return parserRule("currency", currency.Parse) }

// CurrencyNumeric returns a strict active numeric currency rule.
func CurrencyNumeric() validation.Validator[string] {
	return parserRule("currency_numeric", currency.ParseNumeric)
}

// CallingCode returns a supported calling-code rule.
func CallingCode() validation.Validator[string] {
	return parserRule("calling_code", phone.ParseCallingCode)
}

// Phone snapshots options and checks parseability, not Possible or Valid metadata.
func Phone(options phone.ParseOptions) validation.Validator[string] {
	return parserRule("phone", func(value string) (phone.Number, error) {
		return phone.Parse(value, options)
	})
}

// ValidPhone snapshots options and requires provider Valid metadata, not ownership
// or reachability. A false validity result uses the safe invalid cause.
func ValidPhone(options phone.ParseOptions) validation.Validator[string] {
	return validation.ValidatorFunc[string](func(ctx validation.Context, value string) validation.Report {
		number, err := phone.Parse(value, options)
		if err != nil || !number.Valid() {
			return invalid(ctx, "valid_phone", err)
		}
		return validation.NewReport(ctx.Limits())
	})
}

// Postal returns a bounded opaque postal-value rule.
func Postal(context country.Code) validation.Validator[string] {
	return parserRule("postal", func(value string) (postal.Code, error) {
		return postal.Parse(value, context)
	})
}

func parserRule[T any](code string, parse func(string) (T, error)) validation.Validator[string] {
	return validation.ValidatorFunc[string](func(ctx validation.Context, value string) validation.Report {
		if _, err := parse(value); err != nil {
			return invalid(ctx, code, err)
		}
		return validation.NewReport(ctx.Limits())
	})
}

func invalid(ctx validation.Context, code string, cause error) validation.Report {
	safeCause := international.ErrInvalid
	if errors.Is(cause, international.ErrResourceLimit) {
		safeCause = international.ErrResourceLimit
	}
	return validation.NewReport(ctx.Limits()).Add(validation.NewViolation(
		ctx.Path(), code, validation.Error, nil, safeCause,
	))
}
