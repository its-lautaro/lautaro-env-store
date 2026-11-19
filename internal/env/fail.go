package env

import (
	"fmt"
	"os"

	"github.com/its-lautaro/lautaro-env-store/internal/color"
)

// Fail imprime el error con ❌ y un hint para ver la ayuda, y termina con código 1.
func Fail(format string, a ...any) {
	fmt.Fprintln(os.Stderr, color.Red("❌ "+fmt.Sprintf(format, a...)))
	fmt.Fprintln(os.Stderr, color.Dim("💡 Para ver la ayuda: ")+color.Bold("lautaro-env-store --help"))
	os.Exit(1)
}
