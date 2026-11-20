package commands

import (
	"fmt"

	"github.com/its-lautaro/lautaro-env-store/internal/color"
)

func (a App) Remove(names []string) error {
	removed, missing, err := a.Store.Delete(names)
	if err != nil {
		return err
	}
	for _, n := range removed {
		fmt.Fprintf(a.Out, "🗑  %s %s\n", color.Cyan(n), color.Green("eliminada"))
	}
	for _, n := range missing {
		fmt.Fprintf(a.Out, "🤷 %s %s\n", color.Cyan(n), color.Yellow("no estaba guardada"))
	}
	if len(removed) > 0 {
		fmt.Fprintln(a.Out, "🚀 "+unsetHint(removed))
	}
	return nil
}
