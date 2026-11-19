package env

import (
	"os"
	"path/filepath"
)

// storePath es LAUTARO_ENV_FILE si está definida, si no ~/.lautaro-env.
func storePath() string {
	if p := os.Getenv("LAUTARO_ENV_FILE"); p != "" {
		return p
	}
	return filepath.Join(home, ".lautaro-env")
}
