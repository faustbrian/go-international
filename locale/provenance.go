package locale

import (
	international "github.com/faustbrian/go-international/v2"
	intlLanguage "github.com/faustbrian/go-international/v2/language"
)

// DatasetProvenance returns the IANA and x/text parsing provenance.
func DatasetProvenance() international.Provenance {
	return intlLanguage.DatasetProvenance()
}
