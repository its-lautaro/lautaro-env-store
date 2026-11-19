package commands

import (
	"fmt"
	"strings"

	"github.com/its-lautaro/lautaro-env-store/internal/color"
	"github.com/its-lautaro/lautaro-env-store/internal/env"
)

// Remove borra las variables indicadas del store.
func Remove(names []string) {
	if len(names) == 0 {
		env.Fail("--remove necesita al menos un nombre")
	}
	missing := env.Remove(names)
	var removed []string
	for _, n := range names {
		if missing[n] {
			continue
		}
		removed = append(removed, n)
		fmt.Printf("🗑  %s %s\n", color.Cyan(n), color.Green("eliminada"))
	}
	for _, n := range names {
		if missing[n] {
			fmt.Printf("🤷 %s %s\n", color.Cyan(n), color.Yellow("no estaba guardada"))
		}
	}
	if len(removed) > 0 {
		fmt.Printf("🚀 En esta terminal además: %s\n", color.Bold("unset "+strings.Join(removed, " ")))
	}
}
