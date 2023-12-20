package main

import (
	"os"

	"github.com/rarimo/rarime-auth-svc/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
