package commands

import (
	"strings"
	"testing"

	"github.com/its-lautaro/lautaro-env-store/internal/store"
)

func storeAt(path string) store.Store { return store.Store{Path: path} }

func TestHelpRow(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		cmd  Command
		want []string
	}{
		{
			name: "without arg",
			cmd:  Command{Names: []string{"--list", "ls"}, Description: "Lista"},
			want: []string{"--list", "ls", "Lista"},
		},
		{
			name: "with arg",
			cmd:  Command{Names: []string{"--remove"}, Arg: "NOMBRE...", Description: "Borra"},
			want: []string{"--remove", "NOMBRE...", "Borra"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := helpRow(tt.cmd)
			for _, w := range tt.want {
				if !strings.Contains(got, w) {
					t.Errorf("helpRow() = %q, falta %q", got, w)
				}
			}
			if !strings.HasPrefix(got, "  ") {
				t.Errorf("helpRow() = %q, want sangría de dos espacios", got)
			}
		})
	}
}
