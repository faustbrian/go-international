// Command international-generate updates pinned international metadata.
package main

import (
	"fmt"
	"os"

	"github.com/faustbrian/go-international/v2/internal/generate"
)

func main() {
	if err := generate.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
