package main

import (
	"os"
	"path/filepath"
)

func storePath(home string) string {
	if p := os.Getenv("LAUTARO_ENV_FILE"); p != "" {
		return p
	}
	return filepath.Join(home, ".lautaro-env")
}
