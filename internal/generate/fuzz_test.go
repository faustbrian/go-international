package generate

import (
	"bytes"
	"strings"
	"testing"
)

// FuzzGeneratedDataXML exercises the local limited-reader XML decoders with
// arbitrary bytes. It bypasses network authentication and output generation,
// testing parser panic resilience rather than authentic dataset conformance.
func FuzzGeneratedDataXML(fuzzer *testing.F) {
	fuzzer.Add([]byte(`<supplementalData><idValidity><id type="region" idStatus="regular">FI</id></idValidity></supplementalData>`))
	fuzzer.Add([]byte(`<ISO_4217 Pblshd="v"><CcyTbl></CcyTbl></ISO_4217>`))
	fuzzer.Add([]byte{0xff, 0xfe})
	fuzzer.Add([]byte(strings.Repeat("<id>", 4_096)))
	fuzzer.Fuzz(func(_ *testing.T, input []byte) {
		_, _ = parseRegionValidity(bytes.NewReader(input))
		_, _ = parseSubdivisionValidity(bytes.NewReader(input))
		_, _ = decodeCurrencyDocument(bytes.NewReader(input))
	})
}

// FuzzGeneratedDataRanges exercises both compact range helpers directly with
// arbitrary tokens, including malformed separators and non-ASCII endpoints.
// It is not an assertion of generated table completeness or range correctness.
func FuzzGeneratedDataRanges(fuzzer *testing.F) {
	for _, seed := range []string{
		"AA", "AA~Z", "fi18", "ad02~9", "", "\xff", "ＦＩ", "aa1~９",
		strings.Repeat("a", 4_096),
	} {
		fuzzer.Add(seed)
	}
	fuzzer.Fuzz(func(_ *testing.T, input string) {
		_, _ = expandCodeRange(input)
		_, _ = expandSubdivisionRange(input)
	})
}
