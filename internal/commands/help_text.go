package commands

import (
	"strings"

	"github.com/its-lautaro/lautaro-env-store/internal/color"
)

func (a App) helpText() string {
	bin := color.Bold(Name)
	var b strings.Builder
	w := func(s string) { b.WriteString(s + "\n") }

	w("🔐 " + bin + " — guarda variables de entorno en " + a.Store.Path + " y las deja listas para el shell")
	w("")
	w(color.Bold("Uso:"))
	w("  " + bin + " " + color.Cyan("NOMBRE") + "=" + color.Cyan("valor") + " [" + color.Cyan("NOMBRE") + "=" + color.Cyan("valor") + " ...]")
	w("  " + bin + " " + color.Cyan("<comando>") + " [" + color.Cyan("args") + "]")
	w("")
	w(color.Bold("Ejemplos:"))
	w("  " + bin + " " + color.Cyan("API_USER") + "=usuario")
	w("  " + bin + " " + color.Cyan("API_USER") + "=usuario " + color.Cyan("API_PASSWORD") + "=secreto")
	example := `="Un Usuario"`
	w("  " + bin + " " + color.Cyan("API_USER") + example + pad(Name+" API_USER"+example) + color.Dim("# espacios o símbolos: comillas del shell"))
	w("")
	w(color.Bold("Comandos:"))
	for _, c := range a.Commands() {
		w(helpRow(c))
	}
	w("")
	w(color.Bold("Notas:"))
	w(color.Dim("  Los nombres de las variables guardadas se copian al clipboard; los valores nunca se imprimen."))
	w(color.Dim("  El store se crea con permisos 600 y " + strings.Join(a.RCFiles, " o ") + " lo carga en cada terminal nueva."))
	w(color.Dim("  " + StoreEnvVar + " cambia la ruta del store."))
	return strings.TrimRight(b.String(), "\n")
}
