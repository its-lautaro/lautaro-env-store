package commands

import (
	"io"

	"github.com/its-lautaro/lautaro-env-store/internal/store"
)

// App recibe sus entradas y salidas en lugar de usar os.Stdin, os.Stdout y el clipboard del sistema
// directamente, para que main las conecte y los tests puedan reemplazarlas.
type App struct {
	Store     store.Store
	RCFiles   []string
	Out       io.Writer
	In        io.Reader
	Clipboard func(text string) error
}
