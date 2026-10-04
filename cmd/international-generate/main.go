// Command international-generate updates pinned international metadata.
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/faustbrian/go-international/v3/internal/generate"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	if err := generate.RunContext(ctx, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
