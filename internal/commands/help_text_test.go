package commands

import (
	"strings"
	"testing"
)

func TestApp_helpText(t *testing.T) {
	t.Parallel()
	a := App{Store: storeAt("/x/.lautaro-env"), RCFiles: []string{"/x/.zshrc", "/x/.bashrc"}}
	got := a.helpText()
	want := []string{Name, "/x/.lautaro-env", "/x/.zshrc o /x/.bashrc", StoreEnvVar}
	for _, c := range a.Commands() {
		want = append(want, c.Names...)
		want = append(want, c.Description)
	}
	for _, w := range want {
		if !strings.Contains(got, w) {
			t.Errorf("la ayuda no menciona %q", w)
		}
	}
	if strings.HasSuffix(got, "\n") {
		t.Error("la ayuda termina en salto de línea; Println agrega otro")
	}
}
