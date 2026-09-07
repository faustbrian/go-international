// Package internationalvalidation integrates strict international parsers with
// validation without introducing validation dependencies into core types.
//
// Deprecated: use github.com/faustbrian/go-international/adapters/validation.
// This package remains supported for the longer of 180 days after successor
// public availability and two subsequently published stable root-module minor
// releases.
package internationalvalidation

import (
	adapter "github.com/faustbrian/go-international/adapters/validation"
	"github.com/faustbrian/go-international/country"
	"github.com/faustbrian/go-international/phone"
	validation "github.com/faustbrian/go-validation"
)

// Country returns a strict current ISO 3166-1 alpha-2 rule.
func Country() validation.Validator[string] { return adapter.Country() }

// CountryAlpha3 returns a strict current ISO 3166-1 alpha-3 rule.
func CountryAlpha3() validation.Validator[string] {
	return adapter.CountryAlpha3()
}

// CountryNumeric returns a strict current ISO 3166-1 numeric rule.
func CountryNumeric() validation.Validator[string] {
	return adapter.CountryNumeric()
}

// Subdivision returns a strict current ISO 3166-2 rule.
func Subdivision() validation.Validator[string] {
	return adapter.Subdivision()
}

// Language returns a strict current canonical ISO 639 rule.
func Language() validation.Validator[string] { return adapter.Language() }

// Locale returns a bounded standards-aware BCP 47 rule.
func Locale() validation.Validator[string] { return adapter.Locale() }

// Currency returns a strict active ISO 4217 alphabetic rule.
func Currency() validation.Validator[string] { return adapter.Currency() }

// CurrencyNumeric returns a strict current ISO 4217 numeric rule.
func CurrencyNumeric() validation.Validator[string] {
	return adapter.CurrencyNumeric()
}

// CallingCode returns a supported plus-prefixed ITU calling-code rule.
func CallingCode() validation.Validator[string] {
	return adapter.CallingCode()
}

// Phone returns a bounded libphonenumber-backed parseability rule. It does not
// claim ownership, reachability, identity, or assignment validity.
func Phone(options phone.ParseOptions) validation.Validator[string] {
	return adapter.Phone(options)
}

// ValidPhone additionally requires current metadata to classify the parsed
// number as valid. This still makes no ownership or reachability claim.
func ValidPhone(options phone.ParseOptions) validation.Validator[string] {
	return adapter.ValidPhone(options)
}

// Postal returns a bounded value rule that preserves the explicit country
// context. It makes no syntax or deliverability claim.
func Postal(context country.Code) validation.Validator[string] {
	return adapter.Postal(context)
}
