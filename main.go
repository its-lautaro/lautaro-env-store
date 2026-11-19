// lautaro-env-store guarda variables de entorno en ~/.lautaro-env y las deja listas para el shell.
package main

import (
	"os"

	"github.com/its-lautaro/lautaro-env-store/internal/commands"
)

var subcommands = map[string]func([]string){
	"-h":       commands.Help,
	"--help":   commands.Help,
	"--list":   commands.List,
	"ls":       commands.List,
	"--remove": commands.Remove,
	"rm":       commands.Remove,
	"--clear":  commands.Clear,
	"--init":   commands.Init,
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		commands.Help(args)
	}
	if run, ok := subcommands[args[0]]; ok {
		run(args[1:])
		return
	}
	commands.Save(args)
}
