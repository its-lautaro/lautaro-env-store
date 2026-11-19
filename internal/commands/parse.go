package commands

import (
	"fmt"
	"strings"

	"github.com/its-lautaro/lautaro-env-store/internal/store"
)

func parse(args []string) ([]store.Var, error) {
	vars := make([]store.Var, 0, len(args))
	var bad []string
	seen := map[string]bool{}
	for _, a := range args {
		name, value, ok := strings.Cut(a, "=")
		if !ok {
			bad = append(bad, a)
			continue
		}
		if seen[name] {
			return nil, fmt.Errorf("'%s' está repetida; pasá cada variable una sola vez", name)
		}
		seen[name] = true
		vars = append(vars, store.Var{Name: name, Value: value})
	}
	if len(bad) > 0 {
		return nil, fmt.Errorf("cada argumento debe ser NOMBRE=valor; no entiendo: %s", strings.Join(bad, " "))
	}
	return vars, nil
}
