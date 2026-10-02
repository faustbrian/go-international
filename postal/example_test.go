package postal_test

import (
	"fmt"

	"github.com/faustbrian/go-international/v2/country"
	"github.com/faustbrian/go-international/v2/postal"
)

func Example() {
	finland, _ := country.Parse("FI")
	code, _ := postal.Parse(" 00100 ", finland)
	normalized, _ := code.Normalize(postal.NormalizeOptions{Spaces: postal.SpacesCollapseASCII})
	fmt.Println(normalized.Country(), normalized.Raw())
	// Output: FI 00100
}
