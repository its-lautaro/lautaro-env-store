package commands

import (
	"fmt"
	"strings"

	"github.com/its-lautaro/lautaro-env-store/internal/store"
)

func parse(args []string) ([]store.Var, error) {
	vars := make([]store.Var, 0, len(args))
	var bad []string
	for _, a := range args {
		name, value, ok := strings.Cut(a, "=")
		if !ok {
			bad = append(bad, a)
			continue
		}
		vars = append(vars, store.Var{Name: name, Value: value})
	}
	if len(bad) > 0 {
		return nil, fmt.Errorf("cada argumento debe ser NOMBRE=valor; no entiendo: %s", strings.Join(bad, " "))
	}
	seen := map[string]bool{}
	for _, v := range vars {
		if !store.ValidName(v.Name) {
			return nil, fmt.Errorf("'%s' no es un nombre válido de variable de entorno", v.Name)
		}
		if seen[v.Name] {
			return nil, fmt.Errorf("'%s' está repetida; pasá cada variable una sola vez", v.Name)
		}
		seen[v.Name] = true
	}
	return vars, nil
}
