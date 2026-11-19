package store

import "strings"

// Comillas simples POSIX: dentro de ellas nada se interpreta, y una comilla simple literal se escribe '"'"'.
func quote(value string) string {
	return "'" + strings.ReplaceAll(value, "'", `'"'"'`) + "'"
}
