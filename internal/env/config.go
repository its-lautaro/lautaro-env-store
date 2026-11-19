package env

import (
	"path/filepath"
	"regexp"
)

var (
	home     = homeDir()
	Store    = storePath()
	RCFiles  = []string{filepath.Join(home, ".zshrc"), filepath.Join(home, ".bashrc")}
	NameRE   = regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*$`)
	EntryRE  = regexp.MustCompile(`^export\s+([A-Za-z_][A-Za-z0-9_]*)=`)
	SavedRE  = regexp.MustCompile(`# saved (\d{4}-\d{2}-\d{2} \d{2}:\d{2})$`)
	Marker   = "lautaro-env-store"
	UsageTxt = `🔐 lautaro-env-store — guarda variables de entorno en ~/.lautaro-env y las deja listas para el shell

Uso:
  lautaro-env-store NOMBRE=valor [NOMBRE=valor ...]
  lautaro-env-store <comando> [args]

Ejemplos:
  lautaro-env-store API_USER=usuario
  lautaro-env-store API_USER=usuario API_PASSWORD=secreto
  lautaro-env-store API_USER="Un Usuario"              # espacios o símbolos: comillas del shell

Comandos:
  --list, ls                 Lista los nombres guardados y su fecha (nunca los valores)
  --remove, rm NOMBRE...     Borra esas variables
  --clear [--yes]            Borra todas; pide confirmación salvo --yes
  --init                     Crea el store y agrega el source al rc del shell
  --help, -h                 Muestra esta ayuda

Notas:
  Los nombres de las variables guardadas se copian al clipboard; los valores nunca se imprimen.
  El store se crea con permisos 600 y ~/.zshrc (o ~/.bashrc) lo carga en cada terminal nueva.
  LAUTARO_ENV_FILE cambia la ruta del store.`
)

// SourceLine es la línea que se agrega al rc del shell para cargar el store.
func SourceLine() string {
	return `[ -f "` + Store + `" ] && source "` + Store + `"  # ` + Marker
}
