package commands

import (
	"fmt"

	"github.com/its-lautaro/lautaro-env-store/internal/env"
)

// Init crea el store vacío y agrega el source al rc del shell.
func Init(args []string) {
	env.EnsureStore()
	rc, added := env.EnsureSourced()
	if added {
		fmt.Printf("🧩 Store listo en %s; agregué el source a %s\n", env.Store, rc)
	} else {
		fmt.Printf("🧩 Store listo en %s (el shell ya lo cargaba)\n", env.Store)
	}
}
