package shell

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// fakeTool crea un ejecutable en dir que corre script con sh.
func fakeTool(t *testing.T, dir, name, script string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(dir, name), []byte("#!/bin/sh\n"+script+"\n"), 0o755); err != nil {
		t.Fatal(err)
	}
}

// catTo devuelve un script que guarda su stdin en out, con cat por ruta absoluta (PATH se reemplaza en los tests).
func catTo(t *testing.T, out string) string {
	t.Helper()
	cat, err := exec.LookPath("cat")
	if err != nil {
		t.Skip("cat no disponible")
	}
	return `"` + cat + `" > "` + out + `"`
}

// Estos tests cambian PATH, así que no pueden ser paralelos.

func TestCopyToClipboard_NoTool(t *testing.T) {
	t.Setenv("PATH", t.TempDir())
	err := CopyToClipboard("x")
	if err == nil {
		t.Fatal("CopyToClipboard() sin herramientas: error = nil, want error")
	}
	for _, tool := range []string{"pbcopy", "wl-copy", "xclip", "xsel"} {
		if !strings.Contains(err.Error(), tool) {
			t.Errorf("el error no menciona %s: %v", tool, err)
		}
	}
}

func TestCopyToClipboard_PassesTextOnStdin(t *testing.T) {
	dir := t.TempDir()
	out := filepath.Join(dir, "out")
	fakeTool(t, dir, "xsel", catTo(t, out))
	t.Setenv("PATH", dir)
	if err := CopyToClipboard("A B C"); err != nil {
		t.Fatalf("CopyToClipboard() error = %v", err)
	}
	data, err := os.ReadFile(out)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "A B C" {
		t.Errorf("stdin recibido = %q, want %q", data, "A B C")
	}
}

func TestCopyToClipboard_FallsBackAfterFailure(t *testing.T) {
	dir := t.TempDir()
	out := filepath.Join(dir, "out")
	fakeTool(t, dir, "pbcopy", "exit 1")
	fakeTool(t, dir, "wl-copy", catTo(t, out))
	t.Setenv("PATH", dir)
	if err := CopyToClipboard("x"); err != nil {
		t.Fatalf("CopyToClipboard() error = %v, want fallback a wl-copy", err)
	}
	if _, err := os.Stat(out); err != nil {
		t.Errorf("wl-copy no se ejecutó: %v", err)
	}
}

func TestCopyToClipboard_ReportsLastFailure(t *testing.T) {
	dir := t.TempDir()
	fakeTool(t, dir, "xclip", "exit 3")
	t.Setenv("PATH", dir)
	err := CopyToClipboard("x")
	if err == nil || !strings.Contains(err.Error(), "xclip falló") {
		t.Errorf("CopyToClipboard() error = %v, want 'xclip falló'", err)
	}
}
