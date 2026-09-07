package internationalwire_test

import (
	"testing"

	"github.com/faustbrian/go-international/internationalwire"
	"github.com/faustbrian/go-wire"
)

func FuzzLegacyWireDispatch(fuzzer *testing.F) {
	for _, seed := range []struct {
		format string
		input  []byte
	}{
		{string(wire.FormatJSON), []byte(`{"country":"FI","currency":"EUR"}`)},
		{string(wire.FormatXML), []byte(`<document><country>FI</country><currency>EUR</currency></document>`)},
		{string(wire.FormatYAML), []byte("country: FI\ncurrency: EUR\n")},
		{string(wire.FormatTOML), []byte("country = 'FI'\ncurrency = 'EUR'\n")},
		{string(wire.FormatMessagePack), []byte{0x80}},
		{string(wire.FormatSOAP), nil}, {string(wire.FormatCBOR), nil},
		{string(wire.FormatBSON), nil}, {"unknown", nil}, {"", nil},
	} {
		fuzzer.Add(seed.format, seed.input)
	}
	fuzzer.Fuzz(func(_ *testing.T, format string, input []byte) {
		if len(input) > 4096 {
			return
		}
		selected := wire.Format(format)
		_, _ = internationalwire.Encode(selected, document{})
		var target document
		_ = internationalwire.Decode(selected, input, &target)
	})
}
