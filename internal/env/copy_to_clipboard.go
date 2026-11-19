package env

import (
	"os/exec"
	"strings"
)

// CopyToClipboard copia el texto con el primer comando de clipboard disponible.
func CopyToClipboard(text string) bool {
	candidates := [][]string{
		{"pbcopy"},
		{"wl-copy"},
		{"xclip", "-selection", "clipboard"},
		{"xsel", "--clipboard", "--input"},
	}
	for _, c := range candidates {
		if _, err := exec.LookPath(c[0]); err != nil {
			continue
		}
		cmd := exec.Command(c[0], c[1:]...)
		cmd.Stdin = strings.NewReader(text)
		if err := cmd.Run(); err == nil {
			return true
		}
	}
	return false
}
