package commands

import (
	"fmt"
	"strings"

	"github.com/its-lautaro/lautaro-env-store/internal/color"
	"github.com/its-lautaro/lautaro-env-store/internal/shell"
)

func reportClipboard(names []string) {
	p := plural(len(names))
	if err := shell.CopyToClipboard(strings.Join(names, " ")); err != nil {
		fmt.Println(color.Yellow(fmt.Sprintf("⚠️  %s; los nombres quedan arriba.", err)))
		return
	}
	fmt.Printf("📋 Nombre%s copiado%s al clipboard\n", p, p)
}
