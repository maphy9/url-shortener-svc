package main

import (
	"os"

	"github.com/maphy9/url-shortener-svc/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
