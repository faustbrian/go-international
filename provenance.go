package international

import (
	"fmt"
	"regexp"
	"slices"
	"time"
)

var sha256Pattern = regexp.MustCompile(`\A[0-9a-f]{64}\z`)

// Provenance is mutable descriptive metadata for reproducing and reviewing a
// dataset. Copy Transformations before independently mutating a copied value.
type Provenance struct {
	// Dataset names the projected dataset.
	Dataset string `json:"dataset"`
	// Source describes the retrieval location or upstream inputs.
	Source string `json:"source"`
	// RetrievedAt records the retrieval instant.
	RetrievedAt time.Time `json:"retrieved_at"`
	// UpstreamVersion identifies the pinned data or provider version.
	UpstreamVersion string `json:"upstream_version"`
	// License records the declared upstream license or terms.
	License string `json:"license"`
	// SHA256 records lowercase hexadecimal digest text; Validate does not authenticate it.
	SHA256 string `json:"sha256"`
	// Generator identifies the declared producer.
	Generator string `json:"generator"`
	// Transformations records ordered producer steps; struct copies share this slice.
	Transformations []string `json:"transformations"`
}

// Validate checks required nonempty fields, a nonzero retrieval instant, digest
// spelling and a nonempty transformation list without network access. It does
// not authenticate content, resolve sources, validate licenses or inspect each
// transformation string.
func (provenance Provenance) Validate() error {
	switch {
	case provenance.Dataset == "":
		return invalidProvenance("dataset is required")
	case provenance.Source == "":
		return invalidProvenance("source is required")
	case provenance.RetrievedAt.IsZero():
		return invalidProvenance("retrieval date is required")
	case provenance.UpstreamVersion == "":
		return invalidProvenance("upstream version is required")
	case provenance.License == "":
		return invalidProvenance("license is required")
	case !sha256Pattern.MatchString(provenance.SHA256):
		return invalidProvenance("SHA-256 must be lowercase hexadecimal")
	case provenance.Generator == "":
		return invalidProvenance("generator is required")
	case len(provenance.Transformations) == 0:
		return invalidProvenance("transformations are required")
	default:
		return nil
	}
}

// Equal compares field values and ordered transformations. Retrieval timestamps
// compare as instants, not by their location or textual representation.
func (provenance Provenance) Equal(other Provenance) bool {
	return provenance.Dataset == other.Dataset &&
		provenance.Source == other.Source &&
		provenance.RetrievedAt.Equal(other.RetrievedAt) &&
		provenance.UpstreamVersion == other.UpstreamVersion &&
		provenance.License == other.License &&
		provenance.SHA256 == other.SHA256 &&
		provenance.Generator == other.Generator &&
		slices.Equal(provenance.Transformations, other.Transformations)
}

func invalidProvenance(reason string) error {
	return fmt.Errorf("%w: %s", ErrInvalidProvenance, reason)
}
