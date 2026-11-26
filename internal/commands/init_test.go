package commands

import (
	"os"
	"strings"
	"testing"
)

func TestApp_Init(t *testing.T) {
	t.Parallel()
	a := newTestApp(t)

	if err := a.Init(nil); err != nil {
		t.Fatalf("Init() error = %v", err)
	}
	info, err := os.Stat(a.Store.Path)
	if err != nil {
		t.Fatalf("Init() no creó el store: %v", err)
	}
	if info.Mode().Perm() != 0o600 {
		t.Errorf("permisos = %o, want 600", info.Mode().Perm())
	}
	if !strings.Contains(a.out.String(), "Agregué el source") {
		t.Errorf("primera llamada no reportó el source agregado:\n%s", a.out.String())
	}

	a.out.Reset()
	if err := a.Init(nil); err != nil {
		t.Fatalf("segundo Init() error = %v", err)
	}
	if !strings.Contains(a.out.String(), "el shell ya lo cargaba") {
		t.Errorf("segunda llamada debería avisar que ya estaba:\n%s", a.out.String())
	}
	rc, err := os.ReadFile(a.RCFiles[0])
	if err != nil {
		t.Fatal(err)
	}
	if n := strings.Count(string(rc), "source"); n != 1 {
		t.Errorf("el rc tiene %d líneas source, want 1:\n%s", n, rc)
	}
}

func TestApp_Init_NoRCFile(t *testing.T) {
	t.Parallel()
	a := newTestApp(t)
	a.RCFiles = []string{a.RCFiles[1]} // .bashrc no existe
	if err := a.Init(nil); err != nil {
		t.Fatalf("Init() error = %v", err)
	}
	if !strings.Contains(a.out.String(), "agregá a tu rc") {
		t.Errorf("want instrucción manual de source:\n%s", a.out.String())
	}
}
