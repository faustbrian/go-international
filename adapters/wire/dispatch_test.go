package internationalwire_test

import (
	"testing"

	internationalwire "github.com/faustbrian/go-international/v5/adapters/wire"
	"github.com/faustbrian/go-wire/v3"
)

type payload struct {
	Value string `json:"value" xml:"value" yaml:"value" toml:"value" msgpack:"value"`
}

func TestDispatchUsesTheSelectedCodec(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		format wire.Format
		want   string
	}{
		{wire.FormatJSON, `{"value":"FI"}`},
		{wire.FormatXML, `<payload><value>FI</value></payload>`},
		{wire.FormatYAML, "value: FI\n"},
		{wire.FormatTOML, "value = \"FI\"\n"},
		{wire.FormatMessagePack, "\x81\xa5value\xa2FI"},
	} {
		t.Run(string(test.format), func(t *testing.T) {
			encoded, err := internationalwire.Encode(test.format, payload{Value: "FI"})
			if err != nil || string(encoded) != test.want {
				t.Fatalf("Encode() = %q, %v", encoded, err)
			}
			var decoded payload
			if err := internationalwire.Decode(test.format, []byte(test.want), &decoded); err != nil || decoded.Value != "FI" {
				t.Fatalf("Decode() = %#v, %v", decoded, err)
			}
		})
	}
}

func TestUnsupportedFormatSentinelMessage(t *testing.T) {
	t.Parallel()

	if got := internationalwire.ErrUnsupportedFormat.Error(); got != "unsupported international wire format" {
		t.Fatalf("ErrUnsupportedFormat = %q", got)
	}
}
