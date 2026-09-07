//nolint:staticcheck // This parity fuzz test intentionally exercises the deprecated facade.
package internationalwire_test

//lint:file-ignore SA1019 This parity fuzz test intentionally exercises the deprecated facade.

import (
	"testing"

	successor "github.com/faustbrian/go-international/adapters/wire"
	legacy "github.com/faustbrian/go-international/internationalwire"
	"github.com/faustbrian/go-wire"
)

func FuzzSuccessorWireDispatch(fuzzer *testing.F) {
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
	fuzzer.Fuzz(func(t *testing.T, format string, input []byte) {
		if len(input) > 4096 {
			return
		}
		selected := wire.Format(format)
		legacyPayload, legacyEncodeErr := legacy.Encode(selected, document{})
		successorPayload, successorEncodeErr := successor.Encode(selected, document{})
		if (legacyEncodeErr == nil) != (successorEncodeErr == nil) || string(legacyPayload) != string(successorPayload) {
			t.Fatalf("encode parity differs")
		}
		var legacyTarget, successorTarget document
		legacyDecodeErr := legacy.Decode(selected, input, &legacyTarget)
		successorDecodeErr := successor.Decode(selected, input, &successorTarget)
		if (legacyDecodeErr == nil) != (successorDecodeErr == nil) || legacyTarget != successorTarget {
			t.Fatalf("decode parity differs")
		}
	})
}
