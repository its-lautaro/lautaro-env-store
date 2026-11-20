package store

import (
	"os"
	"testing"
)

func TestStore_Create(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		content []string
		want    string
	}{
		{name: "creates missing file", content: nil, want: ""},
		{name: "keeps existing content", content: []string{"# header\nexport A='1'\n"}, want: "# header\nexport A='1'\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := newTestStore(t, tt.content...)
			if err := s.Create(); err != nil {
				t.Fatalf("Create() error = %v", err)
			}
			if got := readFile(t, s.Path); got != tt.want {
				t.Errorf("contenido = %q, want %q", got, tt.want)
			}
			info, err := os.Stat(s.Path)
			if err != nil {
				t.Fatal(err)
			}
			if got := info.Mode().Perm(); got != 0o600 {
				t.Errorf("permisos = %o, want 600", got)
			}
		})
	}
}
