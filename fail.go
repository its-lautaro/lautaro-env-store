package main

import (
	"fmt"
	"os"

	"github.com/its-lautaro/lautaro-env-store/internal/color"
	"github.com/its-lautaro/lautaro-env-store/internal/commands"
)

func fail(err error) {
	fmt.Fprintln(os.Stderr, color.Red("❌ "+err.Error()))
	fmt.Fprintln(os.Stderr, color.Dim("💡 Para ver la ayuda: ")+color.Bold(commands.Name+" --help"))
	os.Exit(1)
}
