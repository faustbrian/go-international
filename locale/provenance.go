package locale

import (
	international "github.com/faustbrian/go-international/v3"
	intlLanguage "github.com/faustbrian/go-international/v3/language"
)

// DatasetProvenance returns the IANA and x/text parsing provenance.
func DatasetProvenance() international.Provenance {
	return intlLanguage.DatasetProvenance()
}
