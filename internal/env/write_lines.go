package env

import (
	"os"
	"strings"
)

// WriteLines reescribe el store con permisos 600.
func WriteLines(lines []string) {
	content := ""
	if len(lines) > 0 {
		content = strings.Join(lines, "\n") + "\n"
	}
	if err := os.WriteFile(Store, []byte(content), 0o600); err != nil {
		Fail("no pude escribir %s: %v", Store, err)
	}
	if err := os.Chmod(Store, 0o600); err != nil {
		Fail("no pude ajustar permisos de %s: %v", Store, err)
	}
}
