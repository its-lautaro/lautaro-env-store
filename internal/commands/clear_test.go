package commands

import (
	"strings"
	"testing"
)

func TestApp_Clear(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		args        []string
		stdin       string
		wantCleared bool
		wantOut     string
	}{
		{name: "--yes skips confirmation", args: []string{"--yes"}, wantCleared: true, wantOut: "unset A B"},
		{name: "-y skips confirmation", args: []string{"-y"}, wantCleared: true, wantOut: "Store vacío"},
		{name: "confirms with y", stdin: "y\n", wantCleared: true, wantOut: "Store vacío"},
		{name: "confirms with sí", stdin: " Sí \n", wantCleared: true, wantOut: "Store vacío"},
		{name: "default is no", stdin: "\n", wantOut: "Cancelado"},
		{name: "anything else cancels", stdin: "nope\n", wantOut: "Cancelado"},
		{name: "EOF cancels", stdin: "", wantOut: "Cancelado"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := newTestApp(t)
			a.mustSave(t, "A=1", "B=2")
			a.In = strings.NewReader(tt.stdin)
			if err := a.Clear(tt.args); err != nil {
				t.Fatalf("Clear() error = %v", err)
			}
			if cleared := len(a.entries(t)) == 0; cleared != tt.wantCleared {
				t.Errorf("store vacío = %v, want %v", cleared, tt.wantCleared)
			}
			if !strings.Contains(a.out.String(), tt.wantOut) {
				t.Errorf("falta %q en:\n%s", tt.wantOut, a.out.String())
			}
		})
	}
}

func TestApp_Clear_EmptyStore(t *testing.T) {
	t.Parallel()
	a := newTestApp(t)
	if err := a.Clear(nil); err != nil {
		t.Fatalf("Clear() error = %v", err)
	}
	if !strings.Contains(a.out.String(), "No había nada que borrar") {
		t.Errorf("salida inesperada:\n%s", a.out.String())
	}
}

func TestApp_Clear_UnknownFlag(t *testing.T) {
	t.Parallel()
	a := newTestApp(t)
	a.mustSave(t, "A=1")
	err := a.Clear([]string{"--force"})
	if err == nil || !strings.Contains(err.Error(), "--force") {
		t.Fatalf("Clear(--force) error = %v, want error que mencione --force", err)
	}
	if len(a.entries(t)) != 1 {
		t.Error("Clear() con flag inválido borró el store")
	}
}
