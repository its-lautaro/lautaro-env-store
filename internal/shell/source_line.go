package shell

import "fmt"

func sourceLine(storePath, marker string) string {
	return fmt.Sprintf("[ -f %q ] && source %q  # %s", storePath, storePath, marker)
}
