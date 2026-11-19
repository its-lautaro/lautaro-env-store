package commands

import (
	"fmt"

	"github.com/its-lautaro/lautaro-env-store/internal/color"
	"github.com/its-lautaro/lautaro-env-store/internal/shell"
)

func (a App) Init(args []string) error {
	if err := a.Store.Create(); err != nil {
		return err
	}
	rc, added, err := shell.EnsureSourced(a.RCFiles, a.Store.Path, Name)
	if err != nil {
		return err
	}
	if rc != "" && !added {
		fmt.Printf("🧩 Store listo en %s %s\n", color.Dim(a.Store.Path), color.Dim("(el shell ya lo cargaba)"))
		return nil
	}
	fmt.Printf("🧩 Store listo en %s\n", color.Dim(a.Store.Path))
	a.reportSourced(rc, added)
	return nil
}
