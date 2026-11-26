package store

import (
	"os/exec"
	"strings"
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
		{name: "newline uses ANSI-C quoting", value: "l1\nl2", want: `$'l1\nl2'`},
		{name: "tab and carriage return", value: "a\tb\r", want: `$'a\tb\r'`},
		{name: "ANSI-C escapes quote and backslash", value: "it's\\\n", want: `$'it\'s\\\n'`},
		{name: "other control chars as hex", value: "a\x01\x7f\n", want: `$'a\x01\x7f\n'`},
		{name: "unicode kept raw", value: "ñ\n", want: "$'ñ\\n'"},
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

// Los valores con caracteres de control usan $'...', que entienden bash y zsh (los rc donde se carga el store).
func TestQuote_ControlCharsRoundTrip(t *testing.T) {
	t.Parallel()
	values := []struct {
		name  string
		value string
	}{
		{name: "multiline", value: "l1\nl2\n"},
		{name: "pem like", value: "-----BEGIN KEY-----\nabc+/=\n-----END KEY-----"},
		{name: "quotes backslashes and newline", value: "it's \"x\" \\n \\\n$HOME `id`"},
		{name: "tab cr and other controls", value: "a\tb\r\x01\x1b[0m\x7f"},
		{name: "hex digit after escape", value: "\x01abc"},
		{name: "unicode", value: "contraseña\n🔐"},
	}
	for _, shell := range []string{"bash", "zsh"} {
		t.Run(shell, func(t *testing.T) {
			t.Parallel()
			path, err := exec.LookPath(shell)
			if err != nil {
				t.Skipf("%s no disponible", shell)
			}
			for _, tt := range values {
				t.Run(tt.name, func(t *testing.T) {
					t.Parallel()
					q := quote(tt.value)
					if strings.ContainsAny(q, "\n\r") {
						t.Fatalf("quote(%q) = %q ocupa más de una línea", tt.value, q)
					}
					out, err := exec.Command(path, "-c", "printf '%s' "+q).Output()
					if err != nil {
						t.Fatalf("%s: %v", shell, err)
					}
					if string(out) != tt.value {
						t.Errorf("%s devolvió %q, want %q", shell, out, tt.value)
					}
				})
			}
		})
	}
}
