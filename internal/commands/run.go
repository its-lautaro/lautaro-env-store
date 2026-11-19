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
		if c.Arg == "" && len(args) > 1 {
			return fmt.Errorf("%s no lleva argumentos", args[0])
		}
		return c.Run(args[1:])
	}
	return a.Save(args)
}
