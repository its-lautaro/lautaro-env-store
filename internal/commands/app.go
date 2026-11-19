package commands

import "github.com/its-lautaro/lautaro-env-store/internal/store"

type App struct {
	Store   store.Store
	RCFiles []string
}
