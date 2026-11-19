package commands

import (
	"fmt"

	"github.com/its-lautaro/lautaro-env-store/internal/color"
	"github.com/its-lautaro/lautaro-env-store/internal/shell"
)

func (a App) Save(args []string) error {
	vars, err := parse(args)
	if err != nil {
		return err
	}
	replaced, err := a.Store.Put(vars)
	if err != nil {
		return err
	}
	rc, added, err := shell.EnsureSourced(a.RCFiles, a.Store.Path, Name)
	if err != nil {
		return err
	}
	names := make([]string, len(vars))
	for i, v := range vars {
		names[i] = v.Name
	}
	p := plural(len(names))
	fmt.Printf("🔐 Variable%s guardada%s:\n", p, p)
	for _, name := range names {
		if replaced[name] {
			fmt.Printf("   ♻️  %s %s\n", color.Cyan(name), color.Yellow("(reemplazada)"))
		} else {
			fmt.Printf("   %s %s\n", color.Green("✅"), color.Cyan(name))
		}
	}
	reportClipboard(names)
	a.reportSourced(rc, added)
	return nil
}
