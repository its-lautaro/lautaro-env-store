package shell

import (
	"fmt"
	"os/exec"
	"strings"
)

var clipboardTools = [][]string{
	{"pbcopy"},
	{"wl-copy"},
	{"xclip", "-selection", "clipboard"},
	{"xsel", "--clipboard", "--input"},
}

func CopyToClipboard(text string) error {
	var names []string
	for _, t := range clipboardTools {
		names = append(names, t[0])
		if _, err := exec.LookPath(t[0]); err != nil {
			continue
		}
		cmd := exec.Command(t[0], t[1:]...)
		cmd.Stdin = strings.NewReader(text)
		if err := cmd.Run(); err == nil {
			return nil
		}
	}
	return fmt.Errorf("no encontré un comando de clipboard (%s)", strings.Join(names, "/"))
}
