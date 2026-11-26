package commands

import (
	"slices"
	"strings"
	"testing"
)

func TestApp_Remove(t *testing.T) {
	t.Parallel()
	a := newTestApp(t)
	a.mustSave(t, "A=1", "B=2", "C=3")
	if err := a.Remove([]string{"A", "Z"}); err != nil {
		t.Fatalf("Remove() error = %v", err)
	}
	if got := a.entries(t); !slices.Equal(got, []string{"B", "C"}) {
		t.Errorf("entries = %q, want [B C]", got)
	}
	for _, w := range []string{"eliminada", "no estaba guardada", "unset A"} {
		if !strings.Contains(a.out.String(), w) {
			t.Errorf("falta %q en:\n%s", w, a.out.String())
		}
	}
}

func TestApp_Remove_NothingRemovedHasNoUnsetHint(t *testing.T) {
	t.Parallel()
	a := newTestApp(t)
	if err := a.Remove([]string{"Z"}); err != nil {
		t.Fatalf("Remove() error = %v", err)
	}
	if strings.Contains(a.out.String(), "unset") {
		t.Errorf("sugirió unset sin haber borrado nada:\n%s", a.out.String())
	}
}
