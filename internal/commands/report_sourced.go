package commands

import (
	"fmt"
	"strings"

	"github.com/its-lautaro/lautaro-env-store/internal/color"
)

func (a App) reportSourced(rc string, added bool) {
	switch {
	case rc == "":
		fmt.Println(color.Yellow(fmt.Sprintf("⚠️  No encontré %s; agregá a tu rc:  source %q", strings.Join(a.RCFiles, " ni "), a.Store.Path)))
	case added:
		fmt.Printf("🧩 Agregué el source de %s a %s\n", color.Dim(a.Store.Path), color.Dim(rc))
	}
}
