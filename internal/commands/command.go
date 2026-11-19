package commands

// Unlimited como MaxArgs admite cualquier cantidad de argumentos.
const Unlimited = -1

type Command struct {
	Names       []string
	Arg         string
	Description string
	MinArgs     int
	MaxArgs     int
	Run         func(args []string) error
}
