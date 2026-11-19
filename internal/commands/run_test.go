package commands

import (
	"slices"
	"strings"
	"testing"
)

func TestApp_Run_ArityErrors(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		args        []string
		wantErrPart string
	}{
		{name: "remove needs a name", args: []string{"--remove"}, wantErrPart: "--remove necesita NOMBRE..."},
		{name: "rm alias needs a name", args: []string{"rm"}, wantErrPart: "rm necesita"},
		{name: "list takes no args", args: []string{"--list", "x"}, wantErrPart: "--list no lleva argumentos"},
		{name: "help takes no args", args: []string{"-h", "x"}, wantErrPart: "-h no lleva argumentos"},
		{name: "init takes no args", args: []string{"--init", "x"}, wantErrPart: "--init no lleva argumentos"},
		{name: "clear takes at most one", args: []string{"--clear", "--yes", "x"}, wantErrPart: "como máximo 1"},
		{name: "unknown flag falls through to save", args: []string{"--nope"}, wantErrPart: "NOMBRE=valor"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := newTestApp(t).Run(tt.args)
			if err == nil || !strings.Contains(err.Error(), tt.wantErrPart) {
				t.Errorf("Run(%q) error = %v, want que contenga %q", tt.args, err, tt.wantErrPart)
			}
		})
	}
}

func TestApp_Run_Dispatch(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		args        []string
		wantOut     string
		wantEntries []string
	}{
		{name: "no args shows help", args: nil, wantOut: "Comandos:", wantEntries: []string{"A"}},
		{name: "--help", args: []string{"--help"}, wantOut: "Comandos:", wantEntries: []string{"A"}},
		{name: "ls alias", args: []string{"ls"}, wantOut: "1 variable", wantEntries: []string{"A"}},
		{name: "rm alias removes", args: []string{"rm", "A"}, wantOut: "eliminada", wantEntries: []string{}},
		{name: "clear --yes", args: []string{"--clear", "--yes"}, wantOut: "Store vacío", wantEntries: []string{}},
		{name: "assignments save", args: []string{"B=2"}, wantOut: "guardada", wantEntries: []string{"A", "B"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := newTestApp(t)
			a.mustSave(t, "A=1")
			if err := a.Run(tt.args); err != nil {
				t.Fatalf("Run(%q) error = %v", tt.args, err)
			}
			if !strings.Contains(a.out.String(), tt.wantOut) {
				t.Errorf("falta %q en:\n%s", tt.wantOut, a.out.String())
			}
			if got := a.entries(t); !slices.Equal(got, tt.wantEntries) {
				t.Errorf("entries = %q, want %q", got, tt.wantEntries)
			}
		})
	}
}
