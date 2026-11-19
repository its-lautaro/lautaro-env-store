package commands

type Command struct {
	Names       []string
	Arg         string
	Description string
	Run         func(args []string) error
}
