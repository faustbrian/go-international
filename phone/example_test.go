package phone_test

import (
	"fmt"

	"github.com/faustbrian/go-international/v3/country"
	"github.com/faustbrian/go-international/v3/phone"
)

func Example() {
	finland, _ := country.Parse("FI")
	number, _ := phone.Parse("040 123 4567", phone.ParseOptions{RegionHint: finland})
	fmt.Println(number.E164(), number.Possible(), number.Valid())
	// Output: +358401234567 true true
}
