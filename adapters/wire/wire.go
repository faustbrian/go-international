// Package internationalwire provides bounded strict wire dispatch for
// structures containing International scalar values.
package internationalwire

import (
	"github.com/faustbrian/go-international/v4/internal/wireadapter"
	"github.com/faustbrian/go-wire/v3"
)

// ErrUnsupportedFormat marks formats without a scalar-safe adapter.
var ErrUnsupportedFormat error = wireadapter.UnsupportedFormatError{}

// Encode encodes a value in a supported format.
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
