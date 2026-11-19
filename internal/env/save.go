package env

import (
	"fmt"
	"time"
)

// Save guarda los pares (reemplazando los existentes) y devuelve los nombres que ya estaban.
func Save(pairs []Pair) map[string]bool {
	existing := map[string]bool{}
	for _, e := range ReadEntries() {
		existing[e.Name] = true
	}
	names := map[string]bool{}
	for _, p := range pairs {
		names[p.Name] = true
	}
	lines := DropNames(ReadLines(), names)
	stamp := time.Now().Format("2006-01-02 15:04")
	for _, p := range pairs {
		lines = append(lines, fmt.Sprintf("export %s=%s  # saved %s", p.Name, ShellQuote(p.Value), stamp))
	}
	WriteLines(lines)
	replaced := map[string]bool{}
	for name := range names {
		if existing[name] {
			replaced[name] = true
		}
	}
	return replaced
}
