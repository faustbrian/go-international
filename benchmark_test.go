package international_test

import (
	"testing"

	"github.com/faustbrian/go-international/v5/country"
	"github.com/faustbrian/go-international/v5/currency"
	"github.com/faustbrian/go-international/v5/locale"
	"github.com/faustbrian/go-international/v5/phone"
)

// BenchmarkCountryLookup and the neighboring benchmarks use fixed public example
// inputs with already compiled metadata. Compare only equivalent provider/data
// versions, Go toolchains and hardware; no IO or dataset refresh is timed.
func BenchmarkCountryLookup(benchmark *testing.B) {
	for benchmark.Loop() {
		_, _ = country.Parse("FI")
	}
}

// BenchmarkLocaleCanonicalization includes both Parse and Canonical, not only
// canonicalization of a previously parsed tag.
func BenchmarkLocaleCanonicalization(benchmark *testing.B) {
	for benchmark.Loop() {
		tag, _ := locale.Parse("zh-hant-tw-u-ca-gregory")
		_, _ = tag.Canonical()
	}
}

func BenchmarkCurrencyLookup(benchmark *testing.B) {
	for benchmark.Loop() {
		_, _ = currency.Parse("EUR")
	}
}

func BenchmarkPhoneParse(benchmark *testing.B) {
	for benchmark.Loop() {
		_, _ = phone.ParseE164("+16502530000")
	}
}

// BenchmarkBulkCountryConversion excludes All allocation/setup and times one
// alpha-3 conversion for every current country per iteration. Corpus size changes
// alter work per iteration and must be accounted for in comparisons.
func BenchmarkBulkCountryConversion(benchmark *testing.B) {
	codes := country.All()
	benchmark.ResetTimer()
	for benchmark.Loop() {
		for _, code := range codes {
			_, _ = code.Alpha3()
		}
	}
}
