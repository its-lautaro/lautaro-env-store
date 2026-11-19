package shell

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// EnsureSourced agrega la línea source al primer rc que exista. Devuelve rc vacío si no existe ninguno.
func EnsureSourced(rcFiles []string, storePath, marker string) (rc string, added bool, err error) {
	line := sourceLine(storePath, marker)
	for _, rc := range rcFiles {
		data, err := os.ReadFile(rc)
		if errors.Is(err, os.ErrNotExist) {
			continue
		}
		if err != nil {
			return rc, false, fmt.Errorf("no pude leer %s: %w", rc, err)
		}
		if strings.Contains(string(data), line) {
			return rc, false, nil
		}
		f, err := os.OpenFile(rc, os.O_APPEND|os.O_WRONLY, 0o644)
		if err != nil {
			return rc, false, fmt.Errorf("no pude escribir %s: %w", rc, err)
		}
		defer f.Close()
		block := fmt.Sprintf("\n# %s: variables guardadas con `%s`\n%s\n", marker, marker, line)
		if _, err := f.WriteString(block); err != nil {
			return rc, false, fmt.Errorf("no pude escribir %s: %w", rc, err)
		}
		return rc, true, nil
	}
	return "", false, nil
}
