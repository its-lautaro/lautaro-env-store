package store

import (
	"fmt"
	"strings"
)

// quote devuelve value listo para el shell, siempre en una sola línea: el store se lee y se edita línea a línea.
//
// Sin caracteres de control usa comillas simples POSIX: dentro de ellas nada se interpreta, y una comilla
// simple literal se escribe '"'"'. Con caracteres de control (saltos de línea, tabs, ...) usa $'...' (bash y zsh)
// para escaparlos, porque en comillas simples un salto de línea partiría la variable en varias líneas.
func quote(value string) string {
	if !strings.ContainsFunc(value, isControl) {
		return "'" + strings.ReplaceAll(value, "'", `'"'"'`) + "'"
	}
	var b strings.Builder
	b.WriteString("$'")
	for i := 0; i < len(value); i++ {
		switch c := value[i]; c {
		case '\\':
			b.WriteString(`\\`)
		case '\'':
			b.WriteString(`\'`)
		case '\n':
			b.WriteString(`\n`)
		case '\t':
			b.WriteString(`\t`)
		case '\r':
			b.WriteString(`\r`)
		default:
			if isControl(rune(c)) {
				fmt.Fprintf(&b, `\x%02x`, c)
			} else {
				b.WriteByte(c)
			}
		}
	}
	b.WriteString("'")
	return b.String()
}

func isControl(r rune) bool {
	return r < 0x20 || r == 0x7f
}
