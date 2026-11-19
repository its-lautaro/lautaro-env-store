package env

import "strings"

// Parse convierte argumentos NOMBRE=valor en pares; falla si alguno no tiene '='.
func Parse(args []string) []Pair {
	if len(args) == 0 {
		Usage(1)
	}
	var bad []string
	for _, a := range args {
		if !strings.Contains(a, "=") {
			bad = append(bad, a)
		}
	}
	if len(bad) > 0 {
		Fail("cada argumento debe ser NOMBRE=valor; no entiendo: %s", strings.Join(bad, " "))
	}
	pairs := make([]Pair, 0, len(args))
	for _, a := range args {
		name, value, _ := strings.Cut(a, "=")
		pairs = append(pairs, Pair{Name: name, Value: value})
	}
	return pairs
}
