package currency

import (
	"time"

	international "github.com/faustbrian/go-international/v2"
)

// DatasetProvenance returns the pinned SIX List One and List Three provenance.
func DatasetProvenance() international.Provenance {
	return international.Provenance{
		Dataset:         "iso-4217-current-and-historic",
		Source:          "SIX ISO 4217 List One and List Three XML",
		RetrievedAt:     time.Date(2026, time.October, 2, 0, 0, 0, 0, time.UTC),
		UpstreamVersion: "ISO 4217 List One 2026-09-17; List Three 2026-01-01",
		License:         "SIX ISO 4217 Terms of Use",
		SHA256:          "091c69a6a28c78a3d20c34a7c77cc3236b1ed1c704dd7949765e1bfd40476ad9",
		Generator:       "international-generate/v1",
		Transformations: []string{
			"deduplicate country-specific currency rows",
			"preserve official names and withdrawal text",
			"sort by alphabetic code",
		},
	}
}
