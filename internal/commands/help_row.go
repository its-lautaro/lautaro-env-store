package commands

import (
	"fmt"
	"strings"

	"github.com/its-lautaro/lautaro-env-store/internal/color"
)

func helpRow(c Command) string {
	flags := strings.Join(c.Names, ", ")
	plain := flags
	text := color.Bold(flags)
	if c.Arg != "" {
		plain += " " + c.Arg
		text += " " + color.Cyan(c.Arg)
	}
	return fmt.Sprintf("  %s%s%s", text, pad(plain), c.Description)
}
