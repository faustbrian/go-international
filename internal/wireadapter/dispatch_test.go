package wireadapter_test

import (
	"errors"
	"testing"

	"github.com/faustbrian/go-international/internal/wireadapter"
	"github.com/faustbrian/go-wire"
)

type payload struct {
	Value string `json:"value" xml:"value" yaml:"value" toml:"value" msgpack:"value"`
}

func TestDispatchUsesTheSelectedCodec(t *testing.T) {
	t.Parallel()

	tests := []struct {
		format wire.Format
		want   string
	}{
		{wire.FormatJSON, `{"value":"FI"}`},
		{wire.FormatXML, `<payload><value>FI</value></payload>`},
		{wire.FormatYAML, "value: FI\n"},
		{wire.FormatTOML, "value = \"FI\"\n"},
		{wire.FormatMessagePack, "\x81\xa5value\xa2FI"},
	}
	for _, test := range tests {
		t.Run(string(test.format), func(t *testing.T) {
			encoded, err := wireadapter.Encode(test.format, payload{Value: "FI"}, errors.New("unsupported"))
			if err != nil || string(encoded) != test.want {
				t.Fatalf("Encode() = %q, %v", encoded, err)
			}
			var decoded payload
			if err := wireadapter.Decode(test.format, []byte(test.want), &decoded, errors.New("unsupported")); err != nil || decoded.Value != "FI" {
				t.Fatalf("Decode() = %#v, %v", decoded, err)
			}
		})
	}
}

func TestDispatchReturnsTheCallingPackageSentinel(t *testing.T) {
	t.Parallel()

	unsupported := errors.New("selected unsupported format")
	for _, format := range []wire.Format{wire.FormatSOAP, wire.FormatCBOR, wire.FormatBSON, wire.Format("unknown")} {
		if _, err := wireadapter.Encode(format, payload{}, unsupported); !sameErrorIdentity(err, unsupported) {
			t.Fatalf("Encode(%s) = %v", format, err)
		}
		if err := wireadapter.Decode(format, nil, &payload{}, unsupported); !sameErrorIdentity(err, unsupported) {
			t.Fatalf("Decode(%s) = %v", format, err)
		}
	}
}

func TestUnsupportedFormatErrorHasStableValueIdentity(t *testing.T) {
	t.Parallel()

	var first error = wireadapter.UnsupportedFormatError{}
	var second error = wireadapter.UnsupportedFormatError{}
	if !sameErrorIdentity(first, second) || first.Error() != "unsupported international wire format" {
		t.Fatalf("identity/message = %v/%q", sameErrorIdentity(first, second), first.Error())
	}
}

func sameErrorIdentity(left, right error) bool {
	return left == right //nolint:errorlint // Direct identity is the helper contract under test.
}
