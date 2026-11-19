package store

import (
	"os"
	"testing"
)

func TestStore_Clear(t *testing.T) {
	t.Parallel()
	s := newTestStore(t, "# header\nexport A='1'\n")
	if err := s.Clear(); err != nil {
		t.Fatalf("Clear() error = %v", err)
	}
	if got := readFile(t, s.Path); got != "" {
		t.Errorf("contenido después de Clear() = %q, want vacío", got)
	}
}

func TestStore_Clear_CreatesMissingFile(t *testing.T) {
	t.Parallel()
	s := newTestStore(t)
	if err := s.Clear(); err != nil {
		t.Fatalf("Clear() error = %v", err)
	}
	if _, err := os.Stat(s.Path); err != nil {
		t.Errorf("Clear() no creó el archivo: %v", err)
	}
}
