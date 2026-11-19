package commands

import (
	"fmt"
	"strings"

	"github.com/its-lautaro/lautaro-env-store/internal/color"
)

func (a App) reportClipboard(names []string) {
	p := plural(len(names))
	if err := a.Clipboard(strings.Join(names, " ")); err != nil {
		fmt.Fprintln(a.Out, color.Yellow(fmt.Sprintf("⚠️  %s; los nombres quedan arriba.", err)))
		return
	}
	fmt.Fprintf(a.Out, "📋 Nombre%s copiado%s al clipboard\n", p, p)
}
