package internationalwire_test

import (
	"errors"
	"fmt"
	"testing"

	successor "github.com/faustbrian/go-international/v4/adapters/wire"
	//lint:ignore SA1019 This import exercises the supported deprecated facade.
	//nolint:staticcheck // Compatibility coverage requires the supported deprecated facade.
	legacy "github.com/faustbrian/go-international/v4/internationalwire"
	"github.com/faustbrian/go-wire/v3"
	"github.com/faustbrian/go-wire/v3/msgpackwire"
)

func TestWireV3ErrorCategoriesReachBothPublicPaths(t *testing.T) {
	t.Parallel()

	for name, decode := range map[string]func(wire.Format, []byte, any) error{
		"canonical": successor.Decode,
		//nolint:staticcheck // Compatibility coverage requires the supported deprecated facade.
		"facade": legacy.Decode,
	} {
		t.Run(name, func(t *testing.T) {
			for _, test := range []struct {
				name   string
				input  []byte
				target any
				kind   wire.ErrorKind
				want   error
			}{
				{"malformed", []byte(`{"country":`), &document{}, wire.ErrorKindParse, wire.ErrParse},
				{"invalid-target", []byte(`{}`), document{}, wire.ErrorKindTarget, wire.ErrTarget},
			} {
				t.Run(test.name, func(t *testing.T) {
					err := decode(wire.FormatJSON, test.input, test.target)
					var categorized *wire.Error
					if !errors.Is(err, test.want) || !errors.As(err, &categorized) {
						t.Fatalf("Decode() = %v, want public Wire v3 category %v", err, test.want)
					}
					if categorized.Kind != test.kind || categorized.Format != wire.FormatJSON || categorized.Op != "decode" {
						t.Fatalf("Wire error category = %#v", categorized)
					}
				})
			}
		})
	}
}

func TestMessagePackWorkLimitPreservesTargetsThroughBothPublicPaths(t *testing.T) {
	t.Parallel()

	// Ordinary encoding produces a shallow map below byte, map, and aggregate
	// value limits, but its distinct keys exceed default comparison work.
	values := make(map[string]string, 3000)
	for index := range 3000 {
		values[fmt.Sprintf("k%04d", index)] = "FI"
	}
	payload, err := msgpackwire.Encode(values, msgpackwire.EncodeOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if len(values) >= msgpackwire.DefaultMaxMapPairs ||
		2*len(values)+1 >= msgpackwire.DefaultMaxTotalValues ||
		int64(len(payload)) >= msgpackwire.DefaultMaxBytes {
		t.Fatal("fixture unexpectedly exceeds another default admission limit")
	}

	for name, decode := range map[string]func(wire.Format, []byte, any) error{
		"canonical": successor.Decode,
		//nolint:staticcheck // Compatibility coverage requires the supported deprecated facade.
		"facade": legacy.Decode,
	} {
		t.Run(name, func(t *testing.T) {
			target := map[string]string{"preserved": "FI"}
			err := decode(wire.FormatMessagePack, payload, &target)
			var categorized *wire.Error
			if !errors.Is(err, wire.ErrSizeLimit) || !errors.As(err, &categorized) {
				t.Fatalf("Decode() = %v, want public Wire v3 size-limit category", err)
			}
			if categorized.Kind != wire.ErrorKindSizeLimit || categorized.Format != wire.FormatMessagePack || categorized.Op != "decode" {
				t.Fatalf("Wire error category = %#v", categorized)
			}
			if len(target) != 1 || target["preserved"] != "FI" {
				t.Fatalf("refused decode changed target: %#v", target)
			}
		})
	}
}
