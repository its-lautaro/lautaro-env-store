package commands

import "fmt"

func (a App) Help(args []string) error {
	fmt.Fprintln(a.Out, a.helpText())
	return nil
}
