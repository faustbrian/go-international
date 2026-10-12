// Package language provides ISO 639 identifiers through pinned x/text registry
// data. Use Parse for the canonical lowercase identity; ParseISO3 is a separate
// provider conversion with deprecated-spelling limitations described below.
// Name supplies presentation metadata, not negotiation or locale detection.
package language

import (
	international "github.com/faustbrian/go-international/v4"
	textlanguage "golang.org/x/text/language"
	"golang.org/x/text/language/display"
)

// Code stores an ISO 639 language identifier; its zero value is absent. Parse
// produces canonical spelling, preferring two letters where a mapping exists.
// ParseISO3 can instead produce a deprecated provider spelling that Parse rejects.
// Encoders preserve that spelling rather than canonicalizing it, so a successful
// encode need not round-trip through strict text, JSON or SQL decoding.
type Code struct{ value string }

// Parse accepts a canonical lowercase ISO identifier. It requires ISO 639-1
// when a two-letter representation exists and otherwise retains ISO 639-3.
func Parse(input string) (Code, error) {
	if !validLowerAlpha(input, 2) && !validLowerAlpha(input, 3) {
		return Code{}, international.NewParseError("language code", "invalid ISO 639 identifier")
	}
	base, err := textlanguage.ParseBase(input)
	canonical, _ := textlanguage.DeprecatedBase.Canonicalize(textlanguage.Raw.Make(input))
	if err != nil {
		return Code{}, international.NewParseError("language code", "unknown or obsolete identifier")
	}
	if base.String() != input {
		return Code{}, international.NewParseError("language code", "unknown or obsolete identifier")
	}
	if canonical.String() != input {
		return Code{}, international.NewParseError("language code", "unknown or obsolete identifier")
	}
	return Code{value: input}, nil
}

// ParseISO3 accepts a lowercase three-letter identifier recognized by the pinned
// x/text ParseBase provider and returns Base.String spelling. Unlike Parse, it
// does not additionally reject DeprecatedBase canonicalization differences: for
// example, "heb" yields "iw", which Parse and default persistence decoding reject.
// This is an existing provider-conversion limitation, not a canonicality guarantee.
func ParseISO3(input string) (Code, error) {
	if !validLowerAlpha(input, 3) {
		return Code{}, international.NewParseError("language code", "invalid ISO 639 three-letter identifier")
	}
	base, err := textlanguage.ParseBase(input)
	if err != nil {
		return Code{}, international.NewParseError("language code", "unknown or obsolete identifier")
	}
	return Code{value: base.String()}, nil
}

// String returns the stored identifier spelling, or empty for absence. ParseISO3
// may have stored a deprecated spelling rather than Parse's canonical identity.
func (code Code) String() string { return code.value }

// IsZero reports whether the code represents an absent value.
func (code Code) IsZero() bool { return code.value == "" }

// ISO3 returns the authoritative three-letter mapping, or an empty string when
// the code is absent.
func (code Code) ISO3() string {
	if code.IsZero() {
		return ""
	}
	base, _ := textlanguage.ParseBase(code.value)
	return base.ISO3()
}

// Name returns CLDR display metadata in the requested language.
func Name(code Code, displayLanguage textlanguage.Tag) string {
	if code.IsZero() {
		return ""
	}
	namer := display.Languages(displayLanguage)
	if namer == nil {
		return ""
	}
	return namer.Name(textlanguage.Make(code.value))
}

func validLowerAlpha(value string, length int) bool {
	if len(value) != length {
		return false
	}
	for index := range value {
		if value[index] < 'a' || value[index] > 'z' {
			return false
		}
	}
	return true
}
