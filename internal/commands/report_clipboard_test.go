package commands

import (
	"errors"
	"strings"
	"testing"
)

func TestApp_reportClipboard(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		names   []string
		err     error
		wantOut string
	}{
		{name: "one name", names: []string{"A"}, wantOut: "📋 Nombre copiado al clipboard"},
		{name: "several names", names: []string{"A", "B"}, wantOut: "📋 Nombres copiados al clipboard"},
		{name: "clipboard fails", names: []string{"A"}, err: errors.New("sin pbcopy"), wantOut: "sin pbcopy; los nombres quedan arriba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := newTestApp(t)
			a.clip.err = tt.err
			a.reportClipboard(tt.names)
			if !strings.Contains(a.out.String(), tt.wantOut) {
				t.Errorf("salida = %q, want que contenga %q", a.out.String(), tt.wantOut)
			}
		})
	}
}
