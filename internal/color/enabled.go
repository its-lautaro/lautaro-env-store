// Package color pinta texto con códigos ANSI solo cuando la salida es una terminal y NO_COLOR no está definida.
package color

import "os"

var enabled = detect()

func detect() bool {
	if _, off := os.LookupEnv("NO_COLOR"); off {
		return false
	}
	return isTerminal(os.Stdout) && isTerminal(os.Stderr)
}
