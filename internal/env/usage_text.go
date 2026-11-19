package env

import (
	"strings"

	"github.com/its-lautaro/lautaro-env-store/internal/color"
)

// UsageText arma el texto de ayuda: nombre, uso, ejemplos, comandos y notas.
func UsageText() string {
	bin := color.Bold("lautaro-env-store")
	var b strings.Builder
	w := func(s string) { b.WriteString(s + "\n") }

	w("🔐 " + bin + " — guarda variables de entorno en ~/.lautaro-env y las deja listas para el shell")
	w("")
	w(color.Bold("Uso:"))
	w("  " + bin + " " + color.Cyan("NOMBRE") + "=" + color.Cyan("valor") + " [" + color.Cyan("NOMBRE") + "=" + color.Cyan("valor") + " ...]")
	w("  " + bin + " " + color.Cyan("<comando>") + " [" + color.Cyan("args") + "]")
	w("")
	w(color.Bold("Ejemplos:"))
	w("  " + bin + " " + color.Cyan("API_USER") + "=usuario")
	w("  " + bin + " " + color.Cyan("API_USER") + "=usuario " + color.Cyan("API_PASSWORD") + "=secreto")
	w("  " + bin + " " + color.Cyan("API_USER") + `="Un Usuario"              ` + color.Dim("# espacios o símbolos: comillas del shell"))
	w("")
	w(color.Bold("Comandos:"))
	w(command("--list, ls", "", "Lista los nombres guardados y su fecha (nunca los valores)"))
	w(command("--remove, rm", "NOMBRE...", "Borra esas variables"))
	w(command("--clear", "[--yes]", "Borra todas; pide confirmación salvo --yes"))
	w(command("--init", "", "Crea el store y agrega el source al rc del shell"))
	w(command("--help, -h", "", "Muestra esta ayuda"))
	w("")
	w(color.Bold("Notas:"))
	w(color.Dim("  Los nombres de las variables guardadas se copian al clipboard; los valores nunca se imprimen."))
	w(color.Dim("  El store se crea con permisos 600 y ~/.zshrc (o ~/.bashrc) lo carga en cada terminal nueva."))
	w(color.Dim("  LAUTARO_ENV_FILE cambia la ruta del store."))
	return strings.TrimRight(b.String(), "\n")
}
