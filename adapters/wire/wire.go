// Package internationalwire provides bounded strict wire dispatch for
// structures containing International scalar values.
package internationalwire

import (
	"github.com/faustbrian/go-wire/v3"
	"github.com/faustbrian/go-wire/v3/jsonwire"
	"github.com/faustbrian/go-wire/v3/msgpackwire"
	"github.com/faustbrian/go-wire/v3/tomlwire"
	"github.com/faustbrian/go-wire/v3/xmlwire"
	"github.com/faustbrian/go-wire/v3/yamlwire"
)

// ErrUnsupportedFormat marks formats without a scalar-safe adapter.
var ErrUnsupportedFormat error = unsupportedFormatError{}

// unsupportedFormatError is the comparable initial sentinel value; its zero
// value is ready to use.
type unsupportedFormatError struct{}

// Error returns the stable, value-safe unsupported-format message.
func (unsupportedFormatError) Error() string {
	return "unsupported international wire format"
}

// Encode encodes a value with the bounded default profile of a supported
// format. Unsupported formats return the package's current sentinel.
func Encode(format wire.Format, value any) ([]byte, error) {
	switch format {
	case wire.FormatJSON:
		return jsonwire.Encode(value, jsonwire.EncodeOptions{})
	case wire.FormatXML:
		return xmlwire.Encode(value, xmlwire.EncodeOptions{})
	case wire.FormatYAML:
		return yamlwire.Encode(value, yamlwire.EncodeOptions{})
	case wire.FormatTOML:
		return tomlwire.Encode(value, tomlwire.EncodeOptions{})
	case wire.FormatMessagePack:
		return msgpackwire.Encode(value, msgpackwire.EncodeOptions{})
	case wire.FormatSOAP, wire.FormatCBOR, wire.FormatBSON:
		return nil, ErrUnsupportedFormat
	default:
		return nil, ErrUnsupportedFormat
	}
}

// Decode synchronously borrows payload and a caller-owned target, using the
// selected Wire v3 codec's default profile. There is no context parameter or
// adapter-wide rollback guarantee; target mutation on error is codec-specific.
// Unsupported formats return the package's current sentinel without decoding.
func Decode(format wire.Format, payload []byte, target any) error {
	switch format {
	case wire.FormatJSON:
		return jsonwire.Decode(payload, target, jsonwire.DecodeOptions{})
	case wire.FormatXML:
		return xmlwire.Decode(payload, target, xmlwire.DecodeOptions{})
	case wire.FormatYAML:
		return yamlwire.Decode(payload, target, yamlwire.DecodeOptions{})
	case wire.FormatTOML:
		return tomlwire.Decode(payload, target, tomlwire.DecodeOptions{})
	case wire.FormatMessagePack:
		return msgpackwire.Decode(payload, target, msgpackwire.DecodeOptions{})
	case wire.FormatSOAP, wire.FormatCBOR, wire.FormatBSON:
		return ErrUnsupportedFormat
	default:
		return ErrUnsupportedFormat
	}
}
