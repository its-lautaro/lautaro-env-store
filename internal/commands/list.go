package commands

import (
	"fmt"

	"github.com/its-lautaro/lautaro-env-store/internal/color"
)

func (a App) List(args []string) error {
	entries, err := a.Store.Entries()
	if err != nil {
		return err
	}
	if len(entries) == 0 {
		fmt.Fprintf(a.Out, "📭 No hay variables guardadas en %s\n", color.Dim(a.Store.Path))
		return nil
	}
	fmt.Fprintf(a.Out, "🗂  %d variable%s en %s:\n", len(entries), plural(len(entries)), color.Dim(a.Store.Path))
	names := make([]string, len(entries))
	for i, e := range entries {
		names[i] = e.Name
		if e.Saved != "" {
			fmt.Fprintf(a.Out, "   🔑 %s  %s\n", color.Cyan(e.Name), color.Dim("(guardada "+e.Saved+")"))
		} else {
			fmt.Fprintf(a.Out, "   🔑 %s\n", color.Cyan(e.Name))
		}
	}
	a.reportClipboard(names)
	return nil
}
