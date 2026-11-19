package main

import (
	"path/filepath"
	"testing"

	"github.com/its-lautaro/lautaro-env-store/internal/commands"
)

func TestStorePath(t *testing.T) {
	tests := []struct {
		name string
		env  string
		want string
	}{
		{name: "default under home", env: "", want: filepath.Join("/home/u", ".lautaro-env")},
		{name: "env var overrides", env: "/custom/store", want: "/custom/store"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(commands.StoreEnvVar, tt.env)
			if got := storePath("/home/u"); got != tt.want {
				t.Errorf("storePath() = %q, want %q", got, tt.want)
			}
		})
	}
}
