package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/its-lautaro/lautaro-env-store/internal/color"
)

func (a App) Clear(args []string) error {
	skipConfirm := false
	if len(args) == 1 {
		if args[0] != "--yes" && args[0] != "-y" {
			return fmt.Errorf("--clear solo admite --yes o -y; no entiendo: %s", args[0])
		}
		skipConfirm = true
	}
	entries, err := a.Store.Entries()
	if err != nil {
		return err
	}
	if len(entries) == 0 {
		fmt.Println("📭 No había nada que borrar")
		return nil
	}
	if !skipConfirm {
		fmt.Print(color.Yellow(fmt.Sprintf("⚠️  Vas a borrar %d variable(s) de %s. ¿Seguro? [y/N] ", len(entries), a.Store.Path)))
		answer, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		switch strings.ToLower(strings.TrimSpace(answer)) {
		case "y", "yes", "s", "si", "sí":
		default:
			fmt.Println("👌 Cancelado")
			return nil
		}
	}
	if err := a.Store.Clear(); err != nil {
		return err
	}
	names := make([]string, len(entries))
	for i, e := range entries {
		names[i] = e.Name
	}
	fmt.Println("🧹 Store vacío. " + unsetHint(names))
	return nil
}
