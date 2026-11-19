package commands

import (
	"slices"
	"strings"
	"testing"
)

func TestApp_List(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		save       []string
		want       []string
		notWant    []string
		wantCopied []string
	}{
		{name: "empty store", want: []string{"No hay variables guardadas"}},
		{
			name:       "lists names never values",
			save:       []string{"A=valor-secreto", "B=otro"},
			want:       []string{"2 variables", "A", "B", "guardada", "Nombres copiados"},
			notWant:    []string{"valor-secreto", "otro"},
			wantCopied: []string{"A B"},
		},
		{name: "singular", save: []string{"A=1"}, want: []string{"1 variable en", "Nombre copiado"}, wantCopied: []string{"A"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := newTestApp(t)
			if len(tt.save) > 0 {
				a.mustSave(t, tt.save...)
			}
			if err := a.List(nil); err != nil {
				t.Fatalf("List() error = %v", err)
			}
			out := a.out.String()
			for _, w := range tt.want {
				if !strings.Contains(out, w) {
					t.Errorf("falta %q en:\n%s", w, out)
				}
			}
			for _, w := range tt.notWant {
				if strings.Contains(out, w) {
					t.Errorf("List() imprimió %q:\n%s", w, out)
				}
			}
			if !slices.Equal(a.clip.copied, tt.wantCopied) {
				t.Errorf("clipboard = %q, want %q", a.clip.copied, tt.wantCopied)
			}
		})
	}
}
