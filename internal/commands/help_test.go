package commands

import (
	"bytes"
	"testing"
)

func TestApp_Help(t *testing.T) {
	t.Parallel()
	var out bytes.Buffer
	a := App{Store: storeAt("/x/.lautaro-env"), RCFiles: []string{"/x/.zshrc"}, Out: &out}
	if err := a.Help(nil); err != nil {
		t.Fatalf("Help() error = %v", err)
	}
	if out.String() != a.helpText()+"\n" {
		t.Errorf("Help() imprimió algo distinto de helpText():\n%s", out.String())
	}
}
