package locale

import (
	international "github.com/faustbrian/go-international/v5"
	intlLanguage "github.com/faustbrian/go-international/v5/language"
)

// DatasetProvenance returns the IANA and x/text parsing provenance.
func DatasetProvenance() international.Provenance {
	return intlLanguage.DatasetProvenance()
}
