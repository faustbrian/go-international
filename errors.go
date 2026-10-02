package international

import (
	"errors"
	"fmt"
)

var (
	// ErrInvalid identifies syntactically invalid input without retaining it.
	ErrInvalid = errors.New("international: invalid value")
	// ErrInvalidProvenance identifies incomplete or malformed source metadata.
	ErrInvalidProvenance = errors.New("international: invalid provenance")
	// ErrInvalidDataset identifies structurally invalid generated data.
	ErrInvalidDataset = errors.New("international: invalid dataset")
	// ErrResourceLimit identifies input rejected before excessive work.
	ErrResourceLimit = errors.New("international: resource limit exceeded")
)

// ParseError is a bounded diagnostic that never stores or echoes caller input.
type ParseError struct {
	kind   string
	reason string
}

// NewParseError creates a redacted parse error for a public value kind. It
// preserves package-defined reason classifications and replaces every other
// reason so caller-provided values cannot enter diagnostics.
func NewParseError(kind, reason string) *ParseError {
	kind = diagnosticKind(kind)
	return &ParseError{kind: kind, reason: diagnosticReason(reason)}
}

// Error returns a bounded, input-redacted diagnostic.
func (err *ParseError) Error() string {
	return fmt.Sprintf("international: invalid %s: %s", err.kind, err.reason)
}

// Unwrap makes all ParseError values match ErrInvalid.
func (err *ParseError) Unwrap() error {
	return ErrInvalid
}

func diagnosticKind(value string) string {
	switch value {
	case "calling code",
		"country alpha-3 code",
		"country code",
		"country numeric code",
		"currency code",
		"currency numeric code",
		"language code",
		"locale",
		"locale tag",
		"phone",
		"phone format",
		"phone number",
		"postal",
		"postal code",
		"postal normalization",
		"status",
		"subdivision code",
		"text",
		"value":
		return value
	default:
		return "value"
	}
}

func diagnosticReason(value string) string {
	switch value {
	case "absent tag",
		"absent value",
		"absent value has no text encoding",
		"ambiguous numeric identifier under selected status policy",
		"expected JSON string",
		"expected JSON string or null",
		"input is not canonical E.164",
		"invalid ISO 639 identifier",
		"invalid ISO 639 three-letter identifier",
		"invalid alpha-3 identifier",
		"invalid bounded value or country context",
		"invalid calling code metadata",
		"invalid canonical persistence text",
		"invalid canonicalization input",
		"invalid country context",
		"invalid extension",
		"invalid numeric identifier",
		"invalid syntax",
		"libphonenumber rejected input",
		"malformed BCP 47 tag",
		"malformed UTF-8",
		"malformed input",
		"malformed persistence text",
		"missing or repeated country context",
		"national input requires a region hint",
		"unaccepted ISO 3166-2 identifier",
		"unaccepted ISO 4217 identifier",
		"unaccepted alpha-2 identifier",
		"unknown alpha-3 identifier",
		"unknown enum value",
		"unknown numeric identifier",
		"unknown or malformed BCP 47 tag",
		"unknown or obsolete identifier",
		"unknown policy",
		"unknown style",
		"unknown wire spelling",
		"unsupported SQL source type",
		"unsupported value":
		return value
	default:
		return "invalid value"
	}
}
