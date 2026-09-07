// Package internationalwire provides bounded wire dispatch for structures
// containing international scalar types.
//
// Deprecated: use github.com/faustbrian/go-international/adapters/wire. This
// package remains supported for the longer of 180 days after successor public
// availability and two subsequently published stable root-module minor
// releases.
package internationalwire

import (
	"github.com/faustbrian/go-international/internal/wireadapter"
	"github.com/faustbrian/go-wire"
)

// ErrUnsupportedFormat marks formats without a scalar-safe default adapter.
var ErrUnsupportedFormat error = wireadapter.UnsupportedFormatError{}

// Encode uses wire's bounded default profile for a supported format.
func Encode(format wire.Format, value any) ([]byte, error) {
	return wireadapter.Encode(format, value, ErrUnsupportedFormat)
}

// Decode uses wire's bounded strict default profile for a supported format.
func Decode(format wire.Format, payload []byte, target any) error {
	return wireadapter.Decode(format, payload, target, ErrUnsupportedFormat)
}
