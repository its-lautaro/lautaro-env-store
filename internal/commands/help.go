package commands

import "github.com/its-lautaro/lautaro-env-store/internal/env"

// Help imprime el uso y termina con código 0.
func Help(args []string) {
	env.Usage(0)
}
