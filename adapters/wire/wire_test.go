//nolint:staticcheck // These compatibility tests intentionally exercise the deprecated facade.
package internationalwire_test

//lint:file-ignore SA1019 These compatibility tests intentionally exercise the deprecated facade.

import (
	"errors"
	"testing"

	successor "github.com/faustbrian/go-international/adapters/wire"
	"github.com/faustbrian/go-international/country"
	"github.com/faustbrian/go-international/currency"
	legacy "github.com/faustbrian/go-international/internationalwire"
	"github.com/faustbrian/go-wire"
)

type document struct {
	Country  country.Code  `json:"country" xml:"country" yaml:"country" toml:"country" msgpack:"country"`
	Currency currency.Code `json:"currency" xml:"currency" yaml:"currency" toml:"currency" msgpack:"currency"`
}

func TestSupportedFormatsPreserveLegacyBytesAndValues(t *testing.T) {
	t.Parallel()

	countryCode, _ := country.Parse("FI")
	currencyCode, _ := currency.Parse("EUR")
	want := document{Country: countryCode, Currency: currencyCode}
	for _, format := range []wire.Format{wire.FormatJSON, wire.FormatXML, wire.FormatYAML, wire.FormatTOML, wire.FormatMessagePack} {
		t.Run(string(format), func(t *testing.T) {
			legacyPayload, legacyErr := legacy.Encode(format, want)
			successorPayload, successorErr := successor.Encode(format, want)
			if legacyErr != nil || successorErr != nil || string(successorPayload) != string(legacyPayload) {
				t.Fatalf("legacy=%q/%v successor=%q/%v", legacyPayload, legacyErr, successorPayload, successorErr)
			}
			var got document
			if err := successor.Decode(format, successorPayload, &got); err != nil || got != want {
				t.Fatalf("Decode() = %#v, %v", got, err)
			}
		})
	}
}

func TestUnsupportedFormatsUseIndependentCurrentSentinels(t *testing.T) {
	if !sameErrorIdentity(successor.ErrUnsupportedFormat, legacy.ErrUnsupportedFormat) {
		t.Fatal("default sentinel identity differs")
	}
	legacyOriginal := legacy.ErrUnsupportedFormat
	successorOriginal := successor.ErrUnsupportedFormat
	t.Cleanup(func() {
		legacy.ErrUnsupportedFormat = legacyOriginal
		successor.ErrUnsupportedFormat = successorOriginal
	})

	legacyReplacement := errors.New("legacy replacement")
	legacy.ErrUnsupportedFormat = legacyReplacement
	for _, format := range []wire.Format{wire.FormatSOAP, wire.FormatCBOR, wire.FormatBSON, wire.Format("unknown")} {
		if _, err := legacy.Encode(format, document{}); !sameErrorIdentity(err, legacyReplacement) {
			t.Fatalf("legacy Encode(%s) = %v", format, err)
		}
		if _, err := successor.Encode(format, document{}); !sameErrorIdentity(err, successorOriginal) {
			t.Fatalf("successor Encode(%s) = %v", format, err)
		}
	}

	successorReplacement := errors.New("successor replacement")
	successor.ErrUnsupportedFormat = successorReplacement
	for _, format := range []wire.Format{wire.FormatSOAP, wire.FormatCBOR, wire.FormatBSON, wire.Format("unknown")} {
		if err := successor.Decode(format, nil, &document{}); !sameErrorIdentity(err, successorReplacement) {
			t.Fatalf("successor Decode(%s) = %v", format, err)
		}
		if err := legacy.Decode(format, nil, &document{}); !sameErrorIdentity(err, legacyReplacement) {
			t.Fatalf("legacy Decode(%s) = %v", format, err)
		}
	}
}

func sameErrorIdentity(left, right error) bool {
	return left == right //nolint:errorlint // Direct identity is the compatibility contract under test.
}
