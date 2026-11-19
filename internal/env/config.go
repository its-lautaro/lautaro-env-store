package env

import (
	"path/filepath"
	"regexp"
)

var (
	home    = homeDir()
	Store   = storePath()
	RCFiles = []string{filepath.Join(home, ".zshrc"), filepath.Join(home, ".bashrc")}
	NameRE  = regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*$`)
	EntryRE = regexp.MustCompile(`^export\s+([A-Za-z_][A-Za-z0-9_]*)=`)
	SavedRE = regexp.MustCompile(`# saved (\d{4}-\d{2}-\d{2} \d{2}:\d{2})$`)
	Marker  = "lautaro-env-store"
)

// SourceLine es la línea que se agrega al rc del shell para cargar el store.
func SourceLine() string {
	return `[ -f "` + Store + `" ] && source "` + Store + `"  # ` + Marker
}
