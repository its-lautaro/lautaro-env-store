package commands

import (
	"os"
	"slices"
	"strings"
	"testing"
)

func TestApp_Save(t *testing.T) {
	t.Parallel()
	a := newTestApp(t)
	if err := a.Save([]string{"API_USER=usuario", "API_PASSWORD=s3cr3t-value"}); err != nil {
		t.Fatalf("Save() error = %v", err)
	}
	out := a.out.String()
	if got := a.entries(t); !slices.Equal(got, []string{"API_USER", "API_PASSWORD"}) {
		t.Errorf("entries = %q", got)
	}
	if strings.Contains(out, "s3cr3t-value") || strings.Contains(out, "usuario") {
		t.Errorf("Save() imprimió un valor:\n%s", out)
	}
	for _, w := range []string{"API_USER", "API_PASSWORD", "Variables guardadas", "Nombres copiados"} {
		if !strings.Contains(out, w) {
			t.Errorf("la salida no menciona %q:\n%s", w, out)
		}
	}
	if want := []string{"API_USER API_PASSWORD"}; !slices.Equal(a.clip.copied, want) {
		t.Errorf("clipboard = %q, want %q (solo nombres, nunca valores)", a.clip.copied, want)
	}
	rc, err := os.ReadFile(a.RCFiles[0])
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(rc), a.Store.Path) {
		t.Errorf("Save() no agregó el source al rc:\n%s", rc)
	}
}

func TestApp_Save_MarksReplaced(t *testing.T) {
	t.Parallel()
	a := newTestApp(t)
	a.mustSave(t, "A=1")
	if err := a.Save([]string{"A=2", "B=3"}); err != nil {
		t.Fatalf("Save() error = %v", err)
	}
	if n := strings.Count(a.out.String(), "(reemplazada)"); n != 1 {
		t.Errorf("want exactamente A marcada como reemplazada:\n%s", a.out.String())
	}
}

func TestApp_Save_Errors(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		args        []string
		wantErrPart string
	}{
		{name: "not NAME=value", args: []string{"hola"}, wantErrPart: "NOMBRE=valor"},
		{name: "invalid name", args: []string{"1A=x"}, wantErrPart: "no es un nombre válido"},
		{name: "repeated name", args: []string{"A=1", "A=2"}, wantErrPart: "repetida"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := newTestApp(t)
			err := a.Save(tt.args)
			if err == nil || !strings.Contains(err.Error(), tt.wantErrPart) {
				t.Fatalf("Save(%q) error = %v, want que contenga %q", tt.args, err, tt.wantErrPart)
			}
			if a.out.Len() != 0 {
				t.Errorf("Save() con error imprimió:\n%s", a.out.String())
			}
			if _, statErr := os.Stat(a.Store.Path); !os.IsNotExist(statErr) {
				t.Error("Save() con error creó el store")
			}
		})
	}
}
