package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		if isUsageError(err) {
			fmt.Fprintln(os.Stderr)
			PrintUsage(os.Stderr)
		}
		os.Exit(1)
	}
}
