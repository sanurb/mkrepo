package main

import (
	"github.com/sanurb/mkrepo/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
