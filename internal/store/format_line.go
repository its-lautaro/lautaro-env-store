package store

import "fmt"

const (
	savedMarker = "# saved "
	stampLayout = "2006-01-02 15:04"
)

func formatLine(v Var, stamp string) string {
	return fmt.Sprintf("export %s=%s  %s%s", v.Name, quote(v.Value), savedMarker, stamp)
}
