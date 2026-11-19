package env

import (
	"os"
	"strings"
)

// EnsureSourced agrega la línea de source al primer rc que exista (zsh primero), una sola vez.
// Devuelve el rc usado y si se agregó la línea.
func EnsureSourced() (string, bool) {
	for _, rc := range RCFiles {
		data, err := os.ReadFile(rc)
		if err != nil {
			continue
		}
		if strings.Contains(string(data), Marker) {
			return rc, false
		}
		f, err := os.OpenFile(rc, os.O_APPEND|os.O_WRONLY, 0o644)
		if err != nil {
			Fail("no pude escribir %s: %v", rc, err)
		}
		defer f.Close()
		f.WriteString("\n# lautaro-env-store: variables guardadas con `lautaro-env-store`\n" + SourceLine() + "\n")
		return rc, true
	}
	return "", false
}
