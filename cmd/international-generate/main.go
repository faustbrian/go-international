// Command international-generate produces requested pinned metadata files using
// -country-output, -currency-output and -subdivision-output. Relative paths use
// the current directory. Interrupt cancellation reaches acquisition and checks
// before writes; outputs are not atomically published as one dataset bundle.
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/faustbrian/go-international/v4/internal/generate"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	if err := generate.RunContext(ctx, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
