package color

import "testing"

// Estos tests cambian la variable global enabled, así que no son paralelos.
func setEnabled(t *testing.T, v bool) {
	t.Helper()
	prev := enabled
	enabled = v
	t.Cleanup(func() { enabled = prev })
}

func TestPaint(t *testing.T) {
	tests := []struct {
		name    string
		enabled bool
		fn      func(string) string
		want    string
	}{
		{name: "disabled returns text untouched", enabled: false, fn: Red, want: "hi"},
		{name: "bold", enabled: true, fn: Bold, want: "\033[1mhi\033[0m"},
		{name: "dim", enabled: true, fn: Dim, want: "\033[2mhi\033[0m"},
		{name: "red", enabled: true, fn: Red, want: "\033[31mhi\033[0m"},
		{name: "green", enabled: true, fn: Green, want: "\033[32mhi\033[0m"},
		{name: "yellow", enabled: true, fn: Yellow, want: "\033[33mhi\033[0m"},
		{name: "cyan", enabled: true, fn: Cyan, want: "\033[36mhi\033[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setEnabled(t, tt.enabled)
			if got := tt.fn("hi"); got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
