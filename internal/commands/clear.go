package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/its-lautaro/lautaro-env-store/internal/env"
)

// Clear borra todas las variables, pidiendo confirmación salvo --yes/-y.
func Clear(args []string) {
	entries := env.ReadEntries()
	if len(entries) == 0 {
		fmt.Println("📭 No había nada que borrar")
		return
	}
	if !hasFlag(args, "--yes") && !hasFlag(args, "-y") {
		fmt.Printf("⚠️  Vas a borrar %d variable(s) de %s. ¿Seguro? [y/N] ", len(entries), env.Store)
		answer, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		switch strings.ToLower(strings.TrimSpace(answer)) {
		case "y", "yes", "s", "si", "sí":
		default:
			fmt.Println("👌 Cancelado")
			return
		}
	}
	env.WriteLines(nil)
	names := make([]string, 0, len(entries))
	for _, e := range entries {
		names = append(names, e.Name)
	}
	fmt.Printf("🧹 Store vacío. En esta terminal además: unset %s\n", strings.Join(names, " "))
}
