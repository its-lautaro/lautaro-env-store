package store

import (
	"os/exec"
	"testing"
)

func TestQuote(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{name: "empty", value: "", want: "''"},
		{name: "plain", value: "abc", want: "'abc'"},
		{name: "spaces", value: "un usuario", want: "'un usuario'"},
		{name: "single quote", value: "it's", want: `'it'"'"'s'`},
		{name: "only single quote", value: "'", want: `''"'"''`},
		{name: "shell metacharacters untouched", value: `$HOME "x" \n`, want: `'$HOME "x" \n'`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := quote(tt.value); got != tt.want {
				t.Errorf("quote(%q) = %q, want %q", tt.value, got, tt.want)
			}
		})
	}
}

// El contrato real de quote: el shell tiene que devolver el valor exacto, sin interpretar nada.
func TestQuote_ShellRoundTrip(t *testing.T) {
	t.Parallel()
	sh, err := exec.LookPath("sh")
	if err != nil {
		t.Skip("sh no disponible")
	}
	values := []struct {
		name  string
		value string
	}{
		{name: "empty", value: ""},
		{name: "spaces", value: "un usuario"},
		{name: "single quotes", value: "it's 'quoted'"},
		{name: "double quotes", value: `say "hi"`},
		{name: "variable expansion", value: "$HOME ${PATH}"},
		{name: "command substitution", value: "`id` $(id)"},
		{name: "backslashes", value: `a\b\\c\n`},
		{name: "glob and comment", value: "* ? [a] # no comment"},
		{name: "unicode", value: "contraseña ñ 🔐"},
	}
	for _, tt := range values {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			out, err := exec.Command(sh, "-c", "printf '%s' "+quote(tt.value)).Output()
			if err != nil {
				t.Fatalf("sh: %v", err)
			}
			if string(out) != tt.value {
				t.Errorf("sh devolvió %q, want %q", out, tt.value)
			}
		})
	}
}
