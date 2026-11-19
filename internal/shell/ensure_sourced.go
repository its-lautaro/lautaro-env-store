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
	for _, candidate := range rcFiles {
		data, err := os.ReadFile(candidate)
		if errors.Is(err, os.ErrNotExist) {
			continue
		}
		if err != nil {
			return candidate, false, fmt.Errorf("no pude leer %s: %w", candidate, err)
		}
		if strings.Contains(string(data), line) {
			return candidate, false, nil
		}
		f, err := os.OpenFile(candidate, os.O_APPEND|os.O_WRONLY, 0o644)
		if err != nil {
			return candidate, false, fmt.Errorf("no pude escribir %s: %w", candidate, err)
		}
		block := fmt.Sprintf("\n# %s: variables guardadas con `%s`\n%s\n", marker, marker, line)
		_, writeErr := f.WriteString(block)
		if err := errors.Join(writeErr, f.Close()); err != nil {
			return candidate, false, fmt.Errorf("no pude escribir %s: %w", candidate, err)
		}
		return candidate, true, nil
	}
	return "", false, nil
}
