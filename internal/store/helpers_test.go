package store

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// newTestStore devuelve un Store en un directorio temporal, opcionalmente con contenido inicial.
func newTestStore(t *testing.T, content ...string) Store {
	t.Helper()
	s := Store{Path: filepath.Join(t.TempDir(), "env")}
	if len(content) > 0 {
		if err := os.WriteFile(s.Path, []byte(strings.Join(content, "")), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	return s
}

func readFile(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}
