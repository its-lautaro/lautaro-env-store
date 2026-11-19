package main

import (
	"os"
	"path/filepath"

	"github.com/its-lautaro/lautaro-env-store/internal/commands"
)

func storePath(home string) string {
	if p := os.Getenv(commands.StoreEnvVar); p != "" {
		return p
	}
	return filepath.Join(home, ".lautaro-env")
}
