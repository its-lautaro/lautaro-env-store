package shell

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const testStore = "/tmp/store"

func writeRC(t *testing.T, path, content string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestEnsureSourced(t *testing.T) {
	t.Parallel()
	line := sourceLine(testStore, "tool")
	tests := []struct {
		name      string
		existing  map[string]string // rc -> contenido inicial; los demás no existen
		wantRC    string
		wantAdded bool
	}{
		{name: "no rc exists", existing: nil, wantRC: "", wantAdded: false},
		{name: "appends to first existing rc", existing: map[string]string{"zshrc": "# mine\n", "bashrc": ""}, wantRC: "zshrc", wantAdded: true},
		{name: "skips missing rc", existing: map[string]string{"bashrc": "# mine\n"}, wantRC: "bashrc", wantAdded: true},
		{name: "already sourced", existing: map[string]string{"zshrc": "x\n" + line + "\n"}, wantRC: "zshrc", wantAdded: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			dir := t.TempDir()
			rcFiles := []string{filepath.Join(dir, "zshrc"), filepath.Join(dir, "bashrc")}
			for name, content := range tt.existing {
				writeRC(t, filepath.Join(dir, name), content)
			}
			rc, added, err := EnsureSourced(rcFiles, testStore, "tool")
			if err != nil {
				t.Fatalf("EnsureSourced() error = %v", err)
			}
			wantRC := ""
			if tt.wantRC != "" {
				wantRC = filepath.Join(dir, tt.wantRC)
			}
			if rc != wantRC || added != tt.wantAdded {
				t.Fatalf("EnsureSourced() = (%q, %v), want (%q, %v)", rc, added, wantRC, tt.wantAdded)
			}
			if rc == "" {
				return
			}
			data, err := os.ReadFile(rc)
			if err != nil {
				t.Fatal(err)
			}
			got := string(data)
			if !strings.HasPrefix(got, tt.existing[tt.wantRC]) {
				t.Errorf("se perdió el contenido previo del rc: %q", got)
			}
			if n := strings.Count(got, line); n != 1 {
				t.Errorf("la línea source aparece %d veces, want 1:\n%s", n, got)
			}
		})
	}
}

func TestEnsureSourced_Idempotent(t *testing.T) {
	t.Parallel()
	rc := filepath.Join(t.TempDir(), "zshrc")
	writeRC(t, rc, "")
	for i, wantAdded := range []bool{true, false, false} {
		_, added, err := EnsureSourced([]string{rc}, testStore, "tool")
		if err != nil {
			t.Fatalf("llamada %d: error = %v", i, err)
		}
		if added != wantAdded {
			t.Errorf("llamada %d: added = %v, want %v", i, added, wantAdded)
		}
	}
}

func TestEnsureSourced_UnreadableRC(t *testing.T) {
	t.Parallel()
	dir := t.TempDir() // un directorio existe pero no se puede leer como rc
	rc, _, err := EnsureSourced([]string{dir}, testStore, "tool")
	if err == nil {
		t.Fatal("EnsureSourced() con rc ilegible: error = nil, want error")
	}
	if rc != dir {
		t.Errorf("rc = %q, want %q", rc, dir)
	}
}
