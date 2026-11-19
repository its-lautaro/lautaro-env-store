package commands

import (
	"fmt"
	"os"

	"github.com/its-lautaro/lautaro-env-store/internal/color"
)

func Fail(err error) {
	fmt.Fprintln(os.Stderr, color.Red("❌ "+err.Error()))
	fmt.Fprintln(os.Stderr, color.Dim("💡 Para ver la ayuda: ")+color.Bold(Name+" --help"))
	os.Exit(1)
}
