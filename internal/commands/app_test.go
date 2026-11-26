package commands

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/its-lautaro/lautaro-env-store/internal/store"
)

// fakeClipboard registra lo que se copia; si err no es nil, falla con ese error.
type fakeClipboard struct {
	mu     sync.Mutex
	copied []string
	err    error
}

func (c *fakeClipboard) Copy(text string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.err != nil {
		return c.err
	}
	c.copied = append(c.copied, text)
	return nil
}

// testApp agrupa la App con lo que los tests inspeccionan: su salida y el clipboard falso.
type testApp struct {
	App
	out  *bytes.Buffer
	clip *fakeClipboard
}

// newTestApp arma una App sobre un directorio temporal, con un rc existente, stdin vacío,
// salida en un buffer y un clipboard falso.
func newTestApp(t *testing.T) testApp {
	t.Helper()
	dir := t.TempDir()
	rc := filepath.Join(dir, ".zshrc")
	if err := os.WriteFile(rc, nil, 0o644); err != nil {
		t.Fatal(err)
	}
	out := &bytes.Buffer{}
	clip := &fakeClipboard{}
	return testApp{
		App: App{
			Store:     store.Store{Path: filepath.Join(dir, ".lautaro-env")},
			RCFiles:   []string{rc, filepath.Join(dir, ".bashrc")},
			Out:       out,
			In:        strings.NewReader(""),
			Clipboard: clip.Copy,
		},
		out:  out,
		clip: clip,
	}
}

func (a testApp) entries(t *testing.T) []string {
	t.Helper()
	entries, err := a.Store.Entries()
	if err != nil {
		t.Fatal(err)
	}
	names := make([]string, len(entries))
	for i, e := range entries {
		names[i] = e.Name
	}
	return names
}

// mustSave guarda args y limpia la salida y el clipboard para que el test vea solo lo suyo.
func (a testApp) mustSave(t *testing.T, args ...string) {
	t.Helper()
	if err := a.Save(args); err != nil {
		t.Fatalf("Save(%q) error = %v", args, err)
	}
	a.out.Reset()
	a.clip.copied = nil
}
