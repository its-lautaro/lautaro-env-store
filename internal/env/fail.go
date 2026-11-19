package env

import (
	"fmt"
	"os"
)

// Fail imprime el error con ❌, el uso, y termina con código 1.
func Fail(format string, a ...any) {
	fmt.Fprintf(os.Stderr, "❌ "+format+"\n", a...)
	Usage(1)
}
