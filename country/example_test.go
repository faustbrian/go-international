package country_test

import (
	"fmt"

	"github.com/faustbrian/go-international/v3/country"
)

func Example() {
	finland, _ := country.Parse("FI")
	alpha3, _ := finland.Alpha3()
	fmt.Println(finland, alpha3)
	// Output: FI FIN
}
