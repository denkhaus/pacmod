package main

import (
	"os"

	"github.com/denkhaus/pacmod/internal/commands"
)

func main() {
	if err := commands.NewDefaultCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
