package env

import (
	"fmt"
	"strings"

	"github.com/its-lautaro/lautaro-env-store/internal/color"
)

// command arma una fila de la sección Comandos del help, alineando la descripción en una columna fija.
func command(flags, arg, description string) string {
	plain := flags
	if arg != "" {
		plain += " " + arg
	}
	padding := strings.Repeat(" ", max(27-len([]rune(plain)), 2))
	text := color.Bold(flags)
	if arg != "" {
		text += " " + color.Cyan(arg)
	}
	return fmt.Sprintf("  %s%s%s", text, padding, description)
}
