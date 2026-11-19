package store

import (
	"fmt"
	"os"
	"strings"
)

func (s Store) writeLines(lines []string) error {
	content := ""
	if len(lines) > 0 {
		content = strings.Join(lines, "\n") + "\n"
	}
	if err := os.WriteFile(s.Path, []byte(content), 0o600); err != nil {
		return fmt.Errorf("no pude escribir %s: %w", s.Path, err)
	}
	// os.WriteFile aplica el modo solo al crear el archivo; el Chmod cubre un store que ya existía con otros permisos.
	if err := os.Chmod(s.Path, 0o600); err != nil {
		return fmt.Errorf("no pude ajustar permisos de %s: %w", s.Path, err)
	}
	return nil
}
