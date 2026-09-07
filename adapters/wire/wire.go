// Package internationalwire provides bounded strict wire dispatch for
// structures containing International scalar values.
package internationalwire

import (
	"github.com/faustbrian/go-international/internal/wireadapter"
	"github.com/faustbrian/go-wire"
)

// ErrUnsupportedFormat marks formats without a scalar-safe adapter.
var ErrUnsupportedFormat error = wireadapter.UnsupportedFormatError{}

// Encode encodes a value in a supported format.
func Encode(format wire.Format, value any) ([]byte, error) {
	return wireadapter.Encode(format, value, ErrUnsupportedFormat)
}

// Decode decodes a value in a supported format.
func Decode(format wire.Format, payload []byte, target any) error {
	return wireadapter.Decode(format, payload, target, ErrUnsupportedFormat)
}
