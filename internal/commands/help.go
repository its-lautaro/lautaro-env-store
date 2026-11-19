package commands

import "fmt"

func (a App) Help(args []string) error {
	fmt.Println(a.helpText())
	return nil
}
