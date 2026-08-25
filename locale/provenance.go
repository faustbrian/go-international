package locale

import (
	international "github.com/faustbrian/go-international"
	intlLanguage "github.com/faustbrian/go-international/language"
)

// DatasetProvenance returns the IANA and x/text parsing provenance.
func DatasetProvenance() international.Provenance {
	return intlLanguage.DatasetProvenance()
}
