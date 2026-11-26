package store

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStore_writeLines(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		lines []string
		want  string
	}{
		{name: "nil writes empty file", lines: nil, want: ""},
		{name: "lines end with newline", lines: []string{"a", "b"}, want: "a\nb\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := newTestStore(t)
			if err := s.writeLines(tt.lines); err != nil {
				t.Fatalf("writeLines() error = %v", err)
			}
			if got := readFile(t, s.Path); got != tt.want {
				t.Errorf("contenido = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestStore_writeLines_FixesPermissionsOfExistingFile(t *testing.T) {
	t.Parallel()
	s := newTestStore(t, "old\n")
	if err := os.Chmod(s.Path, 0o644); err != nil {
		t.Fatal(err)
	}
	if err := s.writeLines([]string{"new"}); err != nil {
		t.Fatalf("writeLines() error = %v", err)
	}
	info, err := os.Stat(s.Path)
	if err != nil {
		t.Fatal(err)
	}
	if got := info.Mode().Perm(); got != 0o600 {
		t.Errorf("permisos = %o, want 600", got)
	}
}

func TestStore_writeLines_MissingDirectory(t *testing.T) {
	t.Parallel()
	s := Store{Path: filepath.Join(t.TempDir(), "no", "existe", "env")}
	if err := s.writeLines([]string{"a"}); err == nil {
		t.Error("writeLines() en un directorio inexistente: error = nil, want error")
	}
}
