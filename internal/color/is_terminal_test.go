package color

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIsTerminal(t *testing.T) {
	t.Parallel()
	f, err := os.Create(filepath.Join(t.TempDir(), "f"))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if isTerminal(f) {
		t.Error("isTerminal(archivo regular) = true, want false")
	}
	f.Close()
	if isTerminal(f) {
		t.Error("isTerminal(archivo cerrado) = true, want false")
	}
}
