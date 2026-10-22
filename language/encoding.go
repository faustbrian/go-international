package language

import (
	"database/sql/driver"

	"github.com/faustbrian/go-international/v5/internal/codec"
)

// MarshalText encodes the stored spelling without canonicalizing it, rejecting
// absence. A deprecated spelling returned by ParseISO3 encodes successfully but
// is rejected by UnmarshalText; see Code's conversion limitation.
func (code Code) MarshalText() ([]byte, error) { return codec.MarshalText(code.value, "language code") }

// UnmarshalText parses strict current canonical ISO 639 text.
func (code *Code) UnmarshalText(input []byte) error {
	parsed, err := Parse(string(input))
	if err == nil {
		*code = parsed
	}
	return err
}

// MarshalJSON encodes the stored spelling as a string, or absence as null.
// It does not canonicalize deprecated ParseISO3 results; UnmarshalJSON rejects
// those spellings even when this encoding succeeds.
func (code Code) MarshalJSON() ([]byte, error) { return codec.MarshalJSON(code.value) }

// UnmarshalJSON accepts only a current canonical ISO 639 string or null.
func (code *Code) UnmarshalJSON(input []byte) error {
	parsed, absent, err := codec.DecodeJSON(input, "language code", Parse)
	if err == nil {
		if absent {
			*code = Code{}
		} else {
			*code = parsed
		}
	}
	return err
}

// Value returns stored spelling, or SQL NULL for absence, without canonicalizing
// deprecated ParseISO3 results. Such stored text is rejected by strict Scan.
func (code Code) Value() (driver.Value, error) { return codec.DatabaseValue(code.value) }

// Scan accepts SQL NULL, string, or byte text.
func (code *Code) Scan(source any) error {
	parsed, absent, err := codec.Scan(source, "language code", Parse)
	if err == nil {
		if absent {
			*code = Code{}
		} else {
			*code = parsed
		}
	}
	return err
}
