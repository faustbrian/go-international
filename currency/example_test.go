package currency_test

import (
	"fmt"

	"github.com/faustbrian/go-international/v2/currency"
)

func Example() {
	euro, _ := currency.Parse("EUR")
	numeric, _ := euro.Numeric()
	minorUnits, specified := euro.MinorUnits()
	fmt.Println(numeric, minorUnits, specified)
	// Output: 978 2 true
}
