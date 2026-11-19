package commands

import (
	"fmt"

	"github.com/its-lautaro/lautaro-env-store/internal/color"
	"github.com/its-lautaro/lautaro-env-store/internal/env"
)

// Init crea el store vacío y agrega el source al rc del shell.
func Init(args []string) {
	env.EnsureStore()
	rc, added := env.EnsureSourced()
	if added {
		fmt.Printf("🧩 Store listo en %s; agregué el source a %s\n", color.Dim(env.Store), color.Dim(rc))
	} else {
		fmt.Printf("🧩 Store listo en %s %s\n", color.Dim(env.Store), color.Dim("(el shell ya lo cargaba)"))
	}
}
