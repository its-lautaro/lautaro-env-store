package commands

import (
	"fmt"
	"strings"

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
		fmt.Printf("🗑  %s eliminada\n", n)
	}
	for _, n := range names {
		if missing[n] {
			fmt.Printf("🤷 %s no estaba guardada\n", n)
		}
	}
	if len(removed) > 0 {
		fmt.Printf("🚀 En esta terminal además: unset %s\n", strings.Join(removed, " "))
	}
}
