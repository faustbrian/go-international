// Package wireadapter owns the shared bounded dispatch used by both public
// International wire paths.
package wireadapter

import (
	"github.com/faustbrian/go-wire"
	"github.com/faustbrian/go-wire/jsonwire"
	"github.com/faustbrian/go-wire/msgpackwire"
	"github.com/faustbrian/go-wire/tomlwire"
	"github.com/faustbrian/go-wire/xmlwire"
	"github.com/faustbrian/go-wire/yamlwire"
)

// UnsupportedFormatError is the immutable initial identity shared by public
// paths. Its zero value is ready to use.
type UnsupportedFormatError struct{}

// Error returns the stable, value-safe unsupported-format message.
func (UnsupportedFormatError) Error() string {
	return "unsupported international wire format"
}

// Encode dispatches to the bounded default profile for supported formats.
func Encode(format wire.Format, value any, unsupported error) ([]byte, error) {
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
		return nil, unsupported
	default:
		return nil, unsupported
	}
}

// Decode dispatches to the bounded strict default profile for supported formats.
func Decode(format wire.Format, payload []byte, target any, unsupported error) error {
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
		return unsupported
	default:
		return unsupported
	}
}
