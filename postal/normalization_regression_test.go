package postal_test

import (
	"errors"
	"strings"
	"testing"

	international "github.com/faustbrian/go-international/v4"
	"github.com/faustbrian/go-international/v4/country"
	"github.com/faustbrian/go-international/v4/postal"
)

func TestNFCNormalizationPreservesCanonicalComposition(t *testing.T) {
	t.Parallel()
	context, err := country.Parse("FI")
	if err != nil {
		t.Fatal(err)
	}
	// Frozen expectations cover the NFC corrections in golang/text v0.42.0,
	// composition_test.go and the Unicode 17 canonical composition tables.
	tests := []struct{ name, input, want string }{
		{"intervening starter", "i\U000113c2\u0300\u0316", "i\U000113c2\u0316\u0300"},
		{"ordinary composition", "i\u0300\u0316", "\u00ec\u0316"},
		{"supplementary plane", "\U00010041\u0300", "\U00010041\u0300"},
		{"composition after Hangul", "\u1100\u1161\U000113c2\U000113c2", "\uac00\U000113c5"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			original, err := postal.Parse(test.input, context)
			if err != nil {
				t.Fatal(err)
			}
			preserved, err := original.Normalize(postal.NormalizeOptions{})
			if err != nil || preserved.Raw() != test.input || preserved.Country() != context {
				t.Fatalf("preserve = %q, %v", preserved.Raw(), err)
			}
			normalized, err := original.Normalize(postal.NormalizeOptions{Unicode: postal.UnicodeNFC})
			if err != nil || normalized.Raw() != test.want || normalized.Country() != context {
				t.Fatalf("NFC = %q, %v; want %q", normalized.Raw(), err, test.want)
			}
			if original.Raw() != test.input || original.Country() != context {
				t.Fatal("normalization changed the original")
			}
			for _, value := range []postal.Code{original, normalized} {
				encoded := "FI\t" + value.Raw()
				text, err := value.MarshalText()
				if err != nil || string(text) != encoded {
					t.Fatalf("text = %q, %v", text, err)
				}
				stored, err := value.Value()
				if err != nil || stored != encoded {
					t.Fatalf("SQL value = %q, %v", stored, err)
				}
				var decoded postal.Code
				if err := decoded.Scan(stored); err != nil || decoded.Raw() != value.Raw() || decoded.Country() != context {
					t.Fatalf("SQL decode = %q, %v", decoded.Raw(), err)
				}
			}
		})
	}
}

func TestNFCNormalizationRechecksExpandedByteBudget(t *testing.T) {
	t.Parallel()
	context, err := country.Parse("FI")
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct{ name, input, want string }{
		{"below limit", strings.Repeat("\u0344", 7) + "abc", strings.Repeat("\u0308\u0301", 7) + "abc"},
		{"at limit", strings.Repeat("\u0344", 8), strings.Repeat("\u0308\u0301", 8)},
		{"above limit", strings.Repeat("\u0344", 8) + "a", ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			original, err := postal.Parse(test.input, context)
			if err != nil {
				t.Fatal(err)
			}
			normalized, err := original.Normalize(postal.NormalizeOptions{Unicode: postal.UnicodeNFC})
			if test.want == "" {
				if !errors.Is(err, international.ErrResourceLimit) || !normalized.IsZero() {
					t.Fatalf("overflow result = %q, %v", normalized.Raw(), err)
				}
			} else if err != nil || normalized.Raw() != test.want || normalized.Country() != context {
				t.Fatalf("NFC = %q, %v; want %q", normalized.Raw(), err, test.want)
			}
			if original.Raw() != test.input || original.Country() != context {
				t.Fatal("normalization changed original")
			}
		})
	}
}
