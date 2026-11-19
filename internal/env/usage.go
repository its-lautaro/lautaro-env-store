package env

import (
	"fmt"
	"os"
)

// Usage imprime el texto de ayuda en stderr y termina con el código indicado.
func Usage(code int) {
	fmt.Fprintln(os.Stderr, UsageTxt)
	os.Exit(code)
}
