// Package internationalvalidation exposes pure synchronous Validation rules
// backed by International's strict, bounded parsers.
package internationalvalidation

import (
	"errors"

	international "github.com/faustbrian/go-international"
	"github.com/faustbrian/go-international/country"
	"github.com/faustbrian/go-international/currency"
	"github.com/faustbrian/go-international/language"
	"github.com/faustbrian/go-international/locale"
	"github.com/faustbrian/go-international/phone"
	"github.com/faustbrian/go-international/postal"
	"github.com/faustbrian/go-international/subdivision"
	validation "github.com/faustbrian/go-validation"
)

// Country returns a country rule.
func Country() validation.Validator[string] { return parserRule("country", country.Parse) }

// CountryAlpha3 returns a country alpha-3 rule.
func CountryAlpha3() validation.Validator[string] {
	return parserRule("country_alpha3", country.ParseAlpha3)
}

// CountryNumeric returns a country numeric rule.
func CountryNumeric() validation.Validator[string] {
	return parserRule("country_numeric", country.ParseNumeric)
}

// Subdivision returns a subdivision rule.
func Subdivision() validation.Validator[string] {
	return parserRule("subdivision", subdivision.Parse)
}

// Language returns a language rule.
func Language() validation.Validator[string] { return parserRule("language", language.Parse) }

// Locale returns a locale rule.
func Locale() validation.Validator[string] { return parserRule("locale", locale.Parse) }

// Currency returns a currency rule.
func Currency() validation.Validator[string] { return parserRule("currency", currency.Parse) }

// CurrencyNumeric returns a currency numeric rule.
func CurrencyNumeric() validation.Validator[string] {
	return parserRule("currency_numeric", currency.ParseNumeric)
}

// CallingCode returns a calling-code rule.
func CallingCode() validation.Validator[string] {
	return parserRule("calling_code", phone.ParseCallingCode)
}

// Phone returns a phone parseability rule.
func Phone(options phone.ParseOptions) validation.Validator[string] {
	return parserRule("phone", func(value string) (phone.Number, error) {
		return phone.Parse(value, options)
	})
}

// ValidPhone returns a valid-phone rule.
func ValidPhone(options phone.ParseOptions) validation.Validator[string] {
	return validation.ValidatorFunc[string](func(ctx validation.Context, value string) validation.Report {
		number, err := phone.Parse(value, options)
		if err != nil || !number.Valid() {
			return invalid(ctx, "valid_phone", err)
		}
		return validation.NewReport(ctx.Limits())
	})
}

// Postal returns a postal rule.
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
