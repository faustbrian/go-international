// Package internationalwire provides bounded wire dispatch for structures
// containing international scalar types.
package internationalwire

import (
	"errors"

	"github.com/faustbrian/go-wire"
	"github.com/faustbrian/go-wire/jsonwire"
	"github.com/faustbrian/go-wire/msgpackwire"
	"github.com/faustbrian/go-wire/tomlwire"
	"github.com/faustbrian/go-wire/xmlwire"
	"github.com/faustbrian/go-wire/yamlwire"
)

// ErrUnsupportedFormat marks formats without a scalar-safe default adapter.
var ErrUnsupportedFormat = errors.New("unsupported international wire format")

// Encode uses wire's bounded default profile for a supported format.
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

// Decode uses wire's bounded strict default profile for a supported format.
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
