package main

import (
	"os"
	"path/filepath"

	"github.com/its-lautaro/lautaro-env-store/internal/commands"
	"github.com/its-lautaro/lautaro-env-store/internal/shell"
	"github.com/its-lautaro/lautaro-env-store/internal/store"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		fail(err)
	}
	app := commands.App{
		Store:     store.Store{Path: storePath(home)},
		RCFiles:   []string{filepath.Join(home, ".zshrc"), filepath.Join(home, ".bashrc")},
		Out:       os.Stdout,
		In:        os.Stdin,
		Clipboard: shell.CopyToClipboard,
	}
	if err := app.Run(os.Args[1:]); err != nil {
		fail(err)
	}
}
