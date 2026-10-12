// Package internationalwire provides bounded wire dispatch for structures
// containing international scalar types.
//
// Deprecated: use github.com/faustbrian/go-international/v4/adapters/wire. This
// package remains supported for the longer of 180 days after successor public
// availability and two subsequently published stable root-module minor
// releases.
package internationalwire

import (
	"github.com/faustbrian/go-international/v4/internal/wireadapter"
	"github.com/faustbrian/go-wire/v3"
)

// ErrUnsupportedFormat marks formats without a scalar-safe default adapter.
var ErrUnsupportedFormat error = wireadapter.UnsupportedFormatError{}

// Encode uses wire's bounded default profile for a supported format.
func Encode(format wire.Format, value any) ([]byte, error) {
	return wireadapter.Encode(format, value, ErrUnsupportedFormat)
}

// Decode synchronously borrows payload and a caller-owned target, using the
// selected Wire v3 codec's default profile. There is no context parameter or
// adapter-wide rollback guarantee; target mutation on error is codec-specific.
// Unsupported formats return the package's current sentinel without decoding.
func Decode(format wire.Format, payload []byte, target any) error {
	return wireadapter.Decode(format, payload, target, ErrUnsupportedFormat)
}
