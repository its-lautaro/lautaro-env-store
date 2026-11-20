package store

import (
	"maps"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func TestStore_Put(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		initial      []string
		vars         []Var
		wantReplaced map[string]bool
		wantNames    []string
	}{
		{
			name:         "new store",
			vars:         []Var{{Name: "A", Value: "1"}, {Name: "B", Value: "2"}},
			wantReplaced: map[string]bool{},
			wantNames:    []string{"A", "B"},
		},
		{
			name:         "replaces existing and keeps the rest",
			initial:      []string{"# header\nexport A='old'  # saved x\nexport C='3'  # saved x\n"},
			vars:         []Var{{Name: "A", Value: "new"}, {Name: "B", Value: "2"}},
			wantReplaced: map[string]bool{"A": true},
			wantNames:    []string{"C", "A", "B"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := newTestStore(t, tt.initial...)
			replaced, err := s.Put(tt.vars)
			if err != nil {
				t.Fatalf("Put() error = %v", err)
			}
			if !maps.Equal(replaced, tt.wantReplaced) {
				t.Errorf("replaced = %v, want %v", replaced, tt.wantReplaced)
			}
			entries, err := s.Entries()
			if err != nil {
				t.Fatal(err)
			}
			var names []string
			for _, e := range entries {
				names = append(names, e.Name)
				if _, err := time.Parse(stampLayout, e.Saved); e.Name != "C" && err != nil {
					t.Errorf("%s: fecha %q no respeta %q: %v", e.Name, e.Saved, stampLayout, err)
				}
			}
			if strings.Join(names, ",") != strings.Join(tt.wantNames, ",") {
				t.Errorf("nombres = %v, want %v", names, tt.wantNames)
			}
			if len(tt.initial) > 0 && !strings.HasPrefix(readFile(t, s.Path), "# header\n") {
				t.Error("Put() no preservó las líneas que no son variables")
			}
		})
	}
}

func TestStore_Put_InvalidNameWritesNothing(t *testing.T) {
	t.Parallel()
	const initial = "export A='1'\n"
	s := newTestStore(t, initial)
	_, err := s.Put([]Var{{Name: "B", Value: "2"}, {Name: "1BAD", Value: "x"}})
	if err == nil {
		t.Fatal("Put() con nombre inválido: error = nil, want error")
	}
	if !strings.Contains(err.Error(), "1BAD") {
		t.Errorf("el error no menciona el nombre inválido: %v", err)
	}
	if got := readFile(t, s.Path); got != initial {
		t.Errorf("Put() modificó el store: %q, want %q", got, initial)
	}
}

// Lo que importa del store es que el shell lo pueda cargar con source y obtenga los valores exactos.
func TestStore_Put_SourcedByShell(t *testing.T) {
	t.Parallel()
	sh, err := exec.LookPath("sh")
	if err != nil {
		t.Skip("sh no disponible")
	}
	s := newTestStore(t)
	want := map[string]string{"USER_X": "un usuario", "PASS_X": `it's "$secret" $(id)`}
	if _, err := s.Put([]Var{{Name: "USER_X", Value: want["USER_X"]}, {Name: "PASS_X", Value: want["PASS_X"]}}); err != nil {
		t.Fatalf("Put() error = %v", err)
	}
	cmd := exec.Command(sh, "-c", `. "$1" && printf '%s\n%s' "$USER_X" "$PASS_X"`, "sh", s.Path)
	cmd.Env = []string{"PATH=" + os.Getenv("PATH")}
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("source del store falló: %v", err)
	}
	got := strings.SplitN(string(out), "\n", 2)
	if len(got) != 2 || got[0] != want["USER_X"] || got[1] != want["PASS_X"] {
		t.Errorf("valores tras source = %q, want [%q %q]", got, want["USER_X"], want["PASS_X"])
	}
}
