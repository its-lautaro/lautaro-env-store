package env

import "strings"

// ShellQuote envuelve el valor en comillas simples, escapando las que contenga.
func ShellQuote(value string) string {
	return "'" + strings.ReplaceAll(value, "'", `'"'"'`) + "'"
}
