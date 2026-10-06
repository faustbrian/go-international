package locale_test

import (
	"fmt"

	"github.com/faustbrian/go-international/v4/locale"
)

func Example() {
	tag, _ := locale.Parse("EN-us")
	canonical, _ := tag.Canonical()
	fmt.Println(tag.String(), canonical.String())
	// Output: EN-us en-US
}
