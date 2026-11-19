package commands

import (
	"fmt"
	"slices"
)

func (a App) Run(args []string) error {
	if len(args) == 0 {
		return a.Help(nil)
	}
	for _, c := range a.Commands() {
		if !slices.Contains(c.Names, args[0]) {
			continue
		}
		rest := args[1:]
		if len(rest) < c.MinArgs {
			return fmt.Errorf("%s necesita %s", args[0], c.Arg)
		}
		if c.MaxArgs == 0 && len(rest) > 0 {
			return fmt.Errorf("%s no lleva argumentos", args[0])
		}
		if c.MaxArgs > 0 && len(rest) > c.MaxArgs {
			return fmt.Errorf("%s admite como máximo %d argumento(s): %s", args[0], c.MaxArgs, c.Arg)
		}
		return c.Run(rest)
	}
	return a.Save(args)
}
