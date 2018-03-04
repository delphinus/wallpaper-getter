package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	ctx := context.Background()
	if err := process(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "error in process: %v", err)
		os.Exit(1)
	}
}
