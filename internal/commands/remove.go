package commands

import (
	"errors"
	"fmt"

	"github.com/its-lautaro/lautaro-env-store/internal/color"
)

func (a App) Remove(names []string) error {
	if len(names) == 0 {
		return errors.New("--remove necesita al menos un nombre")
	}
	removed, missing, err := a.Store.Delete(names)
	if err != nil {
		return err
	}
	for _, n := range removed {
		fmt.Printf("🗑  %s %s\n", color.Cyan(n), color.Green("eliminada"))
	}
	for _, n := range missing {
		fmt.Printf("🤷 %s %s\n", color.Cyan(n), color.Yellow("no estaba guardada"))
	}
	if len(removed) > 0 {
		fmt.Println("🚀 " + unsetHint(removed))
	}
	return nil
}
