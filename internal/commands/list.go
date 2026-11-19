package commands

import (
	"fmt"
	"strings"

	"github.com/its-lautaro/lautaro-env-store/internal/color"
	"github.com/its-lautaro/lautaro-env-store/internal/env"
)

// List muestra los nombres guardados con su fecha y los copia al clipboard.
func List(args []string) {
	entries := env.ReadEntries()
	if len(entries) == 0 {
		fmt.Printf("📭 No hay variables guardadas en %s\n", color.Dim(env.Store))
		return
	}
	plural := ""
	if len(entries) > 1 {
		plural = "s"
	}
	fmt.Printf("🗂  %d variable%s en %s:\n", len(entries), plural, color.Dim(env.Store))
	names := make([]string, 0, len(entries))
	for _, e := range entries {
		names = append(names, e.Name)
		if e.Saved != "" {
			fmt.Printf("   🔑 %s  %s\n", color.Cyan(e.Name), color.Dim("(guardada "+e.Saved+")"))
		} else {
			fmt.Printf("   🔑 %s\n", color.Cyan(e.Name))
		}
	}
	env.CopyToClipboard(strings.Join(names, " "))
}
