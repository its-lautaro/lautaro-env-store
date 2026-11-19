package commands

import (
	"fmt"
	"strings"

	"github.com/its-lautaro/lautaro-env-store/internal/color"
)

func unsetHint(names []string) string {
	return fmt.Sprintf("En esta terminal además: %s", color.Bold("unset "+strings.Join(names, " ")))
}
