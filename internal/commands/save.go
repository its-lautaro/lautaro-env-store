package commands

import (
	"fmt"
	"strings"

	"github.com/its-lautaro/lautaro-env-store/internal/color"
	"github.com/its-lautaro/lautaro-env-store/internal/env"
)

// Save es el comando por defecto: guarda NOMBRE=valor y copia los nombres al clipboard.
func Save(args []string) {
	pairs := env.Parse(args)
	for _, p := range pairs {
		if !env.NameRE.MatchString(p.Name) {
			env.Fail("'%s' no es un nombre válido de variable de entorno", p.Name)
		}
	}

	replaced := env.Save(pairs)
	rc, added := env.EnsureSourced()
	names := make([]string, 0, len(pairs))
	for _, p := range pairs {
		names = append(names, p.Name)
	}
	copied := env.CopyToClipboard(strings.Join(names, " "))

	plural := ""
	if len(names) > 1 {
		plural = "s"
	}
	fmt.Printf("🔐 Variable%s guardada%s:\n", plural, plural)
	for _, name := range names {
		if replaced[name] {
			fmt.Printf("   ♻️  %s %s\n", color.Cyan(name), color.Yellow("(reemplazada)"))
		} else {
			fmt.Printf("   %s %s\n", color.Green("✅"), color.Cyan(name))
		}
	}
	if copied {
		fmt.Printf("📋 Nombre%s copiado%s al clipboard\n", plural, plural)
	} else {
		fmt.Println(color.Yellow("⚠️  No encontré un comando de clipboard (pbcopy/xclip/wl-copy); los nombres quedan arriba."))
	}
	if added {
		fmt.Printf("🧩 Agregué el source de %s a %s\n", color.Dim(env.Store), color.Dim(rc))
	}
}
