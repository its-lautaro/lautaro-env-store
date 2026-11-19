package env

import (
	"errors"
	"os"
	"strings"
)

// ReadLines devuelve las líneas del store; vacío si no existe.
func ReadLines() []string {
	data, err := os.ReadFile(Store)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		Fail("no pude leer %s: %v", Store, err)
	}
	text := strings.TrimRight(string(data), "\n")
	if text == "" {
		return nil
	}
	return strings.Split(text, "\n")
}
