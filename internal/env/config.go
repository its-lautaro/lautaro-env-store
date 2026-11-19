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
	UsageTxt = `lautaro-env-store: guarda variables de entorno en ~/.lautaro-env y las deja listas para el shell.

Formatos aceptados:
  lautaro-env-store API_USER=usuario                          # una
  lautaro-env-store API_USER=usuario API_PASSWORD=secreto     # varias
  lautaro-env-store API_USER="Un Usuario"                     # valores con espacios: comillas del shell
  lautaro-env-store --init                                    # solo crea el store y el source en el rc

Mantenimiento:
  lautaro-env-store --list                     # nombres guardados y fecha (nunca los valores)
  lautaro-env-store --remove API_USER          # borra esas variables
  lautaro-env-store --clear [--yes]            # borra todas`
)

// SourceLine es la línea que se agrega al rc del shell para cargar el store.
func SourceLine() string {
	return `[ -f "` + Store + `" ] && source "` + Store + `"  # ` + Marker
}
